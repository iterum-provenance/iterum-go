package messageq

import (
	"sync"
	"time"

	"github.com/prometheus/common/log"
	"github.com/streadway/amqp"

	desc "github.com/iterum-provenance/iterum-go/descriptors"
	"github.com/iterum-provenance/iterum-go/transmit"
	"github.com/iterum-provenance/iterum-go/util"
)

// Sender is the structure that listens to a channel and redirects messages to rabbitMQ
type Sender struct {
	toSend         <-chan transmit.Serializable // data.RemoteFragmentDesc
	ToLineate      chan<- transmit.Serializable // desc.RemoteFragmentDesc
	Publishers     map[string]QPublisher
	publisherGroup *sync.WaitGroup
	TargetQueue    string
	BrokerURL      string
	fragments      int
}

// NewSender creates a new sender which receives messages from a channel and sends them on the message queue.
func NewSender(toSend, toLineate chan transmit.Serializable, brokerURL, targetQueue string) (sender Sender, err error) {

	sender = Sender{
		toSend,
		toLineate,
		make(map[string]QPublisher),
		&sync.WaitGroup{},
		targetQueue,
		brokerURL,
		0,
	}
	return
}

// spawnPublisher creates a new publisher for a specific queue using the same amqp Connection
func (sender *Sender) spawnPublisher(conn *amqp.Connection, targetQueue string) {
	ch, err := conn.Channel() // Eventually closed by the QPublisher
	util.Ensure(err, "Opened channel")
	publisher := NewQPublisher(make(chan transmit.Serializable, 10), ch, targetQueue)
	publisher.Start(sender.publisherGroup)
	sender.Publishers[targetQueue] = publisher
}

// StartBlocking listens to the channel, and send remoteFragments to the message queue on the OUTPUT_QUEUE queue.
func (sender *Sender) StartBlocking() {
	log.Infof("Connecting to %s.\n", sender.BrokerURL)
	conn, err := amqp.Dial(sender.BrokerURL)
	util.Ensure(err, "Connected to RabbitMQ")
	defer conn.Close()

	sender.spawnPublisher(conn, sender.TargetQueue)

	for remoteFragMsg := range sender.toSend {
		remoteFragment := *remoteFragMsg.(*desc.RemoteFragmentDesc)

		queueName := sender.TargetQueue
		if remoteFragment.Metadata.TargetQueue != nil {
			queueName = *remoteFragment.Metadata.TargetQueue
			if _, ok := sender.Publishers[queueName]; !ok {
				sender.spawnPublisher(conn, queueName)
			}
		}

		log.Debugf("Fragment sender sending %v to toLineate and Publisher\n", remoteFragment)
		// Wrap in messagequeue specific struct
		mqFragment := newFragmentDesc(remoteFragment)
		// hand to publisher
		sender.Publishers[queueName].ToPublish <- &mqFragment
		// lineate data
		sender.ToLineate <- &remoteFragment
		sender.fragments++
	}

	sender.Stop()
}

// Start asychronously calls StartBlocking via Gorouting
func (sender Sender) Start(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		startTime := time.Now()
		sender.StartBlocking()
		log.Infof("sender ran for %v", time.Now().Sub(startTime))
	}()
}

// Stop finishes up and notifies the user of its progress
func (sender Sender) Stop() {
	log.Infof("MQSender finishing up, published %v messages\n", sender.fragments)
	close(sender.ToLineate)
	for _, pub := range sender.Publishers {
		close(pub.ToPublish)
	}
	sender.publisherGroup.Wait()
}

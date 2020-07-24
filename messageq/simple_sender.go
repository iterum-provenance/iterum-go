package messageq

import (
	"sync"
	"time"

	"github.com/prometheus/common/log"
	"github.com/streadway/amqp"

	"github.com/iterum-provenance/iterum-go/transmit"
	"github.com/iterum-provenance/iterum-go/util"
)

// SimpleSender is the structure that listens to a channel and redirects messages to rabbitMQ
type SimpleSender struct {
	ToSend      chan transmit.Serializable
	TargetQueue string
	BrokerURL   string
	messages    int
	publisher   *QPublisher
}

// NewSimpleSender creates a new sender which receives messages from a channel and sends them on the message queue.
func NewSimpleSender(toSend chan transmit.Serializable, brokerURL, targetQueue string) (sender SimpleSender) {
	return SimpleSender{
		toSend,
		targetQueue,
		brokerURL,
		0,
		nil,
	}
}

// spawnPublisher creates a new publisher for a specific queue using the same amqp Connection
func (sender *SimpleSender) spawnPublisher(conn *amqp.Connection) {
	ch, err := conn.Channel() // Eventually closed by the QPublisher
	util.Ensure(err, "Opened channel")
	pub := NewQPublisher(make(chan transmit.Serializable, 10), ch, sender.TargetQueue)
	sender.publisher = &pub
}

// StartBlocking listens to the channel, and send remoteFragments to the message queue on the OUTPUT_QUEUE queue.
func (sender *SimpleSender) StartBlocking() {
	log.Infof("Connecting to %s.\n", sender.BrokerURL)
	conn, err := amqp.Dial(sender.BrokerURL)
	util.Ensure(err, "Connected to RabbitMQ")
	defer conn.Close()

	sender.spawnPublisher(conn)
	wg := &sync.WaitGroup{}
	defer wg.Wait()
	sender.publisher.Start(wg)

	for msg := range sender.ToSend {
		log.Debugf("Simple sender got %v\n", msg)
		// hand to publisher
		sender.publisher.ToPublish <- msg
		sender.messages++
	}
	sender.Stop()
}

// Start asychronously calls StartBlocking via Gorouting
func (sender *SimpleSender) Start(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		startTime := time.Now()
		sender.StartBlocking()
		log.Infof("sender ran for %v", time.Now().Sub(startTime))
	}()
}

// Stop finishes up and notifies the user of its progress
func (sender *SimpleSender) Stop() {
	log.Infof("SimpleSender finishing up, published %v messages\n", sender.messages)
	close(sender.publisher.ToPublish)
}

package messageq

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/iterum-provenance/iterum-go/transmit"
	"github.com/iterum-provenance/iterum-go/util"
	"github.com/prometheus/common/log"
	"github.com/streadway/amqp"
)

// Consumer consumes messages from the messagequeue without acknowledging them
type Consumer struct {
	Output    chan<- transmit.Serializable // *desc.RemoteFragmentDesc
	Unacked   chan amqp.Delivery
	Exit      chan bool
	mqChannel *amqp.Channel
	QueueName string
	consumed  int
}

// NewConsumer creates a message consumer for a listener
func NewConsumer(out chan transmit.Serializable, mqChannel *amqp.Channel, inputQueue string) (consumer Consumer) {
	return Consumer{
		out,
		make(chan amqp.Delivery, 10),
		make(chan bool, 1),
		mqChannel,
		inputQueue,
		0,
	}
}

func (consumer *Consumer) handleRemoteFragment(message amqp.Delivery) {
	var mqFragment MqFragmentDesc
	err := mqFragment.Deserialize(message.Body)
	if err != nil {
		log.Errorln(err)
	}
	log.Debugf("Received a mqFragment: %v\n", mqFragment)
	var remoteFragment = mqFragment.RemoteFragmentDesc

	consumer.Output <- &remoteFragment
	consumer.Unacked <- message
}

// StartBlocking listens on the rabbitMQ messagequeue and redirects messages on the INPUT_QUEUE to a channel
func (consumer *Consumer) StartBlocking() {
	q, err := consumer.mqChannel.QueueDeclare(
		consumer.QueueName, // name
		false,              // durable
		false,              // delete when unused
		false,              // exclusive
		false,              // no-wait
		nil,                // arguments
	)
	util.Ensure(err, fmt.Sprintf("Created queue '%v'", consumer.QueueName))

	id := rand.Int()
	consumerName := fmt.Sprintf("%v%v%v", q.Name, "-consumer-", id)
	mqMessages, err := consumer.mqChannel.Consume(
		q.Name,       // queue
		consumerName, // consumer
		false,        // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	util.Ensure(err, "Registered consumer to queue")

	log.Infof("Started consuming messages from the MQ.\n")
	for consumer.Exit != nil || mqMessages != nil {
		select {
		case message, ok := <-mqMessages:
			if !ok {
				mqMessages = nil
				log.Warnln("Consuming message from remote MQ returned !ok, either final stage or an error occurred. Anyway, stopping...")
				message.Ack(true)
			} else {
				consumer.handleRemoteFragment(message)
				consumer.consumed++
			}
		case <-consumer.Exit:
			// Allowed to stop now. Finish current messages left in channel
			consumer.mqChannel.Cancel(consumerName, false)
			consumer.Exit = nil
		}
	}
	consumer.Stop()
}

// Start asychronously calls StartBlocking via a Goroutine
func (consumer *Consumer) Start(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		startTime := time.Now()
		consumer.StartBlocking()
		log.Infof("consumer ran for %v", time.Now().Sub(startTime))
	}()
}

// Stop finishes up the consumer
func (consumer *Consumer) Stop() {
	log.Infof("MQConsumer finishing up, consumed %v messages\n", consumer.consumed)
	close(consumer.Output)
	close(consumer.Unacked)
}

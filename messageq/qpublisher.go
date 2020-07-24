package messageq

import (
	"fmt"
	"sync"
	"time"

	"github.com/iterum-provenance/iterum-go/transmit"
	"github.com/iterum-provenance/iterum-go/util"
	"github.com/prometheus/common/log"
	"github.com/streadway/amqp"
)

// QPublisher is the structure that listens to a channel and publishes messages to rabbitMQ
type QPublisher struct {
	ToPublish   chan transmit.Serializable // data.RemoteFragmentDesc
	Channel     *amqp.Channel
	Queue       *amqp.Queue
	TargetQueue string
	fragments   int
}

// NewQPublisher creates a new qpublisher which receives messages from a channel and sends them on the message queue.
func NewQPublisher(toPublish chan transmit.Serializable, channel *amqp.Channel, targetQueue string) QPublisher {
	return QPublisher{
		toPublish,
		channel,
		nil,
		targetQueue,
		0,
	}
}

// DeclareQueue defines the target queue
func (qpublisher *QPublisher) DeclareQueue() {
	q, err := qpublisher.Channel.QueueDeclare(
		qpublisher.TargetQueue, // name
		false,                  // durable
		false,                  // delete when unused
		false,                  // exclusive
		false,                  // no-wait
		nil,                    // arguments
	)
	util.Ensure(err, fmt.Sprintf("Created queue '%v'", qpublisher.TargetQueue))
	qpublisher.Queue = &q
}

// StartBlocking listens to the channel, and send remoteFragments to the message queue on the OUTPUT_QUEUE queue.
func (qpublisher *QPublisher) StartBlocking() {
	qpublisher.DeclareQueue()

	for msg := range qpublisher.ToPublish {

		log.Debugf("Sending %v to queue '%v'\n", msg, qpublisher.Queue.Name)

		body, err := msg.Serialize()
		if err != nil {
			log.Errorln(err)
		}

		err = qpublisher.Channel.Publish(
			"",                    // exchange
			qpublisher.Queue.Name, // routing key
			false,                 // mandatory
			false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "application/json",
				Body:         body,
			})
		if err != nil {
			log.Errorln(err)
		}
		qpublisher.fragments++
	}

	qpublisher.Stop()
}

// Start asychronously calls StartBlocking via Gorouting
func (qpublisher *QPublisher) Start(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		startTime := time.Now()
		qpublisher.StartBlocking()
		log.Infof("qpublisher ran for %v", time.Now().Sub(startTime))
	}()
}

// Stop finishes up and notifies the user of its progress
func (qpublisher *QPublisher) Stop() {
	log.Infof("MQPublisher finishing up, published %v messages to %v\n", qpublisher.fragments, qpublisher.Queue.Name)
	qpublisher.Channel.Close()
}

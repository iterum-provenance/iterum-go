package messageq

import (
	"fmt"
	"sync"
	"time"

	"github.com/prometheus/common/log"
	"github.com/streadway/amqp"

	"github.com/iterum-provenance/iterum-go/transmit"
	"github.com/iterum-provenance/iterum-go/util"
)

// Listener is the structure that listens to RabbitMQ and redirects messages to a channel
type Listener struct {
	BrokerURL     string
	PrefetchCount int
	TargetQueue   string
	CanExit       chan bool
	exit          chan bool
	consumer      Consumer
	acknowledger  Acknowledger
}

// NewListener creates a new message queue listener
func NewListener(output, toAcknowledge chan transmit.Serializable, brokerURL, inputQueue string, prefetchCount int) (listener Listener, err error) {
	consumer := NewConsumer(output, nil, inputQueue)
	acknowledger := NewAcknowledger(consumer.Unacked, toAcknowledge)

	listener = Listener{
		BrokerURL:     brokerURL,
		PrefetchCount: prefetchCount,
		TargetQueue:   inputQueue,
		CanExit:       make(chan bool, 1),
		exit:          make(chan bool, 1),
		consumer:      consumer,
		acknowledger:  acknowledger,
	}
	return
}

func (listener *Listener) messagesLeftChecker(ch *amqp.Channel) {
	var err error = nil
	qChecker := amqp.Queue{Messages: 9999999} /// >0 init value. form of do-while
	for qChecker.Messages > 0 {
		qChecker, err = ch.QueueInspect(listener.TargetQueue)
		if err != nil {
			log.Fatal(err)
		}
		log.Infof("%v messages remaining in Queue. Closing when zero\n", qChecker.Messages)
		time.Sleep(5 * time.Second)
	}
	log.Infof("MQListener consumed all messages\n")
	listener.consumer.Exit <- true
	close(listener.consumer.Exit)
	listener.exit <- true
	close(listener.exit)
}

// StartBlocking listens on the rabbitMQ messagequeue and redirects messages on the INPUT_QUEUE to a channel
func (listener *Listener) StartBlocking() {
	wg := &sync.WaitGroup{}

	log.Infof("Connecting to %s.\n", listener.BrokerURL)
	conn, err := amqp.Dial(listener.BrokerURL)
	util.Ensure(err, "Connected to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	util.Ensure(err, "Opened channel")
	defer ch.Close()
	if listener.PrefetchCount > 0 {
		err := ch.Qos(
			listener.PrefetchCount, // prefetch count
			0,                      // prefetch size
			false,                  // global
		)
		util.Ensure(err, fmt.Sprintf("QoS was set to %v", listener.PrefetchCount))
	}

	listener.consumer.mqChannel = ch

	listener.consumer.Start(wg)
	listener.acknowledger.Start(wg)

	for listener.CanExit != nil || listener.exit != nil {
		select {
		case canExitOnEmpty := <-listener.CanExit:
			if canExitOnEmpty {
				listener.CanExit = nil
				go listener.messagesLeftChecker(ch)
			}
		case <-listener.exit:
			listener.exit = nil
		}
	}

	log.Infof("MQListener awaiting acknowledger and consumer\n")
	wg.Wait()
	log.Infof("MQListener finished\n")
}

// Start asychronously calls StartBlocking via Gorouting
func (listener *Listener) Start(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		startTime := time.Now()
        listener.StartBlocking()
	    log.Infof("listener ran for %v", time.Now().Sub(startTime))
	}()
}

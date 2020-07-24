package messageq

import (
	"sync"
	"time"

	desc "github.com/iterum-provenance/iterum-go/descriptors"
	"github.com/iterum-provenance/iterum-go/transmit"
	"github.com/prometheus/common/log"
	"github.com/streadway/amqp"
)

// Acknowledger is responsible for sending acknowledgement messages to the MQBroker
type Acknowledger struct {
	acknowledge  <-chan transmit.Serializable    // *desc.FinishedFragmentMessage
	pending      map[desc.IterumID]amqp.Delivery // unacknowledged
	consumed     <-chan amqp.Delivery
	acknowledged int
}

// NewAcknowledger isntantiates a new Acknowledger struct
func NewAcknowledger(consumed chan amqp.Delivery, toAcknowledge chan transmit.Serializable) Acknowledger {
	return Acknowledger{
		acknowledge:  toAcknowledge,
		consumed:     consumed,
		pending:      map[desc.IterumID]amqp.Delivery{},
		acknowledged: 0,
	}
}

// StartBlocking listens on the two channels for new messages that were consumed but not acknowledged
// and for messages to acknowledge
func (ack *Acknowledger) StartBlocking() {
	for ack.consumed != nil || ack.acknowledge != nil {
		select {
		case msg, ok := <-ack.consumed:
			if !ok {
				ack.consumed = nil
				continue
			}
			var mqFragment MqFragmentDesc
			err := mqFragment.Deserialize(msg.Body)
			if err != nil {
				log.Errorln(err)
			}
			if _, ok := ack.pending[mqFragment.Metadata.FragmentID]; ok {
				log.Errorf("Duplicate FragmentID found for pending map: '%v'\n", mqFragment.Metadata.FragmentID)
			}
			ack.pending[mqFragment.Metadata.FragmentID] = msg
		case msg, ok := <-ack.acknowledge:
			if !ok {
				ack.acknowledge = nil
				continue
			}
			doneMsg := *msg.(*desc.FinishedFragmentMessage)
			delivery, ok := ack.pending[doneMsg.FragmentID]
			if !ok {
				log.Errorf("Received finished FragmentID that was not pending: '%v'\n", doneMsg.FragmentID)
				continue
			}
			err := delivery.Ack(false)
			if err != nil {
				log.Errorln(err)
			}
			delete(ack.pending, doneMsg.FragmentID)
			ack.acknowledged++
		}
	}
	if len(ack.pending) != 0 {
		log.Errorln("Acknowledger.pending is not empty at end of lifecycle, unacknowledged messages remain")
	}
	log.Infof("Both input channels are closed, finishing up Acknowledger. Acknowledged %v messages", ack.acknowledged)
}

// Start asychronously calls StartBlocking via Gorouting
func (ack *Acknowledger) Start(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		startTime := time.Now()
		ack.StartBlocking()
		log.Infof("ack ran for %v", time.Now().Sub(startTime))
	}()
}

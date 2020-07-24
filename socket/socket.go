package socket

import (
	"net"
	"os"
	"sync"
	"time"

	"github.com/prometheus/common/log"

	"github.com/iterum-provenance/iterum-go/util"

	"github.com/iterum-provenance/iterum-go/transmit"
)

// Socket is a structure holding a listener, accepting connections
// Channel is a channel that external things can post messages on take from that are
// supposed to be sent to or from the connections
type Socket struct {
	Listener net.Listener
	Channel  chan transmit.Serializable
	handler  ConnHandler
	stop     chan bool
}

// ConnHandler is a handler function ran in a goroutine upon a socket accepting a new connection
type ConnHandler func(Socket, net.Conn)

// NewSocket sets up a listener at the given socketPath and links the passed channel
// with the given bufferSize. It returns an error on failure
func NewSocket(socketPath string, channel chan transmit.Serializable, handler ConnHandler) (socket Socket, err error) {
	defer util.ReturnErrOnPanic(&err)
	if _, errExist := os.Stat(socketPath); !os.IsNotExist(errExist) {
		err = os.Remove(socketPath)
		util.Ensure(err, "Existing socket file removed")
	}

	listener, err := net.Listen("unix", socketPath)
	util.Ensure(err, "Server created")

	socket = Socket{
		listener,
		channel,
		handler,
		make(chan bool, 1),
	}
	return
}

// StartBlocking enters an endless loop accepting connections and calling the handler function
// in a goroutine
func (socket Socket) StartBlocking() {
	for {
		conn, err := socket.Listener.Accept()
		select {
		case <-socket.stop:
			// If we were told to stop (an error would occur, but this is expected)
			return
		default:
			if err != nil {
				// If we weren't told to stop but an error occurred
				log.Errorln(err)
				return
			}
			defer conn.Close()
			go socket.handler(socket, conn)
		}
	}
}

// Start asychronously calls StartBlocking via Gorouting
func (socket Socket) Start(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		startTime := time.Now()
		socket.StartBlocking()
		log.Infof("s ran for %v", time.Now().Sub(startTime))
	}()
}

// Stop tries to close the listener of the socket and returns an error on failure
func (socket *Socket) Stop() error {
	log.Infoln("Stopping socket server...")
	socket.stop <- true
	return socket.Listener.Close()
}

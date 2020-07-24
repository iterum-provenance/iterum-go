package socket

import (
	"sync"

	"github.com/iterum-provenance/iterum-go/util"

	"github.com/iterum-provenance/iterum-go/transmit"
)

// Pipe represents a bidirectional connection between an iterum sidecar and transformation step
// ToTarget and FromTarget refer to the channels in the two sockets
// Messages supposed to go towards Target can be put on ToTarget and message from the Target are put on FromTarget
type Pipe struct {
	SocketFrom Socket
	SocketTo   Socket
	FromTarget chan transmit.Serializable
	ToTarget   chan transmit.Serializable
}

// NewPipe creates and initiates a new Pipe
func NewPipe(fromFile, toFile string, fromChannel, toChannel chan transmit.Serializable, fromHandler, toHandler ConnHandler) Pipe {
	fromSocket, err := NewSocket(fromFile, fromChannel, fromHandler)
	util.Ensure(err, "From Socket succesfully opened and listening")
	toSocket, err := NewSocket(toFile, toChannel, toHandler)
	util.Ensure(err, "Towards Socket succesfully opened and listening")

	return Pipe{fromSocket, toSocket, fromSocket.Channel, toSocket.Channel}
}

// Start calls start on both of the pipe's sockets
func (p Pipe) Start(wg *sync.WaitGroup) {
	p.SocketFrom.Start(wg)
	p.SocketTo.Start(wg)
}

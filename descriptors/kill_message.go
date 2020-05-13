package descriptors

import (
	"encoding/json"
	"fmt"

	"github.com/iterum-provenance/iterum-go/transmit"
)

// KillMessage is a message noting that either the sidecar or the program should stop
// a KillMessage indicates the final message on a socket after which the connection can be broken
type KillMessage struct {
	Status string `json:"status"`
}

const defaultStatus = "complete"

// NewKillMessage creates a new default instance of KillMessage
func NewKillMessage() KillMessage {
	return KillMessage{}
}

// Serialize tries to transform `km` into a json encoded bytearray. Errors on failure
func (km *KillMessage) Serialize() (data []byte, err error) {
	data, err = json.Marshal(km)
	if err != nil {
		err = transmit.ErrSerialization(err)
	}
	return

}

// Deserialize tries to decode a json encoded byte array into `km`. Errors on failure
func (km *KillMessage) Deserialize(data []byte) (err error) {
	err = json.Unmarshal(data, km)
	if err != nil {
		err = transmit.ErrSerialization(err)
	}
	if km.Status != defaultStatus {
		err = transmit.ErrSerialization(fmt.Errorf("Incomplete KillMessage, status not '%v'", defaultStatus))
	}

	return
}

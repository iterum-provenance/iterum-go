package lineage

import (
	"encoding/json"

	desc "github.com/iterum-provenance/iterum-go/descriptors"
	"github.com/iterum-provenance/iterum-go/transmit"
)

// Message is the structure holding the lineage information of a fragment
type Message struct {
	ProcessName string                  `json:"transformation_step"`
	Fragment    desc.RemoteFragmentDesc `json:"description"`
}

// Serialize tries to transform `msg` into a json encoded bytearray. Errors on failure
func (msg *Message) Serialize() (data []byte, err error) {
	data, err = json.Marshal(msg)
	if err != nil {
		err = transmit.ErrSerialization(err)
	}
	return
}

// Deserialize tries to decode a json encoded byte array into `msg`. Errors on failure
func (msg *Message) Deserialize(data []byte) (err error) {
	err = json.Unmarshal(data, msg)
	if err != nil {
		err = transmit.ErrSerialization(err)
	}
	return
}

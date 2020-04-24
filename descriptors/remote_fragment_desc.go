package descriptors

import (
	"encoding/json"

	"github.com/iterum-provenance/iterum-go/transmit"
)

// RemoteFragmentDesc is a structure describing an iterum fragment
// and how it is stored on the remote minio storage
type RemoteFragmentDesc struct {
	Files []RemoteFileDesc `json:"files"`
}

// Serialize tries to transform `f` into a json encoded bytearray. Errors on failure
func (rf *RemoteFragmentDesc) Serialize() (data []byte, err error) {
	data, err = json.Marshal(rf)
	if err != nil {
		err = transmit.ErrSerialization(err)
	}
	return

}

// Deserialize tries to decode a json encoded byte array into `f`. Errors on failure
func (rf *RemoteFragmentDesc) Deserialize(data []byte) (err error) {
	err = json.Unmarshal(data, rf)
	if err != nil {
		err = transmit.ErrSerialization(err)
	}
	return
}

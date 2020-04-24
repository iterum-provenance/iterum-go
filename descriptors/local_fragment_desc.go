package descriptors

import (
	"encoding/json"

	"github.com/iterum-provenance/iterum-go/transmit"
)

// LocalFragmentDesc is a structure describing an iterum fragment
// and how it is stored on the program's machine's local volume
type LocalFragmentDesc struct {
	Files []LocalFileDesc `json:"files"`
}

// Serialize tries to transform `f` into a json encoded bytearray. Errors on failure
func (lf *LocalFragmentDesc) Serialize() (data []byte, err error) {
	data, err = json.Marshal(lf)
	if err != nil {
		err = transmit.ErrSerialization(err)
	}
	return

}

// Deserialize tries to decode a json encoded byte array into `f`. Errors on failure
func (lf *LocalFragmentDesc) Deserialize(data []byte) (err error) {
	err = json.Unmarshal(data, lf)
	if err != nil {
		err = transmit.ErrSerialization(err)
	}
	return
}

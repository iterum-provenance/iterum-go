package descriptors

import (
	"encoding/json"

	"github.com/iterum-provenance/iterum-go/transmit"
)

// LocalFragmentDesc is a structure describing an iterum fragment
// and how it is stored on the program's machine's local volume
type LocalFragmentDesc struct {
	Files    []LocalFileDesc `json:"files"`
	Metadata LocalMetadata   `json:"metadata"` // Additional information on this fragment
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
	if err = json.Unmarshal(data, lf); err != nil {
		return transmit.ErrSerialization(err)
	}
	// Metadata does not need validation here:
	// LocalFragmentDesc come from transformation steps, they do not introduce
	// all necessary information such as a new FragmentID, this is done in the sidecar
	// Therefore a new local fragment can never be valid.
	// Validation is done in the next step where this fragment is deserialized as a remote one
	return err
}

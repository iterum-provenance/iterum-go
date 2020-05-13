package descriptors

import (
	"encoding/json"
	"fmt"

	"github.com/iterum-provenance/iterum-go/transmit"
)

// FinishedFragmentMessage contains an id of a fragment that is no longer
// needed by the transformation step. Meaning we can acknowledge this fragment
type FinishedFragmentMessage struct {
	FragmentID IterumID `json:"fragment_id"`
}

// Serialize tries to transform `ffm` into a json encoded bytearray. Errors on failure
func (ffm *FinishedFragmentMessage) Serialize() (data []byte, err error) {
	data, err = json.Marshal(ffm)
	if err != nil {
		err = transmit.ErrSerialization(err)
	}
	return

}

// Deserialize tries to decode a json encoded byte array into `ffm`. Errors on failure
func (ffm *FinishedFragmentMessage) Deserialize(data []byte) (err error) {
	err = json.Unmarshal(data, ffm)
	if err != nil {
		err = transmit.ErrSerialization(err)
	}
	if !ffm.FragmentID.IsValid() {
		err = transmit.ErrSerialization(fmt.Errorf("No valid FragmentID passed: '%v'", ffm.FragmentID))
	}
	return
}

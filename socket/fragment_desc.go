package socket

import (
	"encoding/json"
	"errors"

	desc "github.com/iterum-provenance/iterum-go/descriptors"
	"github.com/iterum-provenance/iterum-go/transmit"
)

// fragmentDesc is a structure describing a local iterum fragment that
// can be exchanged between the sidecar and the TS
type fragmentDesc struct {
	desc.LocalFragmentDesc
}

// newFragmentDesc makes a new fragment description suited for use within socket package
func newFragmentDesc(files []desc.LocalFileDesc) fragmentDesc {
	sfd := fragmentDesc{desc.LocalFragmentDesc{}}
	sfd.Files = files
	return sfd
}

// Serialize tries to transform `sfd` into a json encoded bytearray. Errors on failure
func (sfd *fragmentDesc) Serialize() (data []byte, err error) {
	data, err = json.Marshal(sfd)
	if err != nil {
		err = transmit.ErrSerialization(err)
	}
	return
}

// Deserialize tries to decode a json encoded byte array into `sfd`. Errors on failure
func (sfd *fragmentDesc) Deserialize(data []byte) (err error) {
	err = json.Unmarshal(data, sfd)
	if err != nil {
		err = transmit.ErrSerialization(err)
	}
	if len(sfd.Files) == 0 {
		err = transmit.ErrSerialization(errors.New("FragmentDescription cannot contain 0 files"))
	}
	return
}

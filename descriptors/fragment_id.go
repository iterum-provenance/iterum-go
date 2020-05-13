package descriptors

import (
	"math/rand"
	"time"
)

// IterumID is an id that belongs to a fragment and used to correlate fragments
type IterumID string

const allowedChars = "abcdefghijklmnopqrstuvwxyz0123456789"
const idLength = 64

func init() {
	rand.Seed(time.Now().UnixNano())
}

// NewIterumID generates a new IterumID of given length
func NewIterumID() IterumID {
	b := make([]byte, idLength)
	for i := range b {
		b[i] = allowedChars[rand.Int63()%int64(len(allowedChars))]
	}
	return IterumID(b)
}

func (iid IterumID) String() string {
	return string(iid)
}

// IsValid checks whether some string parsed as IterumID is a valid one
func (iid IterumID) IsValid() bool {
	return len(iid) == idLength
}

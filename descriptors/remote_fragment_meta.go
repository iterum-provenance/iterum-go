package descriptors

import "errors"

// RemoteMetadata is a structure used in remote fragment descriptions to add additional information
// Mostly used for Iterum internal functioning such as routing
type RemoteMetadata struct {
	FragmentID  IterumID               `json:"fragment_id"`
	TargetQueue *string                `json:"target_queue,omitempty"` // TargetQueue of the message queue
	Custom      map[string]interface{} `json:"custom,omitempty"`       // Custom metadata added by and for the user for across transformation steps
}

// Validate checks whether a metadata struct is valid (as far as possible)
func (rmd RemoteMetadata) Validate() (err error) {
	if !rmd.FragmentID.IsValid() {
		return errors.New("Invalid FragmentID in remote fragment")
	}
	return
}

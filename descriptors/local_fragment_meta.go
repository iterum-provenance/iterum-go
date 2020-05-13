package descriptors

import "errors"

// LocalMetadata is a structure used in local fragment descriptions to add additional information
// Mostly used for Iterum internal functioning such as routing
type LocalMetadata struct {
	FragmentID    IterumID               `json:"fragment_id"`
	OutputChannel *string                `json:"output_channel,omitempty"` // Target output of a transformation step
	Custom        map[string]interface{} `json:"custom,omitempty"`         // Custom metadata added by and for the user for across transformation steps
}

// Validate checks whether a metadata struct is valid (as far as possible)
func (lmd LocalMetadata) Validate() (err error) {
	if !lmd.FragmentID.IsValid() {
		return errors.New("Invalid FragmentID in local fragment")
	}
	return
}

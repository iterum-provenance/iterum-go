package descriptors

// RemoteMetadata is a structure used in remote fragment descriptions to add additional information
// Mostly used for Iterum internal functioning such as routing
type RemoteMetadata struct {
	TargetQueue *string                `json:"target_queue,omitempty"` // TargetQueue of the message queue
	Custom      map[string]interface{} `json:"custom,omitempty"`       // Custom metadata added by and for the user for across transformation steps
}

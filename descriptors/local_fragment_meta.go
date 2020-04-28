package descriptors

// LocalMetadata is a structure used in local fragment descriptions to add additional information
// Mostly used for Iterum internal functioning such as routing
type LocalMetadata struct {
	TargetQueue string `json:"target_queue"` // TargetQueue of the message queue
}

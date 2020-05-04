package descriptors

// LocalMetadata is a structure used in local fragment descriptions to add additional information
// Mostly used for Iterum internal functioning such as routing
type LocalMetadata struct {
	OutputChannel *string                `json:"output_channel,omitempty"` // Target output of a transformation step
	Custom        map[string]interface{} `json:"custom,omitempty"`         // Custom metadata added by and for the user for across transformation steps
}

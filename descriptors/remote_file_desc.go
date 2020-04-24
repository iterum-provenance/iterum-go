package descriptors

// RemoteFileDesc is a description of a file as found in the Minio storage
type RemoteFileDesc struct {
	Name       string `json:"name"`   // Original name as stored in the idv repository
	RemotePath string `json:"path"`   // Remote path to file within MinIO
	Bucket     string `json:"bucket"` // Name of the bucket that the file is stored in
}

// ToLocalPath converts a RemoteFileDesc into a path on the local disk on where to store it
func (rfd RemoteFileDesc) ToLocalPath(prefix string) string {
	return prefix + "/input/" + rfd.Bucket + "/" + rfd.Name
}

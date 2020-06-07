package minio

// ListConfigFiles returns the list of files contained in the config bucket of a minio config
func (cfg *Config) ListConfigFiles() []string {
	objects := cfg.Client.ListObjects(configBucket, "", true, make(chan struct{}, 1))
	results := []string{}
	for object := range objects {
		results = append(results, object.Key)
	}
	return results
}

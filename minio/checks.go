package minio

import (
	"fmt"
	"time"

	"github.com/minio/minio-go"
	"github.com/prometheus/common/log"
)

// Connect tries to initialize the Client element of a minio config
func (config *Config) Connect() error {
	client, err := minio.New(config.Endpoint, config.AccessKey, config.SecretKey, config.UseSSL)
	config.Client = client
	return err
}

// IsConnected returns whether the client of a MinioConfig is initialized
func (config Config) IsConnected() bool {
	return config.Client != nil
}

// EnsureBucket makes sure that the target bucket exists and rights are owned to it.
// It will retry for 'retries' times and return an error if it fails in the end
func (config Config) EnsureBucket(bucket string, retries int) (err error) {
	if !config.IsConnected() {
		return fmt.Errorf("Minio client not initialized, cannot check bucket existence")
	}
	// Check to see if we already own this bucket
	exists, errBucketExists := config.Client.BucketExists(bucket)
	if errBucketExists != nil {
		return fmt.Errorf("Failure of bucket existence checking: '%v'", errBucketExists)
	} else if !exists {
		log.Infof("Bucket '%v' does not exist, creating...\n", bucket)
		errMakeBucket := config.Client.MakeBucket(bucket, "")
		if errMakeBucket != nil {
			if retries > 0 { // retry a number of times
				time.Sleep(1 * time.Second)
				log.Infof("Failed to create bucket '%v', retrying...\n", bucket)
				return config.EnsureBucket(bucket, retries-1)
			}
			return fmt.Errorf("Failed to create bucket '%v' due to: '%v'", bucket, errMakeBucket)
		}
	}
	return
}

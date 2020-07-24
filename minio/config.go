package minio

import (
	minio "github.com/minio/minio-go"

	"github.com/iterum-provenance/iterum-go/process"
)

const (
	dataPrefix   string = "iterum_data"
	configPrefix string = "iterum_config"
)

var configBucket string = process.PipelineHash + "-iterum-config"

// Config is a structure holding all relevant information regarding the minio storage used by Iterum
type Config struct {
	TargetBucket string
	Endpoint     string
	AccessKey    string
	SecretKey    string
	UseSSL       bool
	PutOptions   minio.PutObjectOptions
	GetOptions   minio.GetObjectOptions
	Client       *minio.Client
}

// NewMinioConfig initiates a new minio configuration with all its necessary information
func NewMinioConfig(endpoint, accessKey, secretAccessKey, targetBucket string, useSSL bool) Config {
	putOptions := minio.PutObjectOptions{}
	getOptions := minio.GetObjectOptions{}
	return Config{
		targetBucket,
		endpoint,
		accessKey,
		secretAccessKey,
		useSSL,
		putOptions,
		getOptions,
		nil,
	}
}

// NewMinioConfigFromEnv uses environment variables to initialize a new MinioConfig configured for a step's output
func NewMinioConfigFromEnv() Config {
	// see env.go
	return NewMinioConfig(URL, AccessKey, SecretKey, TargetBucket, UseSSL)
}

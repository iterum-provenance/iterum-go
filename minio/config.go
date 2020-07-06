package minio

import (
	"strconv"

	"github.com/iterum-provenance/iterum-go/env"
	minio "github.com/minio/minio-go"
)

const (
	dataPrefix   string = "iterum_data"
	configPrefix string = "iterum_config"
)

var configBucket string = env.PipelineHash + "-iterum-config"

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
func NewMinioConfigFromEnv() (Config, error) {
	endpoint := env.MinioURL
	accessKeyID := env.MinioAccessKey
	secretAccessKey := env.MinioSecretKey
	useSSL, sslErr := strconv.ParseBool(env.MinioUseSSL)
	targetBucket := env.MinioTargetBucket
	return NewMinioConfig(endpoint, accessKeyID, secretAccessKey, targetBucket, useSSL), sslErr
}

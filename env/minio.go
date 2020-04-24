package env

import (
	"os"
)

const (
	urlEnv          = "MINIO_URL"
	accessKeyEnv    = "MINIO_ACCESS_KEY"
	secretKeyEnv    = "MINIO_SECRET_KEY"
	useSSLEnv       = "MINIO_USE_SSL"
	targetBucketEnv = "MINIO_OUTPUT_BUCKET"
)

// MinioURL is the url at which the minio client can be reached
var MinioURL = os.Getenv(urlEnv)

// MinioAccessKey is the access key for minio
var MinioAccessKey = os.Getenv(accessKeyEnv)

// MinioSecretKey is the secret access key for minio
var MinioSecretKey = os.Getenv(secretKeyEnv)

// MinioUseSSL is a string val denoting whether minio client uses SSL
var MinioUseSSL = os.Getenv(useSSLEnv)

// MinioTargetBucket is the bucket to which storage should go
var MinioTargetBucket = os.Getenv(targetBucketEnv)

// VerifyMinioEnvs checks whether each of the environment variables returned a non-empty value
func VerifyMinioEnvs() error {
	if MinioURL == "" {
		return ErrEnvironment(urlEnv, MinioURL)
	} else if MinioAccessKey == "" {
		return ErrEnvironment(accessKeyEnv, MinioAccessKey)
	} else if MinioSecretKey == "" {
		return ErrEnvironment(secretKeyEnv, MinioSecretKey)
	} else if MinioUseSSL == "" {
		return ErrEnvironment(useSSLEnv, MinioUseSSL)
	} else if MinioTargetBucket == "" {
		return ErrEnvironment(targetBucketEnv, MinioTargetBucket)
	}
	return nil
}

package env

import (
	"os"
)

const (
	minioURLEnv          = "MINIO_URL"
	minioAccessKeyEnv    = "MINIO_ACCESS_KEY"
	minioSecretKeyEnv    = "MINIO_SECRET_KEY"
	minioUseSSLEnv       = "MINIO_USE_SSL"
	minioTargetBucketEnv = "MINIO_OUTPUT_BUCKET"
)

// MinioURL is the url at which the minio client can be reached
var MinioURL = os.Getenv(minioURLEnv)

// MinioAccessKey is the access key for minio
var MinioAccessKey = os.Getenv(minioAccessKeyEnv)

// MinioSecretKey is the secret access key for minio
var MinioSecretKey = os.Getenv(minioSecretKeyEnv)

// MinioUseSSL is a string val denoting whether minio client uses SSL
var MinioUseSSL = os.Getenv(minioUseSSLEnv)

// MinioTargetBucket is the bucket to which storage should go
var MinioTargetBucket = os.Getenv(minioTargetBucketEnv)

// VerifyMinioEnvs checks whether each of the environment variables returned a non-empty value
func VerifyMinioEnvs() error {
	if MinioURL == "" {
		return ErrEnvironment(minioURLEnv, MinioURL)
	} else if MinioAccessKey == "" {
		return ErrEnvironment(minioAccessKeyEnv, MinioAccessKey)
	} else if MinioSecretKey == "" {
		return ErrEnvironment(minioSecretKeyEnv, MinioSecretKey)
	} else if MinioUseSSL == "" {
		return ErrEnvironment(minioUseSSLEnv, MinioUseSSL)
	} else if MinioTargetBucket == "" {
		return ErrEnvironment(minioTargetBucketEnv, MinioTargetBucket)
	}
	return nil
}

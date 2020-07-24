package minio

import (
	"fmt"
	"os"
	"strconv"

	"github.com/prometheus/common/log"

	"github.com/iterum-provenance/iterum-go/env"
	"github.com/iterum-provenance/iterum-go/process"
	"github.com/iterum-provenance/iterum-go/util"
)

func init() {
	minioEnvErr := VerifyMinioEnvs()
	iterumEnvErr := process.VerifyIterumEnvs() // necessary due to usage of env package throughout this package

	err := util.ReturnFirstErr(minioEnvErr, iterumEnvErr)
	if err != nil {
		log.Fatalln(err)
	}
}

const (
	minioURLEnv          = "MINIO_URL"
	minioAccessKeyEnv    = "MINIO_ACCESS_KEY"
	minioSecretKeyEnv    = "MINIO_SECRET_KEY"
	minioUseSSLEnv       = "MINIO_USE_SSL"
	minioTargetBucketEnv = "MINIO_OUTPUT_BUCKET"
)

// URL is the url at which the minio client can be reached
var URL = os.Getenv(minioURLEnv)

// AccessKey is the access key for minio
var AccessKey = os.Getenv(minioAccessKeyEnv)

// SecretKey is the secret access key for minio
var SecretKey = os.Getenv(minioSecretKeyEnv)

// UseSSL is a bool denoting whether minio client uses SSL
var UseSSL = parseUseSSL(minioUseSSLEnv)

// TargetBucket is the bucket to which storage should go
var TargetBucket = os.Getenv(minioTargetBucketEnv)

// VerifyMinioEnvs checks whether each of the environment variables returned a non-empty value
func VerifyMinioEnvs() error {
	if URL == "" {
		return env.ErrEnvironment(minioURLEnv, URL)
	} else if AccessKey == "" {
		return env.ErrEnvironment(minioAccessKeyEnv, AccessKey)
	} else if SecretKey == "" {
		return env.ErrEnvironment(minioSecretKeyEnv, SecretKey)
	} else if minioUseSSLEnv == "" {
		return env.ErrEnvironment(minioUseSSLEnv, fmt.Sprintf("%v", UseSSL))
	} else if TargetBucket == "" {
		return env.ErrEnvironment(minioTargetBucketEnv, TargetBucket)
	}
	return nil
}

func parseUseSSL(envName string) bool {
	value := os.Getenv(envName)
	var useSSL bool = false
	if value != "" {
		parsed, err := strconv.ParseBool(value)
		if err != nil {
			log.Fatalln(env.ErrEnvironment(envName, value))
		}
		useSSL = parsed
	}
	return useSSL
}

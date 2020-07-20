package process

import (
	"os"
	"path"

	"github.com/iterum-provenance/iterum-go/env"
	"github.com/prometheus/common/log"
)

func init() {
	err := VerifyIterumEnvs()

	if err != nil {
		log.Fatalln(err)
	}
}

const (
	dataVolumePathEnv = "DATA_VOLUME_PATH"
	nameEnv           = "ITERUM_NAME"
	configEnv         = "ITERUM_CONFIG"
	configPathEnv     = "ITERUM_CONFIG_PATH"
	pipelineHashEnv   = "PIPELINE_HASH"
)

// DataVolumePath is the path to the shared volume within this pod
var DataVolumePath = os.Getenv(dataVolumePathEnv)

// Name is the name (user defined) for this transformation step/fragmenter/etc
var Name = os.Getenv(nameEnv)

// Config contains a stringified JSON object containing config for the target (allowed to be empty)
var Config = os.Getenv(configEnv)

// ConfigPath contains a string/folder pointing to the folder where iterum config files should be stored (allowed to be empty)
var ConfigPath = path.Join(DataVolumePath, os.Getenv(configPathEnv))

// PipelineHash is the hash associated with this pipeline run
var PipelineHash = os.Getenv(pipelineHashEnv)

// VerifyIterumEnvs checks whether each of the environment variables returned a non-empty value
func VerifyIterumEnvs() error {
	if DataVolumePath == "" {
		return env.ErrEnvironment(dataVolumePathEnv, DataVolumePath)
	}
	if Name == "" {
		return env.ErrEnvironment(nameEnv, Name)
	}
	if PipelineHash == "" {
		return env.ErrEnvironment(pipelineHashEnv, PipelineHash)
	}
	if ConfigPath == DataVolumePath {
		return env.ErrEnvironment(configPathEnv, ConfigPath)
	}
	if Config == "" {
		log.Infof("environment variable %v was empty", configPathEnv)
	}
	if ConfigPath == "" {
		log.Infof("environment variable %v was empty", configPathEnv)
	}
	return nil
}

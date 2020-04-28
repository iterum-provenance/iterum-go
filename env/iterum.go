package env

import "os"

const (
	dataVolumePathEnv = "DATA_VOLUME_PATH"
	nameEnv           = "ITERUM_NAME"
	pipelineHashEnv   = "PIPELINE_HASH"
)

// DataVolumePath is the path to the shared volume within this pod
var DataVolumePath = os.Getenv(dataVolumePathEnv)

// ProcessName is the name (user defined) for this transformation step/fragmenter/etc
var ProcessName = os.Getenv(nameEnv)

// PipelineHash is th hash associated with this pipeline run
var PipelineHash = os.Getenv(pipelineHashEnv)

// VerifyIterumEnvs checks whether each of the environment variables returned a non-empty value
func VerifyIterumEnvs() error {
	if DataVolumePath == "" {
		return ErrEnvironment(dataVolumePathEnv, DataVolumePath)
	}
	if ProcessName == "" {
		return ErrEnvironment(nameEnv, ProcessName)
	}
	if PipelineHash == "" {
		return ErrEnvironment(pipelineHashEnv, PipelineHash)
	}
	return nil
}

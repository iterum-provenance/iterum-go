package env

import "os"

const (
	dataVolumePathEnv = "DATA_VOLUME_PATH"
	nameEnv           = "ITERUM_NAME"
	configEnv         = "ITERUM_CONFIG"
	pipelineHashEnv   = "PIPELINE_HASH"
	managerURLEnv     = "MANAGER_URL"
)

// DataVolumePath is the path to the shared volume within this pod
var DataVolumePath = os.Getenv(dataVolumePathEnv)

// ProcessName is the name (user defined) for this transformation step/fragmenter/etc
var ProcessName = os.Getenv(nameEnv)

// ProcessConfig contains a stringified JSON object containing config for the target (allowed to be empty)
var ProcessConfig = os.Getenv(configEnv)

// PipelineHash is the hash associated with this pipeline run
var PipelineHash = os.Getenv(pipelineHashEnv)

// ManagerURL is the url at which we can reach this pipeline's manager
var ManagerURL = os.Getenv(managerURLEnv)

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
	if ManagerURL == "" {
		return ErrEnvironment(managerURLEnv, ManagerURL)
	}
	return nil
}

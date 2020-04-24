package env

import "os"

const (
	dataVolumePathEnv = "DATA_VOLUME_PATH"
)

// DataVolumePath is the path to the shared volume within this pod
var DataVolumePath = os.Getenv(dataVolumePathEnv)

// VerifyIterumEnvs checks whether each of the environment variables returned a non-empty value
func VerifyIterumEnvs() error {
	if DataVolumePath == "" {
		return ErrEnvironment(dataVolumePathEnv, DataVolumePath)
	}
	return nil
}

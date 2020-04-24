package env

import "os"

const (
	urlEnv        = "DAEMON_URL"
	datasetEnv    = "DAEMON_DATASET"
	commitHashEnv = "DAEMON_COMMIT_HASH"
)

// DaemonURL is the URL at which to reach the idv daemon
var DaemonURL = os.Getenv(urlEnv)

// DaemonDataset is the idv dataset to use
var DaemonDataset = os.Getenv(datasetEnv)

// DaemonCommitHash is the hash of the idv dataset commit to use
var DaemonCommitHash = os.Getenv(commitHashEnv)

// VerifyDaemonEnvs checks whether each of the environment variables returned a non-empty value
func VerifyDaemonEnvs() error {
	if DaemonURL == "" {
		return ErrEnvironment(urlEnv, DaemonURL)
	} else if DaemonCommitHash == "" {
		return ErrEnvironment(datasetEnv, DaemonCommitHash)
	} else if DaemonDataset == "" {
		return ErrEnvironment(commitHashEnv, DaemonDataset)
	}
	return nil
}

package env

import "os"

const (
	daemonURLEnv        = "DAEMON_URL"
	daemonDatasetEnv    = "DAEMON_DATASET"
	daemonCommitHashEnv = "DAEMON_COMMIT_HASH"
)

// DaemonURL is the URL at which to reach the idv daemon
var DaemonURL = os.Getenv(daemonURLEnv)

// DaemonDataset is the idv dataset to use
var DaemonDataset = os.Getenv(daemonDatasetEnv)

// DaemonCommitHash is the hash of the idv dataset commit to use
var DaemonCommitHash = os.Getenv(daemonCommitHashEnv)

// VerifyDaemonEnvs checks whether each of the environment variables returned a non-empty value
func VerifyDaemonEnvs() error {
	if DaemonURL == "" {
		return ErrEnvironment(daemonURLEnv, DaemonURL)
	} else if DaemonCommitHash == "" {
		return ErrEnvironment(daemonDatasetEnv, DaemonCommitHash)
	} else if DaemonDataset == "" {
		return ErrEnvironment(daemonDatasetEnv, DaemonDataset)
	}
	return nil
}

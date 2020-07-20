package daemon

import (
	"log"
	"os"

	"github.com/iterum-provenance/iterum-go/env"
)

func init() {
	err := VerifyDaemonEnvs()
	if err != nil {
		log.Fatalln(err)
	}
}

const (
	daemonURLEnv        = "DAEMON_URL"
	daemonDatasetEnv    = "DAEMON_DATASET"
	daemonCommitHashEnv = "DAEMON_COMMIT_HASH"
)

// URL is the URL at which to reach the idv daemon
var URL = os.Getenv(daemonURLEnv)

// Dataset is the idv dataset to use
var Dataset = os.Getenv(daemonDatasetEnv)

// CommitHash is the hash of the idv dataset commit to use
var CommitHash = os.Getenv(daemonCommitHashEnv)

// VerifyDaemonEnvs checks whether each of the environment variables returned a non-empty value
func VerifyDaemonEnvs() error {
	if URL == "" {
		return env.ErrEnvironment(daemonURLEnv, URL)
	} else if CommitHash == "" {
		return env.ErrEnvironment(daemonDatasetEnv, CommitHash)
	} else if Dataset == "" {
		return env.ErrEnvironment(daemonDatasetEnv, Dataset)
	}
	return nil
}

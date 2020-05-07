package daemon

import "github.com/iterum-provenance/iterum-go/env"

// Config is a structure holding necessary info for interaction with an idv daemon
type Config struct {
	DaemonURL  string
	Dataset    string
	CommitHash string
}

// NewDaemonConfig initiates a new daemon configuration with all its necessary information
func NewDaemonConfig(daemonURL, dataset, commitHash string) Config {
	return Config{
		daemonURL,
		dataset,
		commitHash,
	}
}

// NewDaemonConfigFromEnv uses environment variables to initialize a new DaemonConfig
func NewDaemonConfigFromEnv() Config {
	daemonURL := env.DaemonURL
	dataset := env.DaemonDataset
	commitHash := env.DaemonCommitHash
	return NewDaemonConfig(daemonURL, dataset, commitHash)
}

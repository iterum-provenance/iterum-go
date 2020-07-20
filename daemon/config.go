package daemon

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
	// see env.go
	return NewDaemonConfig(URL, Dataset, CommitHash)
}

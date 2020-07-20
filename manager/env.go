package manager

import (
	"log"
	"os"

	"github.com/iterum-provenance/iterum-go/env"
)

func init() {
	err := VerifyManagerEnvs()
	if err != nil {
		log.Fatalln(err)
	}
}

const (
	urlEnv = "MANAGER_URL"
)

// URL is the url at which we can reach this pipeline's manager
var URL = os.Getenv(urlEnv)

// VerifyManagerEnvs checks whether each of the environment variables returned a non-empty value
func VerifyManagerEnvs() error {
	if URL == "" {
		return env.ErrEnvironment(urlEnv, URL)
	}
	return nil
}

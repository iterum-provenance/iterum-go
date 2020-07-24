package messageq

import (
	"os"
	"strconv"

	"github.com/prometheus/common/log"

	"github.com/iterum-provenance/iterum-go/env"
)

func init() {
	err := VerifyMessageQueueEnvs()
	if err != nil {
		log.Fatalln(err)
	}
}

const (
	mqBrokerURLEnv     = "MQ_BROKER_URL"
	mqOutputQueueEnv   = "MQ_OUTPUT_QUEUE"
	mqInputQueueEnv    = "MQ_INPUT_QUEUE"
	mqPrefetchCountEnv = "MQ_PREFETCH_COUNT"
)

// BrokerURL is the url at which we can reach the message queueing system
var BrokerURL = os.Getenv(mqBrokerURLEnv)

// OutputQueue is the queue into which we push the remote fragment descriptions
var OutputQueue = os.Getenv(mqOutputQueueEnv)

// InputQueue is the queue from which we pull the remote fragment descriptions
var InputQueue = os.Getenv(mqInputQueueEnv)

// PrefetchCount is the amount of messages a MQConsumer will prefetch: maps to rabbitMQ's QualityOfService.PrefetchCount variable
// will be -1 (<0) if the variable was not set
var PrefetchCount = parsePrefetchCount(mqPrefetchCountEnv)

// VerifyMessageQueueEnvs checks whether each of the environment variables returned a non-empty value
func VerifyMessageQueueEnvs() error {
	if BrokerURL == "" {
		return env.ErrEnvironment(mqBrokerURLEnv, BrokerURL)
	} else if OutputQueue == "" {
		return env.ErrEnvironment(mqOutputQueueEnv, OutputQueue)
	} else if InputQueue == "" {
		return env.ErrEnvironment(mqInputQueueEnv, InputQueue)
	}
	return nil
}

func parsePrefetchCount(envName string) int {
	value := os.Getenv(envName)
	var prefetchCount int = -1
	if value != "" {
		parsed, err := strconv.Atoi(value)
		if err != nil || parsed < 0 {
			log.Fatalln(env.ErrEnvironment(envName, value))
		}
		prefetchCount = parsed
	}
	return prefetchCount
}

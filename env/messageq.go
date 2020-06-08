package env

import (
	"os"
	"strconv"
)

const (
	mqBrokerURLEnv     = "MQ_BROKER_URL"
	mqOutputQueueEnv   = "MQ_OUTPUT_QUEUE"
	mqInputQueueEnv    = "MQ_INPUT_QUEUE"
	mqPrefetchCountEnv = "MQ_PREFETCH_COUNT"
)

// MQBrokerURL is the url at which we can reach the message queueing system
var MQBrokerURL = os.Getenv(mqBrokerURLEnv)

// MQOutputQueue is the queue into which we push the remote fragment descriptions
var MQOutputQueue = os.Getenv(mqOutputQueueEnv)

// MQInputQueue is the queue from which we pull the remote fragment descriptions
var MQInputQueue = os.Getenv(mqInputQueueEnv)

// MQPrefetchCount is the amount of messages a MQConsumer will prefetch: maps to rabbitMQ's QualityOfService.PrefetchCount variable
// will be -1 (<0) if the variable was not set
var MQPrefetchCount = parsePrefetchCount(mqPrefetchCountEnv)

// VerifyMessageQueueEnvs checks whether each of the environment variables returned a non-empty value
func VerifyMessageQueueEnvs() error {
	if MQBrokerURL == "" {
		return ErrEnvironment(mqBrokerURLEnv, MQBrokerURL)
	} else if MQOutputQueue == "" {
		return ErrEnvironment(mqOutputQueueEnv, MQOutputQueue)
	} else if MQInputQueue == "" {
		return ErrEnvironment(mqInputQueueEnv, MQInputQueue)
	}
	return nil
}

func parsePrefetchCount(envName string) int {
	value := os.Getenv(envName)
	var prefetchCount int = -1
	if value != "" {
		parsed, err := strconv.Atoi(value)
		if err != nil || parsed < 0 {
			ErrEnvironment(envName, value)
		}
		prefetchCount = parsed
	}
	return prefetchCount
}

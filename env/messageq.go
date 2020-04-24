package env

import "os"

const (
	mqBrokerURLEnv   = "MQ_BROKER_URL"
	mqOutputQueueEnv = "MQ_OUTPUT_QUEUE"
	mqInputQueueEnv  = "MQ_INPUT_QUEUE"
)

// MQBrokerURL is the url at which we can reach the message queueing system
var MQBrokerURL = os.Getenv(mqBrokerURLEnv)

// MQOutputQueue is the queue into which we push the remote fragment descriptions
var MQOutputQueue = os.Getenv(mqOutputQueueEnv)

// MQInputQueue is the queue from which we pull the remote fragment descriptions
var MQInputQueue = os.Getenv(mqInputQueueEnv)

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

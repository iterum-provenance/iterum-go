package env

import "os"

const (
	brokerURLEnv   = "MQ_BROKER_URL"
	outputQueueEnv = "MQ_OUTPUT_QUEUE"
	inputQueueEnv  = "MQ_INPUT_QUEUE"
)

// MQBrokerURL is the url at which we can reach the message queueing system
var MQBrokerURL = os.Getenv(brokerURLEnv)

// MQOutputQueue is the queue into which we push the remote fragment descriptions
var MQOutputQueue = os.Getenv(outputQueueEnv)

// MQInputQueue is the queue from which we pull the remote fragment descriptions
var MQInputQueue = os.Getenv(inputQueueEnv)

// VerifyMessageQueueEnvs checks whether each of the environment variables returned a non-empty value
func VerifyMessageQueueEnvs() error {
	if MQBrokerURL == "" {
		return ErrEnvironment(brokerURLEnv, MQBrokerURL)
	} else if MQOutputQueue == "" {
		return ErrEnvironment(outputQueueEnv, MQOutputQueue)
	} else if MQInputQueue == "" {
		return ErrEnvironment(inputQueueEnv, MQInputQueue)
	}
	return nil
}

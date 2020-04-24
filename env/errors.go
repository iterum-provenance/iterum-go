package env

import "fmt"

// EnvironmentError is raised when an environment variable that is verified turns up empty
type EnvironmentError struct {
	Variable string
	Value    string
}

func (err *EnvironmentError) Error() string {
	return fmt.Sprintf("EnvironmentError: %v with value '%v' is non-valid", err.Variable, err.Value)
}

func (err *EnvironmentError) Unwrap() error {
	return nil
}

// ErrEnvironment is used to raise a EnvironmentError
func ErrEnvironment(variable, value string) *EnvironmentError {
	return &EnvironmentError{Variable: variable, Value: value}
}

package transmit

import (
	"fmt"
)

// SerializationError is raised when serialize or deserialize fails
type SerializationError struct {
	Err error
}

func (err *SerializationError) Error() string {
	return fmt.Sprintf("SerializationError: %v", err.Err)
}

func (err *SerializationError) Unwrap() error {
	return err.Err
}

// ErrSerialization is used to raise a SerializationError embedding another error
func ErrSerialization(err error) *SerializationError { return &SerializationError{err} }

// ConnectionError is raised when something with net.Conn fails (such as reading or writing)
type ConnectionError struct {
	Err error
}

func (err ConnectionError) Error() string {
	return fmt.Sprintf("ConnectionError: %v", err.Err)
}

func (err *ConnectionError) Unwrap() error {
	return err.Err
}

// ErrConnection is used to raise a ConnectionError embedding another error
func ErrConnection(err error) *ConnectionError { return &ConnectionError{err} }

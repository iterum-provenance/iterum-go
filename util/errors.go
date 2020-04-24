package util

import (
	"errors"

	"github.com/prometheus/common/log"
)

// ReturnErrOnPanic catches panics, expects them to be of type error, then stores it in the pointer as recovery
func ReturnErrOnPanic(perr *error) func() {
	return func() {
		if r := recover(); r != nil {
			*perr = r.(error)
		}
	}
}

// Ensure checks for an error. If there is one it panics, if no error it prints the message (given it is not empty)
// `message` should describe what is ensured by the call and verification of `err` being `nil`
func Ensure(err error, message string) {
	if err != nil {
		panic(err)
	}
	if message != "" {
		log.Infoln(message)
	}
}

// PanicIfErr panics in case of an error. If custom message it creates a new error based on that and returns that instead
func PanicIfErr(err error, customMessage string) {
	if err != nil {
		if customMessage != "" {
			panic(errors.New(customMessage))
		} else {
			panic(err)
		}
	}
}

// ReturnFirstErr returns the first not-nil error from a list of errors.
// Used to prevent many copies of if err != nil { return err } when they are indepedent
func ReturnFirstErr(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

// LogIfError logs an error if it occurs, rather than repeating this structure everywhere
func LogIfError(err error) {
	if err != nil {
		log.Errorln(err)
	}
}

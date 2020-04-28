package util

// Parseable is an interface to support parsing from a file
type Parseable interface {
	ParseFromFile(filepath string) error
}

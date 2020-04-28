package util

// Validatable is a struct that can be checked for validity
type Validatable interface {
	IsValid() error
}

// Verify calls IsValid on the parameter and saves the first error in case of an error
// it returns the error in the passed variable, it ensures no if err != nil repetition in code
func Verify(v Validatable, err error) error {
	if err != nil {
		return err
	}
	return v.IsValid()
}

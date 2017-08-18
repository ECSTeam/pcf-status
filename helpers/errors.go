package helpers

import "fmt"

// Error used by the system.
type Error string

// NewErrorf create a new error from a format and parameters.
func NewErrorf(format string, params ...interface{}) (err Error) {
	return Error(fmt.Sprintf(format, params...))
}

// NewError create a new error from a message.
func NewError(message string) (err Error) {
	return Error(message)
}

// String value of the error.
func (err Error) String() string {
	return string(err)
}

// Error as a string
func (err Error) Error() string {
	return err.String()
}

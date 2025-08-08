package util

import "errors"

// ErrorAs checks if an error in err's chain matches target, and if so, sets target to the matching error.
// It is a generic wrapper around errors.As.
func ErrorAs[T error](err error) (asErr T, ok bool) {
	ok = errors.As(err, &asErr)
	return
}
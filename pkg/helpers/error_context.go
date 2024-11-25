package helpers

import "fmt"

// wrapError adds context to errors
func WrapError(err error, context string) error {
	return fmt.Errorf("%s: %v", context, err)
}

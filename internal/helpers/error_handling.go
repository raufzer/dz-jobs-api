package helpers

import (
	"errors"
	"fmt"
)
var (
	ErrEmailAlreadyExists  = errors.New("email already exists")
	ErrInvalidUserData     = errors.New("invalid user data")
	ErrUserCreationFailed  = errors.New("failed to create user")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrUserNotFound = errors.New("user not found")
	ErrTokenGeneration = errors.New("failed to generate token")
)

func WrapError(err error, context string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %v", context, err)
}
func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func IsErrorType(err error, targetErr error) bool {
	return errors.Is(err, targetErr)
}
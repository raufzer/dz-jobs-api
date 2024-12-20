package utils

func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

type CustomError struct {
	StatusCode int
	Message    string
}

func (c *CustomError) Error() string {
	panic("unimplemented")
}

func NewCustomError(statusCode int, message string) *CustomError {
	return &CustomError{
		StatusCode: statusCode,
		Message:    message,
	}
}

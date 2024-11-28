package helpers
import "fmt"

func WrapError(err error, context string) error {
	return fmt.Errorf("%s: %v", context, err)
}
func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}
	
package utils

import (
	"fmt"
	"os"
	"time"
)

func GetEnv(key string, valueType string) interface{} {
	value := os.Getenv(key)
	if value == "" {
		return fmt.Errorf("environment variable %s not set", key)
	}

	switch valueType {
	case "duration":
		duration, err := time.ParseDuration(value)
		if err != nil {
			return fmt.Sprintf("Failed to parse duration for environment variable %s: %v", key, err)
		}
		return duration
	case "string":
		return value
	default:
		return fmt.Sprintf("Unsupported value type %s", valueType)
	}
}

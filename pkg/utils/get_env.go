package utils

import (
	"fmt"
	"os"
	"time"
)

func GetEnv(key string, valueType string) interface{} {
	value := os.Getenv(key)
	if value == "" {
		fmt.Errorf("Environment variable %s not set", key)
	}

	switch valueType {
	case "duration":
		duration, err := time.ParseDuration(value)
		if err != nil {
			fmt.Errorf("Failed to parse duration for environment variable %s: %w", key, err)
		}
		return duration
	case "string":
		return value
	default:
		fmt.Errorf("Unsupported value type %s", valueType)
		return nil
	}
}

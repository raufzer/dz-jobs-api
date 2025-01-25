package utils

import (
	"fmt"
	"os"
	"time"
)

func GetEnv(key string, valueType string) (interface{}, error) {
	value := os.Getenv(key)
	if value == "" {
		return nil, fmt.Errorf("environment variable %s not set", key)
	}

	switch valueType {
	case "duration":
		duration, err := time.ParseDuration(value)
		if err != nil {
			return nil, fmt.Errorf("failed to parse duration for environment variable %s: %v", key, err)
		}
		return duration, nil
	case "string":
		return value, nil
	default:
		return nil, fmt.Errorf("unsupported value type %s", valueType)
	}
}

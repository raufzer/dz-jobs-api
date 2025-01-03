package utils

import (
	"log"
	"os"
	"time"
)

func GetEnv(key string, valueType string) interface{} {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is not set.", key)
	}

	switch valueType {
	case "duration":
		duration, err := time.ParseDuration(value)
		if err != nil {
			log.Fatalf("Error parsing duration from environment variable %s: %v", key, err)
		}
		return duration
	case "string":
		return value
	default:
		log.Fatalf("Unsupported value type %s for environment variable %s", valueType, key)
		return nil
	}
}

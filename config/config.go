package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort     string
	DatabaseURI    string
	JWTSecret      string
	TokenSecret    string
	TokenExpiresIn time.Duration
	TokenMaxAge    int
}

func LoadConfig() (*AppConfig, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Warning: No .env file found, using default environment variables.")
	}

	// Convert token expiration string to duration
	tokenExpStr := getEnvWithDefault("TOKEN_EXPIRES_IN", "24h")
	tokenExpDuration, err := time.ParseDuration(tokenExpStr)
	if err != nil {
		return nil, err
	}

	config := &AppConfig{
		ServerPort:     getEnvWithDefault("SERVER_PORT", "8080"),
		DatabaseURI:    getEnv("DATABASE_URI"),
		JWTSecret:      getEnv("JWT_SECRET"),
		TokenSecret:    getEnv("TOKEN_SECRET"),
		TokenExpiresIn: tokenExpDuration,
		TokenMaxAge:    int(tokenExpDuration.Hours()),
	}

	return config, nil
}

// getEnv gets an environment variable and panics if it's not set
func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is not set.", key)
	}
	return value
}

// getEnvWithDefault gets an environment variable or returns a default value if not set
func getEnvWithDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

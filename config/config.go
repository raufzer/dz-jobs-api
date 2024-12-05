package config

import (
	"dz-jobs-api/helpers"
	"dz-jobs-api/pkg/utils"
	"log"
	"time"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Domain        string
	ServerPort     string
	DatabaseURI    string
	JWTSecret      string
	TokenSecret    string
	TokenExpiresIn time.Duration
	TokenMaxAge    time.Duration
}

func LoadConfig() (*AppConfig, error) {

	err := godotenv.Load("./.env")
	if err != nil {
		log.Println("Warning: No .env file found, using default environment variables.")
	}

	tokenExpiresInStr := utils.GetEnv("TOKEN_EXPIRES_IN")
	tokenExpiresIn, err := time.ParseDuration(tokenExpiresInStr)
	if err != nil {
		return nil, helpers.WrapError(err, "parsing TOKEN_EXPIRES_IN")
	}
	tokenMaxAgeStr := utils.GetEnv("TOKEN_MAX_AGE")
	tokenMaxAge, err := time.ParseDuration(tokenMaxAgeStr)
	if err != nil {
		return nil, helpers.WrapError(err, "parsing TOKEN_MAX_AGE")
	}

	config := &AppConfig{
		Domain:     utils.GetEnv("DOMAIN"),
		ServerPort:     utils.GetEnv("SERVER_PORT"),
		DatabaseURI:    utils.GetEnv("DATABASE_URI"),
		JWTSecret:      utils.GetEnv("JWT_SECRET"),
		TokenSecret:    utils.GetEnv("TOKEN_SECRET"),
		TokenExpiresIn: tokenExpiresIn,
		TokenMaxAge:    tokenMaxAge,
	}

	return config, nil
}


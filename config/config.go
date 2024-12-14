package config

import (
	"dz-jobs-api/internal/helpers"
	"dz-jobs-api/pkg/utils"
	"log"
	"time"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Domain                   string
	ServerPort               string
	DatabaseURI              string
	RedisURI                 string
	RedisPassword            string
	SendGridAPIKey           string
	JWTSecret                string
	TokenSecret              string
	TokenExpiresIn           time.Duration
	AccessTokenMaxAge        time.Duration
	ResetPasswordTokenMaxAge time.Duration
}

func LoadConfig() (*AppConfig, error) {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Println("Warning: No .env file found, using default environment variables.")
	}
	tokenExpriresInStr := utils.GetEnv("TOKEN_EXPIRES_IN")
	tokenExpriresIn, err := time.ParseDuration(tokenExpriresInStr)
	if err != nil {
		return nil, helpers.NewCustomError(500, "parsing ACCESS_TOKEN_MAX_AGE")
	}
	accessTokenMaxAgeStr := utils.GetEnv("ACCESS_TOKEN_MAX_AGE")
	accessTokenMaxAge, err := time.ParseDuration(accessTokenMaxAgeStr)
	if err != nil {
		return nil, helpers.NewCustomError(500, "parsing ACCESS_TOKEN_MAX_AGE")
	}
	resetPasswordTokenMaxAgeStr := utils.GetEnv("RESET_PASSWORD_TOKEN_MAX_AGE")
	resetPasswordTokenMaxAge, err := time.ParseDuration(resetPasswordTokenMaxAgeStr)
	if err != nil {
		return nil, helpers.NewCustomError(500, "parsing RESET_PASSWORD_TOKEN_MAX_AGE")
	}
	config := &AppConfig{
		Domain:                   utils.GetEnv("DOMAIN"),
		ServerPort:               utils.GetEnv("SERVER_PORT"),
		DatabaseURI:              utils.GetEnv("DATABASE_URI"),
		RedisURI:                 utils.GetEnv("REDIS_URI"),
		RedisPassword:            utils.GetEnv("REDIS_PASSWORD"),
		SendGridAPIKey:           utils.GetEnv("SENDGRID_API_KEY"),
		JWTSecret:                utils.GetEnv("JWT_SECRET"),
		TokenSecret:              utils.GetEnv("TOKEN_SECRET"),
		TokenExpiresIn:           tokenExpriresIn,
		AccessTokenMaxAge:        accessTokenMaxAge,
		ResetPasswordTokenMaxAge: resetPasswordTokenMaxAge,
	}
	return config, nil
}

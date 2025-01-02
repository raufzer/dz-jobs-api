package config

import (
	"dz-jobs-api/pkg/utils"
	"log"
	"time"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Domain                   string
	FrontDomain              string
	ServerPort               string
	DatabaseURI              string
	RedisURI                 string
	RedisPassword            string
	SendGridAPIKey           string
	AccessTokenSecret        string
	RefreshTokenSecret       string
	ResetPasswordTokenSecret string
	AccessTokenMaxAge        time.Duration
	RefreshTokenMaxAge       time.Duration
	ResetPasswordTokenMaxAge time.Duration
	GoogleClientID           string
	GoogleClientSecret       string
	GoogleRedirectURL        string
	CloudinaryCloudName      string
	CloudinaryAPIKey         string
	CloudinaryAPISecret      string
	DefaultProfilePicture    string
	DefaultResume            string
	BuildVersion             string
	CommitHash               string
	Environment              string
	DocumentationURL         string
	LastMigration            string
}

func LoadConfig() (*AppConfig, error) {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Println("Warning: No .env file found, using default environment variables.")
	}
	accessTokenMaxAgeStr := utils.GetEnv("ACCESS_TOKEN_MAX_AGE")
	accessTokenMaxAge, err := time.ParseDuration(accessTokenMaxAgeStr)
	if err != nil {
		log.Println("Failed: getting token failed")
	}

	refreshTokenMaxAgeStr := utils.GetEnv("REFRESH_TOKEN_MAX_AGE")
	refreshTokenMaxAge, err := time.ParseDuration(refreshTokenMaxAgeStr)
	if err != nil {
		log.Println("Failed: getting token failed")
	}

	resetPasswordTokenMaxAgeStr := utils.GetEnv("RESET_PASSWORD_TOKEN_MAX_AGE")
	resetPasswordTokenMaxAge, err := time.ParseDuration(resetPasswordTokenMaxAgeStr)
	if err != nil {
		log.Println("Failed: getting token failed")
	}

	config := &AppConfig{
		Domain:                   utils.GetEnv("DOMAIN"),
		FrontDomain:              utils.GetEnv("FRONT_DOMAIN"),
		ServerPort:               utils.GetEnv("SERVER_PORT"),
		DatabaseURI:              utils.GetEnv("DATABASE_URI"),
		RedisURI:                 utils.GetEnv("REDIS_URI"),
		RedisPassword:            utils.GetEnv("REDIS_PASSWORD"),
		SendGridAPIKey:           utils.GetEnv("SENDGRID_API_KEY"),
		AccessTokenSecret:        utils.GetEnv("ACCESS_TOKEN_SECRET"),
		RefreshTokenSecret:       utils.GetEnv("REFRESH_TOKEN_SECRET"),
		ResetPasswordTokenSecret: utils.GetEnv("RESET_PASSWORD_TOKEN_SECRET"),
		AccessTokenMaxAge:        accessTokenMaxAge,
		RefreshTokenMaxAge:       refreshTokenMaxAge,
		ResetPasswordTokenMaxAge: resetPasswordTokenMaxAge,
		GoogleClientID:           utils.GetEnv("GOOGLE_CLIENT_ID"),
		GoogleClientSecret:       utils.GetEnv("GOOGLE_CLIENT_SECRET"),
		GoogleRedirectURL:        utils.GetEnv("GOOGLE_REDIRECT_URL"),
		CloudinaryCloudName:      utils.GetEnv("CLOUDINARY_CLOUD_NAME"),
		CloudinaryAPIKey:         utils.GetEnv("CLOUDINARY_API_KEY"),
		CloudinaryAPISecret:      utils.GetEnv("CLOUDINARY_API_SECRET"),
		DefaultProfilePicture:    utils.GetEnv("DEFAULT_PROFILE_PICTURE"),
		DefaultResume:            utils.GetEnv("DEFAULT_RESUME"),
		BuildVersion:             utils.GetEnv("BUILD_VERSION"),
		CommitHash:               utils.GetEnv("COMMIT_HASH"),
		Environment:              utils.GetEnv("ENVIRONMENT"),
		DocumentationURL:         utils.GetEnv("DOC_URL"),
		LastMigration:            utils.GetEnv("LAST_MIGRATION"),
	}
	return config, nil
}

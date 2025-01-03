package config

import (
	"dz-jobs-api/pkg/utils"
	"log"
	"time"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	BackEndDomain            string
	FrontEndDomain           string
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
	HealthURL                string
	VersionURL               string
	MetricsURL               string
	ServiceEmail             string
}

func LoadConfig() (*AppConfig, error) {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Println("Warning: No .env file found, using default environment variables.")
	}
	config := &AppConfig{
		BackEndDomain:            utils.GetEnv("BACK_END_DOMAIN", "string").(string),
		FrontEndDomain:           utils.GetEnv("FRONT_END_DOMAIN", "string").(string),
		ServerPort:               utils.GetEnv("SERVER_PORT", "string").(string),
		DatabaseURI:              utils.GetEnv("DATABASE_URI", "string").(string),
		RedisURI:                 utils.GetEnv("REDIS_URI", "string").(string),
		RedisPassword:            utils.GetEnv("REDIS_PASSWORD", "string").(string),
		SendGridAPIKey:           utils.GetEnv("SENDGRID_API_KEY", "string").(string),
		AccessTokenSecret:        utils.GetEnv("ACCESS_TOKEN_SECRET", "string").(string),
		RefreshTokenSecret:       utils.GetEnv("REFRESH_TOKEN_SECRET", "string").(string),
		ResetPasswordTokenSecret: utils.GetEnv("RESET_PASSWORD_TOKEN_SECRET", "string").(string),
		AccessTokenMaxAge:        utils.GetEnv("ACCESS_TOKEN_MAX_AGE", "duration").(time.Duration),
		RefreshTokenMaxAge:       utils.GetEnv("REFRESH_TOKEN_MAX_AGE", "duration").(time.Duration),
		ResetPasswordTokenMaxAge: utils.GetEnv("RESET_PASSWORD_TOKEN_MAX_AGE", "duration").(time.Duration),
		GoogleClientID:           utils.GetEnv("GOOGLE_CLIENT_ID", "string").(string),
		GoogleClientSecret:       utils.GetEnv("GOOGLE_CLIENT_SECRET", "string").(string),
		GoogleRedirectURL:        utils.GetEnv("GOOGLE_REDIRECT_URL", "string").(string),
		CloudinaryCloudName:      utils.GetEnv("CLOUDINARY_CLOUD_NAME", "string").(string),
		CloudinaryAPIKey:         utils.GetEnv("CLOUDINARY_API_KEY", "string").(string),
		CloudinaryAPISecret:      utils.GetEnv("CLOUDINARY_API_SECRET", "string").(string),
		DefaultProfilePicture:    utils.GetEnv("DEFAULT_PROFILE_PICTURE", "string").(string),
		DefaultResume:            utils.GetEnv("DEFAULT_RESUME", "string").(string),
		BuildVersion:             utils.GetEnv("BUILD_VERSION", "string").(string),
		CommitHash:               utils.GetEnv("COMMIT_HASH", "string").(string),
		Environment:              utils.GetEnv("ENVIRONMENT", "string").(string),
		DocumentationURL:         utils.GetEnv("DOC_URL", "string").(string),
		LastMigration:            utils.GetEnv("LAST_MIGRATION", "string").(string),
		HealthURL:                utils.GetEnv("HEALTH_URL", "string").(string),
		VersionURL:               utils.GetEnv("VERSION_URL", "string").(string),
		MetricsURL:               utils.GetEnv("METRICS_URL", "string").(string),
		ServiceEmail:             utils.GetEnv("SERVICE_EMAIL", "string").(string),
	}
	return config, nil
}

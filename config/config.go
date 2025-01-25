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
		BackEndDomain:            getEnvOrFatal("BACK_END_DOMAIN", "string").(string),
		FrontEndDomain:           getEnvOrFatal("FRONT_END_DOMAIN", "string").(string),
		ServerPort:               getEnvOrFatal("SERVER_PORT", "string").(string),
		DatabaseURI:              getEnvOrFatal("DATABASE_URI", "string").(string),
		RedisURI:                 getEnvOrFatal("REDIS_URI", "string").(string),
		RedisPassword:            getEnvOrFatal("REDIS_PASSWORD", "string").(string),
		SendGridAPIKey:           getEnvOrFatal("SENDGRID_API_KEY", "string").(string),
		AccessTokenSecret:        getEnvOrFatal("ACCESS_TOKEN_SECRET", "string").(string),
		RefreshTokenSecret:       getEnvOrFatal("REFRESH_TOKEN_SECRET", "string").(string),
		ResetPasswordTokenSecret: getEnvOrFatal("RESET_PASSWORD_TOKEN_SECRET", "string").(string),
		AccessTokenMaxAge:        getEnvOrFatal("ACCESS_TOKEN_MAX_AGE", "duration").(time.Duration),
		RefreshTokenMaxAge:       getEnvOrFatal("REFRESH_TOKEN_MAX_AGE", "duration").(time.Duration),
		ResetPasswordTokenMaxAge: getEnvOrFatal("RESET_PASSWORD_TOKEN_MAX_AGE", "duration").(time.Duration),
		GoogleClientID:           getEnvOrFatal("GOOGLE_CLIENT_ID", "string").(string),
		GoogleClientSecret:       getEnvOrFatal("GOOGLE_CLIENT_SECRET", "string").(string),
		GoogleRedirectURL:        getEnvOrFatal("GOOGLE_REDIRECT_URL", "string").(string),
		CloudinaryCloudName:      getEnvOrFatal("CLOUDINARY_CLOUD_NAME", "string").(string),
		CloudinaryAPIKey:         getEnvOrFatal("CLOUDINARY_API_KEY", "string").(string),
		CloudinaryAPISecret:      getEnvOrFatal("CLOUDINARY_API_SECRET", "string").(string),
		DefaultProfilePicture:    getEnvOrFatal("DEFAULT_PROFILE_PICTURE", "string").(string),
		DefaultResume:            getEnvOrFatal("DEFAULT_RESUME", "string").(string),
		BuildVersion:             getEnvOrFatal("BUILD_VERSION", "string").(string),
		CommitHash:               getEnvOrFatal("COMMIT_HASH", "string").(string),
		Environment:              getEnvOrFatal("ENVIRONMENT", "string").(string),
		DocumentationURL:         getEnvOrFatal("DOC_URL", "string").(string),
		LastMigration:            getEnvOrFatal("LAST_MIGRATION", "string").(string),
		HealthURL:                getEnvOrFatal("HEALTH_URL", "string").(string),
		VersionURL:               getEnvOrFatal("VERSION_URL", "string").(string),
		MetricsURL:               getEnvOrFatal("METRICS_URL", "string").(string),
		ServiceEmail:             getEnvOrFatal("SERVICE_EMAIL", "string").(string),
	}
	return config, nil
}

func getEnvOrFatal(key, expectedType string) interface{} {
	val := utils.GetEnv(key, expectedType)
	return val
}

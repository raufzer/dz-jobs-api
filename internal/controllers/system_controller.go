package controllers

import (
	"dz-jobs-api/config"
	"dz-jobs-api/internal/dto/response"
	"dz-jobs-api/pkg/utils"
	"golang.org/x/oauth2"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type SystemController struct {
	config      *config.AppConfig
	database    *config.DatabaseConfig
	redisClient *config.RedisConfig
}

func NewSystemController(config *config.AppConfig, db *config.DatabaseConfig, redis *config.RedisConfig) *SystemController {
	return &SystemController{
		config:      config,
		database:    db,
		redisClient: redis,
	}
}

func (c *SystemController) DefaultRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response.DefaultResponse{
		Message:       "Welcome to the DZ Jobs API",
		Version:       c.config.VersionURL,
		Health:        c.config.HealthURL,
		Documentation: c.config.DocumentationURL,
		Metrics:       c.config.MetricsURL,
	})
}

func (c *SystemController) GetAPIVersion(ctx *gin.Context) {
	loc, _ := time.LoadLocation("UTC")
	ctx.JSON(http.StatusOK, response.VersionResponse{
		APIVersion:       1,
		BuildVersion:     c.config.BuildVersion,
		CommitHash:       c.config.CommitHash,
		ReleaseDate:      time.Now().In(loc).Format("2006-01-02"),
		Environment:      c.config.Environment,
		DocumentationURL: c.config.DocumentationURL,
		LastMigration:    c.config.LastMigration,
	})
}

func (c *SystemController) GetHealth(ctx *gin.Context) {

	healthStatus := map[string]string{}

	if err := utils.CheckDatabaseHealth(c.database.DB); err != nil {
		healthStatus["database"] = "unhealthy"
	} else {
		healthStatus["database"] = "healthy"
	}

	if err := utils.CheckCacheHealth(c.redisClient.Client); err != nil {
		healthStatus["cache"] = "unhealthy"
	} else {
		healthStatus["cache"] = "healthy"
	}

	if err := utils.CheckSendGridHealth(c.config.SendGridAPIKey); err != nil {
		healthStatus["sendgrid"] = "unhealthy"
	} else {
		healthStatus["sendgrid"] = "healthy"
	}

	if err := utils.CheckCloudinaryHealth(c.config.CloudinaryCloudName, c.config.CloudinaryAPIKey, c.config.CloudinaryAPISecret); err != nil {
		healthStatus["cloudinary"] = "unhealthy"
	} else {
		healthStatus["cloudinary"] = "healthy"
	}

	oauthConfig := &oauth2.Config{
		ClientID:     c.config.GoogleClientID,
		ClientSecret: c.config.GoogleClientSecret,
		RedirectURL:  c.config.GoogleRedirectURL,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://accounts.google.com/o/oauth2/token",
		},
	}
	if err := utils.CheckGoogleOAuthHealth(oauthConfig); err != nil {
		healthStatus["google_oauth"] = "unhealthy"
	} else {
		healthStatus["google_oauth"] = "healthy"
	}

	healthStatus["status"] = response.AggregateHealthStatus(healthStatus)

	ctx.JSON(http.StatusOK, response.HealthResponse{
		Status:   healthStatus["status"],
		Database: healthStatus["database"],
		Cache:    healthStatus["cache"],
		ExternalServices: "SendGrid: " + healthStatus["sendgrid"] + ", " +
			"Cloudinary: " + healthStatus["cloudinary"] + ", " +
			"Google OAuth: " + healthStatus["google_oauth"],
	})
}

func (c *SystemController) GetMetrics(ctx *gin.Context) {
	uptime, requestCount, errorRate := utils.GetMetrics()
	ctx.JSON(http.StatusOK, response.MetricsResponse{
		Uptime:       uptime,
		RequestCount: requestCount,
		ErrorRate:    errorRate,
	})
}

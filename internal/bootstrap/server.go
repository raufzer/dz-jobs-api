package bootstrap

import (
	"dz-jobs-api/config"
	"dz-jobs-api/internal/middlewares"
	v1 "dz-jobs-api/internal/routes/api/v1"
	"dz-jobs-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

func CreateServer(appConfig *config.AppConfig) *gin.Engine {
	// Set Gin mode
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	// CORS setup
	server.Use(config.SetupCORS(appConfig.Domain))

	utils.InitLogger()
	// Global middleware
	server.Use(gin.Recovery())
	server.Use(middlewares.ErrorHandlingMiddleware())
	server.Use(middlewares.LoggingMiddleware())

	// Set-up Docs
	v1.RegisterSwaggerRoutes(server)
	return server
}

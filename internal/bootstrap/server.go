package bootstrap

import (
	"dz-jobs-api/config"
	"dz-jobs-api/internal/middlewares"

	v1 "dz-jobs-api/internal/routes/api/v1"

	"github.com/gin-gonic/gin"
)

func CreateServer(appConfig *config.AppConfig) *gin.Engine {
	// Set Gin mode
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	// CORS setup
	server.Use(config.SetupCORS(appConfig.FrontEndDomain, appConfig.BackEndDomain))

	// Global middleware
	server.Use(gin.Recovery())
	server.Use(middlewares.ErrorHandlingMiddleware())
	server.Use(middlewares.LoggingMiddleware())
	server.Use(middlewares.RateLimiter(20, 10))
	server.MaxMultipartMemory = 8 << 20 // 8 MiB

	// Set-up Docs
	v1.RegisterSwaggerRoutes(server)
	return server
}

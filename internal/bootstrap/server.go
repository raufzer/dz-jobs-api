package bootstrap

import (
	"dz-jobs-api/config"
	"dz-jobs-api/internal/middlewares"
	"dz-jobs-api/pkg/utils"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// CreateServer initializes and returns a configured Gin server instance.
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

	// Swagger setup
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL(appConfig.Domain+"/docs/doc.json"),
	))

	return server
}

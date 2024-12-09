package main

import (
	"dz-jobs-api/config"
	"dz-jobs-api/internal/bootstrap"
	"dz-jobs-api/internal/middlewares"
	v1 "dz-jobs-api/internal/routes/api/v1"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger documentation
	_ "dz-jobs-api/docs"
)

func main() {
	// Load configuration
	appConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize dependencies
	deps, err := bootstrap.InitializeDependencies(appConfig)
	if err != nil {
		log.Fatalf("Failed to initialize dependencies: %v", err)
	}

	// Gin setup
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	// CORS Configuration
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{appConfig.Domain}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Authorization", "Content-Type")
	server.Use(cors.New(corsConfig))

	// Global middleware
	server.Use(gin.Recovery())
	server.Use(middlewares.ErrorHandlingMiddleware())

	// Swagger setup
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL(appConfig.Domain+"/docs/doc.json"),
	))

	// API Routes
	basePath := server.Group("/v1")
	v1.RegisterRoutes(basePath, deps.UserController, deps.AuthController)

	// Server startup
	serverAddr := ":" + appConfig.ServerPort
	log.Printf("🚀 Server starting on %s", serverAddr)
	log.Fatal(server.Run(serverAddr))
}

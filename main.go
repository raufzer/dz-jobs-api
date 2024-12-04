package main

import (
	"dz-jobs-api/config"
	"dz-jobs-api/internal/controllers"
	"dz-jobs-api/internal/middleware"
	"dz-jobs-api/internal/repositories"
	"dz-jobs-api/internal/services"
	v1 "dz-jobs-api/routes/api/v1"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger documentation
	_ "dz-jobs-api/docs"
)

// @title           DzJobs API
// @version         1.0
// @description     Complete API for DzJobs Platform
// @host            dz-jobs-api-production.up.railway.app
// @BasePath        /v1
// @schemes         https
func main() {
	// Load configuration
	appConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Database connection
	dbConfig := config.ConnectDatabase(appConfig)

	// Validator
	validate := validator.New()

	// Repositories
	userRepo := repositories.NewUserRepository(dbConfig.DB)

	// Services
	authService := services.NewAuthServiceImpl(userRepo, validate)

	// Controllers
	userController := controllers.NewUserController(userRepo)
	authController := controllers.NewAuthController(authService, appConfig)

	// Gin setup
	gin.SetMode(gin.ReleaseMode) // Production mode
	server := gin.Default()

	// CORS Configuration
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{appConfig.Domain}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Authorization", "Content-Type")
	server.Use(cors.New(corsConfig))

	// Global middleware
	server.Use(gin.Recovery())
	server.Use(middleware.ErrorHandlingMiddleware()) // Add error handling middleware

	// Swagger setup
	// Alternative configuration
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL(appConfig.Domain+"/docs/doc.json"),
	))
	// API Routes
	basePath := server.Group("/v1")
	v1.UserRoutes(basePath, userController)
	v1.AuthRoutes(basePath, authController)

	// Server startup
	serverAddr := ":" + appConfig.ServerPort
	log.Printf("🚀 Server starting on %s", serverAddr)
	log.Fatal(server.Run(serverAddr))
}

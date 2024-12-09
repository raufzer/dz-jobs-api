package main

import (
	"dz-jobs-api/config"
	"dz-jobs-api/internal/bootstrap"
	"dz-jobs-api/internal/routes/api/v1"
	"log"
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

	// Create and configure Gin server
	server := bootstrap.CreateServer(appConfig)

	// API Routes
	basePath := server.Group("/v1")
	v1.RegisterRoutes(basePath, deps.UserController, deps.AuthController)

	// Server startup
	serverAddr := ":" + appConfig.ServerPort
	log.Printf("🚀 Server starting on %s", serverAddr)
	log.Fatal(server.Run(serverAddr))
}

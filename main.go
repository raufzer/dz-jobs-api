package main

import (
	"dz-jobs-api/config"
	"dz-jobs-api/internal/bootstrap"
	v1 "dz-jobs-api/internal/routes/api/v1"
	"log"
)

func main() {

	// Load app configuration
	appConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize dependencies
	deps, err := bootstrap.InitializeDependencies(appConfig)
	if err != nil {
		log.Fatalf("Failed to initialize dependencies: %v", err)
	}

	// Create server
	server := bootstrap.CreateServer(appConfig)

	// Register all routes (public, protected, role-based)
	v1.RegisterRoutes(server, deps.AuthController, deps.UserController, appConfig)

	// Start the server
	serverAddr := ":" + appConfig.ServerPort
	log.Printf("Server starting on %s", serverAddr)
	log.Fatal(server.Run(serverAddr))
}

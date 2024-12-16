package main

import (
	"dz-jobs-api/config"
	"dz-jobs-api/internal/bootstrap"
	"dz-jobs-api/internal/middlewares"
	"dz-jobs-api/internal/routes/api/v1"
	"log"
)

func main() {

	appConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	deps, err := bootstrap.InitializeDependencies(appConfig)
	if err != nil {
		log.Fatalf("Failed to initialize dependencies: %v", err)
	}

	server := bootstrap.CreateServer(appConfig)

	basePath := server.Group("/v1")
	v1.RegisterPublicRoutes(basePath, deps.AuthController)

	basePath.Use(middlewares.AuthMiddleware(appConfig)) // Apply AuthMiddleware only here
	v1.RegisterPrivateRoutes(basePath, deps.UserController)

	serverAddr := ":" + appConfig.ServerPort
	log.Printf("Server starting on %s", serverAddr)
	log.Fatal(server.Run(serverAddr))
}

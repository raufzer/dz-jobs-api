package main

import (
	"dz-jobs-api/config"
	"dz-jobs-api/internal/bootstrap"
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
	v1.RegisterRoutes(basePath, deps.UserController, deps.AuthController)

	serverAddr := ":" + appConfig.ServerPort
	log.Printf("Server starting on %s", serverAddr)
	log.Fatal(server.Run(serverAddr))
}

package bootstrap

import (
	"dz-jobs-api/config"
	"dz-jobs-api/internal/controllers"
	"dz-jobs-api/internal/repositories"
	"dz-jobs-api/internal/services"

	"github.com/go-playground/validator/v10"
)

type AppDependencies struct {
	UserController *controllers.UserController
	AuthController *controllers.AuthController
}

func InitializeDependencies(cfg *config.AppConfig) (*AppDependencies, error) { // Use 'AppConfig' type here
	// Database connection
	dbConfig := config.ConnectDatabase(cfg) // Pass 'cfg' directly

	// Validator
	validate := validator.New()

	// Repositories
	userRepo := repositories.NewUserRepository(dbConfig.DB)

	// Services
	authService := services.NewAuthServiceImpl(userRepo, validate)

	// Controllers
	userController := controllers.NewUserController(userRepo)
	authController := controllers.NewAuthController(authService, cfg)

	return &AppDependencies{
		UserController: userController,
		AuthController: authController,
	}, nil
}

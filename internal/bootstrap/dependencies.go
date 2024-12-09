package bootstrap

import (
	"dz-jobs-api/config"
	"dz-jobs-api/internal/controllers"
	"dz-jobs-api/internal/repositories"
	"dz-jobs-api/internal/services"
)

type AppDependencies struct {
	UserController *controllers.UserController
	AuthController *controllers.AuthController
}

func InitializeDependencies(cfg *config.AppConfig) (*AppDependencies, error) {

	dbConfig := config.ConnectDatabase(cfg)

	userRepo := repositories.NewUserRepository(dbConfig.DB)

	authService := services.NewAuthService(userRepo)
	userService := services.NewUserService(userRepo)

	userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(authService, cfg)

	return &AppDependencies{
		UserController: userController,
		AuthController: authController,
	}, nil
}

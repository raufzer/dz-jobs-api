package bootstrap

import (
	"dz-jobs-api/config"
	"dz-jobs-api/internal/controllers"
	"dz-jobs-api/internal/repositories/postgresql"
	"dz-jobs-api/internal/repositories/redis"
	"dz-jobs-api/internal/services"
)

type AppDependencies struct {
	UserController *controllers.UserController
	AuthController *controllers.AuthController
	RedisClient    *config.RedisConfig
}

func InitializeDependencies(cfg *config.AppConfig) (*AppDependencies, error) {
	// Initialize PostgreSQL Database
	dbConfig := config.ConnectDatabase(cfg)

	// Initialize Redis
	redisConfig := config.ConnectRedis(cfg)

	// Initialize Repositories
	userRepo := postgresql.NewUserRepository(dbConfig.DB)
	redisRepo := redis.NewUserRepository(redisConfig.Client)

	// Initialize Services
	authService := services.NewAuthService(
		userRepo,
		redisRepo, 
	)
	userService := services.NewUserService(userRepo)

	// Initialize Controllers
	userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(authService, cfg)

	// Return dependencies
	return &AppDependencies{
		UserController: userController,
		AuthController: authController,
		RedisClient:    redisConfig,
	}, nil
}

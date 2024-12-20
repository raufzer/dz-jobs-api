package bootstrap

import (
	"dz-jobs-api/config"
	"dz-jobs-api/internal/controllers"
	candidateControllers "dz-jobs-api/internal/controllers/candidate"
	"dz-jobs-api/internal/repositories/postgresql"
	candidatePostgresql "dz-jobs-api/internal/repositories/postgresql/candidate"
	"dz-jobs-api/internal/repositories/redis"
	"dz-jobs-api/internal/services"
	candidateServices "dz-jobs-api/internal/services/candidate"
)

type AppDependencies struct {
	UserController         *controllers.UserController
	AuthController         *controllers.AuthController
	RedisClient            *config.RedisConfig
	CandidateController    *candidateControllers.CandidateController
	PersonalInfoController *candidateControllers.CandidatePersonalInfoController
	EducationController    *candidateControllers.CandidateEducationController
	ExperienceController   *candidateControllers.CandidateExperienceController
}

func InitializeDependencies(cfg *config.AppConfig) (*AppDependencies, error) {
	// Initialize PostgreSQL Database
	dbConfig := config.ConnectDatabase(cfg)

	// Initialize Redis
	redisConfig := config.ConnectRedis(cfg)

	// Initialize Repositories
	userRepo := postgresql.NewUserRepository(dbConfig.DB)
	redisRepo := redis.NewRedisRepository(redisConfig.Client)
	candidateRepo := candidatePostgresql.NewCandidateRepository(dbConfig.DB)
	personalInfoRepo := candidatePostgresql.NewCandidatePersonalInfoRepository(dbConfig.DB)
	educationRepo := candidatePostgresql.NewCandidateEducationRepository(dbConfig.DB)
	experienceRepo := candidatePostgresql.NewCandidateExperienceRepository(dbConfig.DB)

	// Initialize Services
	authService := services.NewAuthService(
		userRepo,
		redisRepo,
		cfg,
	)
	userService := services.NewUserService(userRepo)
	candidateService := candidateServices.NewCandidateService(candidateRepo)
	personalInfoService := candidateServices.NewCandidatePersonalInfoService(personalInfoRepo)
	educationService := candidateServices.NewCandidateEducationService(educationRepo)
	experienceService := candidateServices.NewCandidateExperienceService(experienceRepo)

	// Initialize Controllers
	userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(authService, cfg)
	candidateController := candidateControllers.NewCandidateController(candidateService)
	personalInfoController := candidateControllers.NewCandidatePersonalInfoController(personalInfoService)
	educationController := candidateControllers.NewCandidateEducationController(educationService)
	experienceController := candidateControllers.NewCandidateExperienceController(experienceService)

	// Return dependencies
	return &AppDependencies{
		UserController:         userController,
		AuthController:         authController,
		RedisClient:            redisConfig,
		CandidateController:    candidateController,
		PersonalInfoController: personalInfoController,
		EducationController:    educationController,
		ExperienceController:   experienceController,
	}, nil
}

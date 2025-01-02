package bootstrap

import (
	"dz-jobs-api/config"
	"dz-jobs-api/internal/controllers"
	"dz-jobs-api/internal/integrations"
	"dz-jobs-api/internal/repositories/postgresql"
	"dz-jobs-api/internal/repositories/redis"
	"dz-jobs-api/internal/services"
	"dz-jobs-api/pkg/utils"
)

type AppDependencies struct {
	UserController           *controllers.UserController
	AuthController           *controllers.AuthController
	RedisClient              *config.RedisConfig
	CandidateController      *controllers.CandidateController
	PersonalInfoController   *controllers.CandidatePersonalInfoController
	EducationController      *controllers.CandidateEducationController
	ExperienceController     *controllers.CandidateExperienceController
	SkillsController         *controllers.CandidateSkillsController
	CertificationsController *controllers.CandidateCertificationsController
	PortfolioController      *controllers.CandidatePortfolioController
	RecruiterController      *controllers.RecruiterController
	JobController            *controllers.JobController
	BookmarksController      *controllers.BookmarksController
	SystemController         *controllers.SystemController
}

func InitializeDependencies(cfg *config.AppConfig) (*AppDependencies, error) {
	// Initialize PostgreSQL Database
	dbConfig := config.ConnectDatabase(cfg)

	// Initialize Redis
	redisConfig := config.ConnectRedis(cfg)

	// Initialize logger
	utils.InitLogger()

	// Initialize Cloudinary
	integrations.InitCloudinary(cfg)

	// Initialize Repositories
	userRepo := postgresql.NewUserRepository(dbConfig.DB)
	redisRepo := redis.NewRedisRepository(redisConfig.Client)
	candidateRepo := postgresql.NewCandidateRepository(dbConfig.DB)
	personalInfoRepo := postgresql.NewCandidatePersonalInfoRepository(dbConfig.DB)
	educationRepo := postgresql.NewCandidateEducationRepository(dbConfig.DB)
	experienceRepo := postgresql.NewCandidateExperienceRepository(dbConfig.DB)
	skillsRepo := postgresql.NewCandidateSkillsRepository(dbConfig.DB)
	certificationRepo := postgresql.NewCandidateCertificationsRepository(dbConfig.DB)
	portfolioRepo := postgresql.NewCandidatePortfolioRepository(dbConfig.DB)
	recruiterRepo := postgresql.NewRecruiterRepository(dbConfig.DB)
	jobRepo := postgresql.NewJobRepository(dbConfig.DB)
	bookmarksRepo := postgresql.NewBookmarskRepository(dbConfig.DB)

	// Initialize Services
	authService := services.NewAuthService(
		userRepo,
		redisRepo,
		cfg,
	)
	userService := services.NewUserService(userRepo)
	candidateService := services.NewCandidateService(candidateRepo, redisRepo, cfg)
	personalInfoService := services.NewCandidatePersonalInfoService(personalInfoRepo)
	educationService := services.NewCandidateEducationService(educationRepo, cfg)
	experienceService := services.NewCandidateExperienceService(experienceRepo)
	skillsService := services.NewCandidateSkillService(skillsRepo)
	certificationsService := services.NewCandidateCertificationsService(certificationRepo)
	portfolioService := services.NewCandidatePortfolioService(portfolioRepo)
	recruiterService := services.NewRecruiterService(recruiterRepo, redisRepo, cfg)
	jobService := services.NewJobService(jobRepo)
	bookmarksService := services.NewBookmarksService(bookmarksRepo)

	// Initialize Controllers
	userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(authService, cfg)
	candidateController := controllers.NewCandidateController(candidateService, cfg)
	personalInfoController := controllers.NewCandidatePersonalInfoController(personalInfoService)
	educationController := controllers.NewCandidateEducationController(educationService)
	experienceController := controllers.NewCandidateExperienceController(experienceService)
	skillsController := controllers.NewCandidateSkillsController(skillsService)
	certificationsController := controllers.NewCandidateCertificationsController(certificationsService)
	portfolioController := controllers.NewCandidatePortfolioController(portfolioService)
	recruiterController := controllers.NewRecruiterController(recruiterService)
	jobController := controllers.NewJobController(jobService)
	bookmarksController := controllers.NewBookmarksController(bookmarksService)
	systemController := controllers.NewSystemController(cfg, dbConfig, redisConfig)

	// Return dependencies
	return &AppDependencies{
		UserController:           userController,
		AuthController:           authController,
		RedisClient:              redisConfig,
		CandidateController:      candidateController,
		PersonalInfoController:   personalInfoController,
		EducationController:      educationController,
		ExperienceController:     experienceController,
		SkillsController:         skillsController,
		CertificationsController: certificationsController,
		PortfolioController:      portfolioController,
		RecruiterController:      recruiterController,
		JobController:            jobController,
		BookmarksController:      bookmarksController,
		SystemController:         systemController,
	}, nil
}

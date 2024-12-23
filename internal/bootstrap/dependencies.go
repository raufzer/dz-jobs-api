package bootstrap

import (
	"dz-jobs-api/config"
	"dz-jobs-api/internal/controllers"
	candidateControllers "dz-jobs-api/internal/controllers/candidate"
	"dz-jobs-api/internal/integrations"
	"dz-jobs-api/internal/repositories/postgresql"
	candidatePostgresql "dz-jobs-api/internal/repositories/postgresql/candidate"
	"dz-jobs-api/internal/repositories/redis"
	"dz-jobs-api/internal/services"
	candidateServices "dz-jobs-api/internal/services/candidate"
	"dz-jobs-api/pkg/utils"
)

type AppDependencies struct {
	UserController           *controllers.UserController
	AuthController           *controllers.AuthController
	RedisClient              *config.RedisConfig
	CandidateController      *candidateControllers.CandidateController
	PersonalInfoController   *candidateControllers.CandidatePersonalInfoController
	EducationController      *candidateControllers.CandidateEducationController
	ExperienceController     *candidateControllers.CandidateExperienceController
	SkillsController         *candidateControllers.CandidateSkillsController
	CertificationsController *candidateControllers.CandidateCertificationsController
	PortfolioController      *candidateControllers.CandidatePortfolioController
	RecruiterController      *controllers.RecruiterController
	JobController            *controllers.JobController
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
	candidateRepo := candidatePostgresql.NewCandidateRepository(dbConfig.DB)
	personalInfoRepo := candidatePostgresql.NewCandidatePersonalInfoRepository(dbConfig.DB)
	educationRepo := candidatePostgresql.NewCandidateEducationRepository(dbConfig.DB)
	experienceRepo := candidatePostgresql.NewCandidateExperienceRepository(dbConfig.DB)
	skillsRepo := candidatePostgresql.NewCandidateSkillsRepository(dbConfig.DB)
	certificationRepo := candidatePostgresql.NewCandidateCertificationsRepository(dbConfig.DB)
	portfolioRepo := candidatePostgresql.NewCandidatePortfolioRepository(dbConfig.DB)
	recruiterRepo := postgresql.NewRecruiterRepository(dbConfig.DB)
	jobRepo := postgresql.NewJobRepository(dbConfig.DB)

	// Initialize Services
	authService := services.NewAuthService(
		userRepo,
		redisRepo,
		cfg,
	)
	userService := services.NewUserService(userRepo)
	candidateService := candidateServices.NewCandidateService(candidateRepo, cfg)
	personalInfoService := candidateServices.NewCandidatePersonalInfoService(personalInfoRepo)
	educationService := candidateServices.NewCandidateEducationService(educationRepo, cfg)
	experienceService := candidateServices.NewCandidateExperienceService(experienceRepo)
	skillsService := candidateServices.NewCandidateSkillService(skillsRepo)
	certificationsService := candidateServices.NewCandidateCertificationsService(certificationRepo)
	portfolioService := candidateServices.NewCandidatePortfolioService(portfolioRepo)
	recruiterService := services.NewRecruiterService(recruiterRepo, cfg)
	jobService := services.NewJobService(jobRepo)

	// Initialize Controllers
	userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(authService, cfg)
	candidateController := candidateControllers.NewCandidateController(candidateService)
	personalInfoController := candidateControllers.NewCandidatePersonalInfoController(personalInfoService)
	educationController := candidateControllers.NewCandidateEducationController(educationService)
	experienceController := candidateControllers.NewCandidateExperienceController(experienceService)
	skillsController := candidateControllers.NewCandidateSkillsController(skillsService)
	certificationsController := candidateControllers.NewCandidateCertificationsController(certificationsService)
	portfolioController := candidateControllers.NewCandidatePortfolioController(portfolioService)
	recruiterController := controllers.NewRecruiterController(recruiterService)
	jobController := controllers.NewJobController(jobService)

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
	}, nil
}

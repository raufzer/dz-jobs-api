package v1

import (
	"dz-jobs-api/config"
	"dz-jobs-api/internal/controllers"
	"dz-jobs-api/internal/dto/response"
	"dz-jobs-api/internal/middlewares"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RegisterRoutes sets up all public and protected routes
func RegisterRoutes(
	router *gin.Engine,
	authController *controllers.AuthController,
	userController *controllers.UserController,
	recruiterController *controllers.RecruiterController,
	candidateController *controllers.CandidateController,
	personalInfoController *controllers.CandidatePersonalInfoController,
	educationController *controllers.CandidateEducationController,
	experienceController *controllers.CandidateExperienceController,
	skillsController *controllers.CandidateSkillsController,
	certificationsController *controllers.CandidateCertificationsController,
	portfolioController *controllers.CandidatePortfolioController,
	jobController *controllers.JobController,
	appConfig *config.AppConfig,
) {
	// Base path
	basePath := router.Group("/v1")

	// Public routes (no authentication required)
	RegisterPublicRoutes(basePath, authController, jobController)

	// Protected routes (authentication required)
	protected := basePath.Group("/")
	protected.Use(middlewares.AuthMiddleware(appConfig))
	RegisterProtectedRoutes(
		protected,
		userController,
		recruiterController,
		candidateController,
		personalInfoController,
		educationController,
		experienceController,
		skillsController,
		certificationsController,
		portfolioController,
		jobController,
	)
}

// RegisterPublicRoutes handles routes that don't require authentication
func RegisterPublicRoutes(
	router *gin.RouterGroup,
	authController *controllers.AuthController,
	jobController *controllers.JobController,
) {
	AuthRoutes(router, authController)
	JobRoutes(router, jobController)
}

// RegisterProtectedRoutes handles routes that require authentication
func RegisterProtectedRoutes(
	router *gin.RouterGroup,
	userController *controllers.UserController,
	recruiterController *controllers.RecruiterController,
	candidateController *controllers.CandidateController,
	personalInfoController *controllers.CandidatePersonalInfoController,
	educationController *controllers.CandidateEducationController,
	experienceController *controllers.CandidateExperienceController,
	skillsController *controllers.CandidateSkillsController,
	certificationsController *controllers.CandidateCertificationsController,
	portfolioController *controllers.CandidatePortfolioController,
	jobController *controllers.JobController,
) {
	// Admin-specific routes
	adminGroup := router.Group("/admin")
	adminGroup.Use(middlewares.RoleMiddleware("admin"))
	RegisterAdminRoutes(adminGroup, userController)

	// Candidate-specific routes
	candidateGroup := router.Group("/candidates")
	candidateGroup.Use(middlewares.RoleMiddleware("candidate", "admin"))
	RegisterCandidateRoutes(
		candidateGroup,
		candidateController,
		personalInfoController,
		educationController,
		experienceController,
		skillsController,
		certificationsController,
		portfolioController,
	)

	// Recruiter-specific routes
	recruiterGroup := router.Group("/recruiters")
	recruiterGroup.Use(middlewares.RoleMiddleware("recruiter", "admin"))
	RegisterRecruiterRoutes(recruiterGroup, recruiterController, jobController)
}

// RegisterAdminRoutes handles routes accessible only to admins
func RegisterAdminRoutes(
	router *gin.RouterGroup,
	userController *controllers.UserController,
) {
	UserRoutes(router, userController)
}

// RegisterCandidateRoutes handles routes accessible only to candidates
func RegisterCandidateRoutes(
	router *gin.RouterGroup,
	candidateController *controllers.CandidateController,
	personalInfoController *controllers.CandidatePersonalInfoController,
	educationController *controllers.CandidateEducationController,
	experienceController *controllers.CandidateExperienceController,
	skillsController *controllers.CandidateSkillsController,
	certificationsController *controllers.CandidateCertificationsController,
	portfolioController *controllers.CandidatePortfolioController,
) {
	candidateRoutes := router.Group("/:candidate_id")
	candidateRoutes.Use(middlewares.CandidateOwnershipMiddleware())
	CandidateRoutes(candidateRoutes, candidateController)
	PersonalInfoRoutes(candidateRoutes, personalInfoController)
	EducationRoutes(candidateRoutes, educationController)
	ExperienceRoutes(candidateRoutes, experienceController)
	SkillsRoutes(candidateRoutes, skillsController)
	CertificationsRoutes(candidateRoutes, certificationsController)
	PortfolioRoutes(candidateRoutes, portfolioController)
}

func RegisterRecruiterRoutes(
	router *gin.RouterGroup,
	recruiterController *controllers.RecruiterController,
	jobController *controllers.JobController,
) {
	recruiterRoutes := router.Group("/:recruiter_id")
	recruiterRoutes.Use(middlewares.RecruiterOwnershipMiddleware())
	RecruiterRoutes(recruiterRoutes, recruiterController)
	RecruiterJobRoutes(recruiterRoutes, jobController)
}

// RegisterSwaggerRoutes handles the Swagger documentation routes
func RegisterSwaggerRoutes(server *gin.Engine) {
	server.GET("/docs/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.URL("/v1/docs/swagger.json"),
	))

	server.GET("/v1/docs/swagger.json", func(ctx *gin.Context) {
		swaggerContent, err := os.ReadFile("./docs/swagger.json")
		if err != nil {
			ctx.JSON(http.StatusNotFound, response.Response{
				Code:    http.StatusNotFound,
				Status:  "Not Found",
				Message: "Swagger documentation not found",
			})
			return
		}
		ctx.Data(http.StatusOK, "application/json", swaggerContent)
	})
}

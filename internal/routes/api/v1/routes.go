package v1

import (
	"dz-jobs-api/config"
	"dz-jobs-api/internal/controllers"
	"os"

	"dz-jobs-api/internal/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

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
	bookmarksController *controllers.BookmarksController,
	systemController *controllers.SystemController,
	appConfig *config.AppConfig,
) {

	basePath := router.Group("/v1")

	RegisterPublicRoutes(basePath, authController, jobController, systemController)

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
		bookmarksController,
	)
}

func RegisterPublicRoutes(
	router *gin.RouterGroup,
	authController *controllers.AuthController,
	jobController *controllers.JobController,
	systemController *controllers.SystemController,
) {
	AuthRoutes(router, authController)
	JobRoutes(router, jobController)
	SystemRoutes(router, systemController)
}

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
	bookmarksController *controllers.BookmarksController,
) {

	adminGroup := router.Group("/admin")
	adminGroup.Use(middlewares.RoleMiddleware("admin"))
	RegisterAdminRoutes(adminGroup, userController)

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
		bookmarksController,
	)

	recruiterGroup := router.Group("/recruiters")
	recruiterGroup.Use(middlewares.RoleMiddleware("recruiter", "admin"))
	RegisterRecruiterRoutes(recruiterGroup, recruiterController, jobController)
}

func RegisterAdminRoutes(
	router *gin.RouterGroup,
	userController *controllers.UserController,
) {
	UserRoutes(router, userController)
}

func RegisterCandidateRoutes(
	router *gin.RouterGroup,
	candidateController *controllers.CandidateController,
	personalInfoController *controllers.CandidatePersonalInfoController,
	educationController *controllers.CandidateEducationController,
	experienceController *controllers.CandidateExperienceController,
	skillsController *controllers.CandidateSkillsController,
	certificationsController *controllers.CandidateCertificationsController,
	portfolioController *controllers.CandidatePortfolioController,
	bookmarksController *controllers.BookmarksController,
) {

	CandidateRoutes(router, candidateController)
	PersonalInfoRoutes(router, personalInfoController)
	EducationRoutes(router, educationController)
	ExperienceRoutes(router, experienceController)
	SkillsRoutes(router, skillsController)
	CertificationsRoutes(router, certificationsController)
	PortfolioRoutes(router, portfolioController)
	BookmarksRoute(router, bookmarksController)
}

func RegisterRecruiterRoutes(
	router *gin.RouterGroup,
	recruiterController *controllers.RecruiterController,
	jobController *controllers.JobController,
) {
	RecruiterRoutes(router, recruiterController)
	RecruiterJobRoutes(router, jobController)
}
func RegisterSwaggerRoutes(server *gin.Engine) {

	server.GET("/docs/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.URL("/v1/docs/swagger.json"),
	))

	server.GET("/v1/docs/swagger.json", func(ctx *gin.Context) {
		swaggerContent, _ := os.ReadFile("docs/swagger.json")
		ctx.Data(http.StatusOK, "application/json", swaggerContent)
	})
}

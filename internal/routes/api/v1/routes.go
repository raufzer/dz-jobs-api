package v1

import (
	"dz-jobs-api/config"
	"dz-jobs-api/internal/controllers"
	candidateControllers "dz-jobs-api/internal/controllers/candidate"
	"dz-jobs-api/internal/dto/response"
	"dz-jobs-api/internal/middlewares"
	v1 "dz-jobs-api/internal/routes/api/v1/candidate"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RegisterRoutes sets up all public and protected routes
func RegisterRoutes(router *gin.Engine, authController *controllers.AuthController, userController *controllers.UserController, candidateController *candidateControllers.CandidateController, personalInfoController *candidateControllers.CandidatePersonalInfoController, educationController *candidateControllers.CandidateEducationController, experienceController *candidateControllers.CandidateExperienceController, appConfig *config.AppConfig) {
	// Base path
	basePath := router.Group("/v1")

	// Public routes (no authentication required)
	RegisterPublicRoutes(basePath, authController)

	// Protected routes (authentication required)
	protected := basePath.Group("/")
	protected.Use(middlewares.AuthMiddleware(appConfig))
	RegisterProtectedRoutes(protected, userController, candidateController, personalInfoController, educationController, experienceController)
}

// RegisterPublicRoutes handles routes that don't require authentication
func RegisterPublicRoutes(router *gin.RouterGroup, authController *controllers.AuthController) {
	AuthRoutes(router, authController)
}

// RegisterProtectedRoutes handles routes that require authentication
func RegisterProtectedRoutes(router *gin.RouterGroup, userController *controllers.UserController, candidateController *candidateControllers.CandidateController, personalInfoController *candidateControllers.CandidatePersonalInfoController, educationController *candidateControllers.CandidateEducationController, experienceController *candidateControllers.CandidateExperienceController) {
	// Admin-specific routes
	adminGroup := router.Group("/admin")
	adminGroup.Use(middlewares.RoleMiddleware("admin"))
	RegisterAdminRoutes(adminGroup, userController)

	// Candidate-specific routes
	candidateGroup := router.Group("/candidates")
	candidateGroup.Use(middlewares.RoleMiddleware("candidate"))
	RegisterCandidateRoutes(candidateGroup, candidateController, personalInfoController, educationController, experienceController)

	// Recruiter-specific routes (Future Implementation)
	// recruiterGroup := router.Group("/recruiters")
	// recruiterGroup.Use(middlewares.RoleMiddleware("recruiter"))
	// RegisterRecruiterRoutes(recruiterGroup, recruiterController)
}

// RegisterAdminRoutes handles routes accessible only to admins
func RegisterAdminRoutes(router *gin.RouterGroup, userController *controllers.UserController) {
	UserRoutes(router, userController)
}

// RegisterCandidateRoutes handles routes accessible only to candidates
func RegisterCandidateRoutes(router *gin.RouterGroup, candidateController *candidateControllers.CandidateController, personalInfoController *candidateControllers.CandidatePersonalInfoController, educationController *candidateControllers.CandidateEducationController, experienceController *candidateControllers.CandidateExperienceController) {
	v1.CandidateRoutes(router, candidateController)
	v1.PersonalInfoRoutes(router, personalInfoController)
	v1.ExperienceRoutes(router, experienceController)
	v1.EducationRoutes(router, educationController)
}

// RegisterRecruiterRoutes handles routes accessible only to recruiters
// func RegisterRecruiterRoutes(router *gin.RouterGroup, recruiterController *controllers.RecruiterController) {
// 	router.POST("/job", recruiterController.PostJob)
// 	router.GET("/applications", recruiterController.GetApplications)
// }

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

// package v1

// import (
// 	controllers "dz-jobs-api/internal/controllers/candidate"
// 	"github.com/gin-gonic/gin"
// )

// func CandidateSkillsRoutes(rg *gin.RouterGroup, candidateSkillsController *controllers.CandidateSkillsController) {
// 	skillsRoute := rg.Group("/candidates/:candidate_id/skills")
// 	{
// 		skillsRoute.POST("/", candidateSkillsController.CreateSkill)
// 		skillsRoute.GET("/", candidateSkillsController.GetSkillsByID)
// 		skillsRoute.PUT("/", candidateSkillsController.UpdateSkill)
// 		skillsRoute.DELETE("/", candidateSkillsController.DeleteSkill)
// 	}
// }

// package v1

// import (
// 	controllers "dz-jobs-api/internal/controllers/candidate"
// 	"github.com/gin-gonic/gin"
// )

// func CandidateCertificationRoutes(rg *gin.RouterGroup, candidateCertificationController *controllers.CandidateCertificationController) {
// 	certificationRoute := rg.Group("/candidates/:candidate_id/certifications")
// 	{
// 		certificationRoute.POST("/", candidateCertificationController.CreateCertification)
// 		certificationRoute.GET("/", candidateCertificationController.GetCertificationByID)
// 		certificationRoute.PUT("/", candidateCertificationController.UpdateCertification)
// 		certificationRoute.DELETE("/", candidateCertificationController.DeleteCertification)
// 	}
// }

// package v1

// import (
// 	controllers "dz-jobs-api/internal/controllers/candidate"
// 	"github.com/gin-gonic/gin"
// )

// func CandidatePortfolioRoutes(rg *gin.RouterGroup, candidatePortfolioController *controllers.CandidatePortfolioController) {
// 	portfolioRoute := rg.Group("/candidates/:candidate_id/portfolio")
// 	{
// 		portfolioRoute.POST("/", candidatePortfolioController.CreatePortfolio)
// 		portfolioRoute.GET("/", candidatePortfolioController.GetPortfolioByID)
// 		portfolioRoute.PUT("/", candidatePortfolioController.UpdatePortfolio)
// 		portfolioRoute.DELETE("/", candidatePortfolioController.DeletePortfolio)
// 	}
// }

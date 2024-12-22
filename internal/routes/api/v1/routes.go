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
func RegisterRoutes(
    router *gin.Engine,
    authController *controllers.AuthController,
    userController *controllers.UserController,
    recruiterController *controllers.RecruiterController,
    candidateController *candidateControllers.CandidateController,
    personalInfoController *candidateControllers.CandidatePersonalInfoController,
    educationController *candidateControllers.CandidateEducationController,
    experienceController *candidateControllers.CandidateExperienceController,
    skillsController *candidateControllers.CandidateSkillsController,
    certificationsController *candidateControllers.CandidateCertificationsController,
    portfolioController *candidateControllers.CandidatePortfolioController,
    appConfig *config.AppConfig,
) {
    // Base path
    basePath := router.Group("/v1")

    // Public routes (no authentication required)
    RegisterPublicRoutes(basePath, authController)

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
    )
}

// RegisterPublicRoutes handles routes that don't require authentication
func RegisterPublicRoutes(
    router *gin.RouterGroup,
    authController *controllers.AuthController,
) {
    AuthRoutes(router, authController)
}

// RegisterProtectedRoutes handles routes that require authentication
func RegisterProtectedRoutes(
    router *gin.RouterGroup,
    userController *controllers.UserController,
    recruiterController *controllers.RecruiterController,
    candidateController *candidateControllers.CandidateController,
    personalInfoController *candidateControllers.CandidatePersonalInfoController,
    educationController *candidateControllers.CandidateEducationController,
    experienceController *candidateControllers.CandidateExperienceController,
    skillsController *candidateControllers.CandidateSkillsController,
    certificationsController *candidateControllers.CandidateCertificationsController,
    portfolioController *candidateControllers.CandidatePortfolioController,
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
    RegisterRecruiterRoutes(recruiterGroup, recruiterController)
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
    candidateController *candidateControllers.CandidateController,
    personalInfoController *candidateControllers.CandidatePersonalInfoController,
    educationController *candidateControllers.CandidateEducationController,
    experienceController *candidateControllers.CandidateExperienceController,
    skillsController *candidateControllers.CandidateSkillsController,
    certificationsController *candidateControllers.CandidateCertificationsController,
    portfolioController *candidateControllers.CandidatePortfolioController,
) {
    v1.CandidateRoutes(router, candidateController)
    v1.PersonalInfoRoutes(router, personalInfoController)
    v1.EducationRoutes(router, educationController)
    v1.ExperienceRoutes(router, experienceController)
    v1.SkillsRoutes(router, skillsController)
    v1.CertificationsRoutes(router, certificationsController)
    v1.PortfolioRoutes(router, portfolioController)
}

func RegisterRecruiterRoutes(
    router *gin.RouterGroup,
    recruiterController *controllers.RecruiterController,
) {
    RecruiterRoutes(router, recruiterController)
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
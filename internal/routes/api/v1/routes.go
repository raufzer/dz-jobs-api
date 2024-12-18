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
func RegisterRoutes(router *gin.Engine, authController *controllers.AuthController, userController *controllers.UserController, appConfig *config.AppConfig) {
	// Base path
	basePath := router.Group("/v1")

	// Public routes (no authentication required)
	RegisterPublicRoutes(basePath, authController)

	// Protected routes (authentication required)
	protected := basePath.Group("/")
	protected.Use(middlewares.AuthMiddleware(appConfig))
	RegisterProtectedRoutes(protected, userController)
}

// RegisterPublicRoutes handles routes that don't require authentication
func RegisterPublicRoutes(router *gin.RouterGroup, authController *controllers.AuthController) {
	AuthRoutes(router, authController)
}

// RegisterProtectedRoutes handles routes that require authentication
func RegisterProtectedRoutes(router *gin.RouterGroup, userController *controllers.UserController) {
	// Admin-specific routes
	adminGroup := router.Group("/admin")
	adminGroup.Use(middlewares.RoleMiddleware("admin"))
	RegisterAdminRoutes(adminGroup, userController)

	// Candidate-specific routes (Future Implementation)
	// candidateGroup := router.Group("/candidate")
	// candidateGroup.Use(middlewares.RoleMiddleware("candidate"))
	// RegisterCandidateRoutes(candidateGroup, candidateController)

	// Recruiter-specific routes (Future Implementation)
	// recruiterGroup := router.Group("/recruiter")
	// recruiterGroup.Use(middlewares.RoleMiddleware("recruiter"))
	// RegisterRecruiterRoutes(recruiterGroup, recruiterController)
}

// RegisterAdminRoutes handles routes accessible only to admins
func RegisterAdminRoutes(router *gin.RouterGroup, userController *controllers.UserController) {
	UserRoutes(router, userController)
}

// RegisterCandidateRoutes handles routes accessible only to candidates
// func RegisterCandidateRoutes(router *gin.RouterGroup, candidateController *controllers.CandidateController) {
// 	router.GET("/profile", candidateController.GetProfile)
// 	router.POST("/apply", candidateController.ApplyForJob)
// }

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

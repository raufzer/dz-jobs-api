package v1

import (
	"dz-jobs-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterPublicRoutes(router *gin.RouterGroup, authController *controllers.AuthController) {

	AuthRoutes(router, authController)
}

func RegisterPrivateRoutes(router *gin.RouterGroup, userController *controllers.UserController) {

	UserRoutes(router, userController)
}

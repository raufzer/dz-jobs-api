package v1

import (
	"dz-jobs-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, userController *controllers.UserController, authController *controllers.AuthController) {

	UserRoutes(router, userController)

	AuthRoutes(router, authController)
}

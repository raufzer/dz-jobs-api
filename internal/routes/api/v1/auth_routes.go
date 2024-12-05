package v1

import (
	"dz-jobs-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(rg *gin.RouterGroup, authController *controllers.AuthController) {
	rg.POST("/sessions", authController.Login)
	rg.POST("/users", authController.Register)
}

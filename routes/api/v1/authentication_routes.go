package v1

import (
	"dz-jobs-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func AuthenticationRoutes(rg *gin.RouterGroup, authenticationController *controllers.AuthenticationController) {
	authenticationRouter := rg.Group("/authentication")
	authenticationRouter.POST("/sessions ", authenticationController.Login)
	authenticationRouter.POST("/users", authenticationController.Register)
}

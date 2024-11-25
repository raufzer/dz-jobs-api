package v1

import (
	"dz-jobs-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup, userController *controllers.UserController) {
	userRoute := rg.Group("/user")
	userRoute.POST("/", userController.CreateUser)
	userRoute.GET("/:name", userController.GetUser)
	userRoute.GET("/", userController.GetAllUsers)
	userRoute.PATCH("/", userController.UpdateUser)
	userRoute.DELETE("/", userController.DeleteUser)
}

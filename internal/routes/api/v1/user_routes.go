package v1

import (
	"dz-jobs-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup, userController *controllers.UserController) {
	usersRoute := rg.Group("/users")
	usersRoute.POST("/", userController.CreateUser)
	usersRoute.GET("/", userController.GetAllUsers)
	usersRoute.GET("/:user_id", userController.GetUser)
	usersRoute.PATCH("/:user_id", userController.UpdateUser)
	usersRoute.DELETE("/:user_id", userController.DeleteUser)
}

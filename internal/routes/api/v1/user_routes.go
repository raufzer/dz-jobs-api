package v1

import (
	"dz-jobs-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup, userController *controllers.UserController) {
	usersRoute := rg.Group("/users")
	usersRoute.POST("/", userController.CreateUser)      
	usersRoute.GET("/", userController.GetAllUsers)      
	usersRoute.GET("/:id", userController.GetUser)       
	usersRoute.PATCH("/:id", userController.UpdateUser)  
	usersRoute.DELETE("/:id", userController.DeleteUser) 
}

package v1

import (
	"dz-jobs-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup, userController *controllers.UserController) {
	usersRoute := rg.Group("/users")
	usersRoute.POST("/", userController.CreateUser)      // Create a new user
	usersRoute.GET("/", userController.GetAllUsers)      // Retrieve all users with optional query params (e.g., pagination)
	usersRoute.GET("/:id", userController.GetUser)       // Retrieve a user by ID
	usersRoute.PATCH("/:id", userController.UpdateUser)  // Update user details by ID
	usersRoute.DELETE("/:id", userController.DeleteUser) // Delete a user by ID
}

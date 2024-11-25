package v1

import (
	"dz-jobs-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup, userController *controllers.UserController) {
	userRoute := rg.Group("/user")
	userRoute.POST("/create", userController.CreateUser)
	userRoute.GET("/get/:name", userController.GetUser)
	userRoute.GET("/getall", userController.GetAllUsers)
	userRoute.PATCH("/update", userController.UpdateUser)
	userRoute.DELETE("/delete", userController.DeleteUser)
}

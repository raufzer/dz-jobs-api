package v1

import (
	"dz-jobs-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RecruiterRoutes(rg *gin.RouterGroup, recruiterController *controllers.RecruiterController) {
	rg.POST("/", recruiterController.CreateRecruiter)
	rg.GET("/", recruiterController.GetRecruiter)
	rg.PUT("/", recruiterController.UpdateRecruiter)
	rg.DELETE("/", recruiterController.DeleteRecruiter)
}

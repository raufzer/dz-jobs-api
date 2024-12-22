package v1

import (
	"dz-jobs-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RecruiterRoutes(rg *gin.RouterGroup, recruiterController *controllers.RecruiterController) {

	rg.POST("/", recruiterController.CreateRecruiter)
	rg.GET("/:id", recruiterController.GetRecruiterByID)
	rg.PUT("/:id", recruiterController.UpdateRecruiter)
	rg.DELETE("/:id", recruiterController.DeleteRecruiter)
}

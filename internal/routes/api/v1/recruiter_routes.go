package v1

import (
	"dz-jobs-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RecruiterRoutes(rg *gin.RouterGroup, recruiterController *controllers.RecruiterController) {

	rg.POST("/", recruiterController.CreateRecruiter)
	rg.GET("/:recruiter_id", recruiterController.GetRecruiter)
	rg.PUT("/:recruiter_id", recruiterController.UpdateRecruiter)
	rg.DELETE("/:recruiter_id", recruiterController.DeleteRecruiter)
}

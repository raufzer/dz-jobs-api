package candidate

import (
	controllers "dz-jobs-api/internal/controllers/candidate"
	"github.com/gin-gonic/gin"
)

func EducationRoutes(rg *gin.RouterGroup, candidateEducationController *controllers.CandidateEducationController) {
	educationRoute := rg.Group("/candidates/:id/education")

	educationRoute.POST("/", candidateEducationController.CreateEducation)
	educationRoute.GET("/", candidateEducationController.GetEducationByID)
	// educationRoute.PUT("/", candidateEducationController.UpdateEducation)
	educationRoute.DELETE("/", candidateEducationController.DeleteEducation)

}

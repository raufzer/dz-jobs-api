package candidate

import (
	controllers "dz-jobs-api/internal/controllers/candidate"
	"github.com/gin-gonic/gin"
)

func EducationRoutes(rg *gin.RouterGroup, candidateEducationController *controllers.CandidateEducationController) {
	educationRoute := rg.Group("/:candidate_id/education")

	educationRoute.POST("/", candidateEducationController.AddEducation)
	educationRoute.GET("/", candidateEducationController.GetEducation)
	educationRoute.DELETE("/", candidateEducationController.DeleteEducation)

}

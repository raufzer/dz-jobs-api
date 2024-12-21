package candidate

import (
	controllers "dz-jobs-api/internal/controllers/candidate"
	"github.com/gin-gonic/gin"
)

func EducationRoutes(rg *gin.RouterGroup, candidateEducationController *controllers.CandidateEducationController) {
	educationRoute := rg.Group("/:id/education")

	educationRoute.POST("/", candidateEducationController.CreateEducation)
	educationRoute.GET("/", candidateEducationController.GetEducationByID)
	educationRoute.DELETE("/", candidateEducationController.DeleteEducation)

}

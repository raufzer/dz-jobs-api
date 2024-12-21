package candidate

import (
	controllers "dz-jobs-api/internal/controllers/candidate"
	"github.com/gin-gonic/gin"
)

func ExperienceRoutes(rg *gin.RouterGroup, candidateExperienceController *controllers.CandidateExperienceController) {
	experienceRoute := rg.Group("/:id/experience")
	
		experienceRoute.POST("/", candidateExperienceController.CreateExperience)
		experienceRoute.GET("/", candidateExperienceController.GetExperienceByID)
		experienceRoute.DELETE("/", candidateExperienceController.DeleteExperience)
	
}

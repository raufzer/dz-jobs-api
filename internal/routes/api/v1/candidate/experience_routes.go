package candidate

import (
	controllers "dz-jobs-api/internal/controllers/candidate"
	"github.com/gin-gonic/gin"
)

func ExperienceRoutes(rg *gin.RouterGroup, candidateExperienceController *controllers.CandidateExperienceController) {
	experienceRoute := rg.Group("/:candidate_id/experience")
	
		experienceRoute.POST("/", candidateExperienceController.AddExperience)
		experienceRoute.GET("/", candidateExperienceController.GetExperience)
		experienceRoute.DELETE("/", candidateExperienceController.DeleteExperience)
	
}

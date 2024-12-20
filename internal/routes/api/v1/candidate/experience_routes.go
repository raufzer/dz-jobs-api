package candidate

import (
	controllers "dz-jobs-api/internal/controllers/candidate"
	"github.com/gin-gonic/gin"
)

func ExperienceRoutes(rg *gin.RouterGroup, candidateExperienceController *controllers.CandidateExperienceController) {
	experienceRoute := rg.Group("/candidates/:id/experience")
	
		experienceRoute.POST("/", candidateExperienceController.CreateExperience)
		experienceRoute.GET("/", candidateExperienceController.GetExperienceByID)
		// experienceRoute.PUT("/", candidateExperienceController.UpdateExperience)
		experienceRoute.DELETE("/", candidateExperienceController.DeleteExperience)
	
}

package v1

import (
	"dz-jobs-api/internal/controllers"
	"github.com/gin-gonic/gin"
)

func ExperienceRoutes(rg *gin.RouterGroup, candidateExperienceController *controllers.CandidateExperienceController) {
	experienceRoute := rg.Group("/experience")
	experienceRoute.POST("/", candidateExperienceController.AddExperience)
	experienceRoute.GET("/", candidateExperienceController.GetExperience)
	experienceRoute.DELETE("/:experienceId", candidateExperienceController.DeleteExperience)

}

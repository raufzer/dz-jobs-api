package candidate

import (
	controllers "dz-jobs-api/internal/controllers/candidate"

	"github.com/gin-gonic/gin"
)

func SkillsRoutes(rg *gin.RouterGroup, candidateSkillsController *controllers.CandidateSkillsController) {
	skillsRoute := rg.Group("/:id/skills")

	skillsRoute.POST("/", candidateSkillsController.CreateSkill)
	skillsRoute.GET("/", candidateSkillsController.GetSkillsByID)
	// educationRoute.PUT("/", candidateEducationController.UpdateEducation)
	skillsRoute.DELETE("/:skill", candidateSkillsController.DeleteSkill)

}

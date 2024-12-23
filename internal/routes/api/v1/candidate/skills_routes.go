package candidate

import (
	controllers "dz-jobs-api/internal/controllers/candidate"

	"github.com/gin-gonic/gin"
)

func SkillsRoutes(rg *gin.RouterGroup, candidateSkillsController *controllers.CandidateSkillsController) {
	skillsRoute := rg.Group("/skills")

	skillsRoute.POST("/", candidateSkillsController.AddSkill)
	skillsRoute.GET("/", candidateSkillsController.GetSkills)
	skillsRoute.DELETE("/:skill_name", candidateSkillsController.DeleteSkill)

}

package candidate

import (
	request "dz-jobs-api/internal/dto/request/candidate"
	"dz-jobs-api/internal/dto/response"
	responseCandidate "dz-jobs-api/internal/dto/response/candidate"
	serviceInterfaces "dz-jobs-api/internal/services/interfaces/candidate"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CandidateSkillsController struct {
	service serviceInterfaces.CandidateSkillsService
}

func NewCandidateSkillsController(service serviceInterfaces.CandidateSkillsService) *CandidateSkillsController {
	return &CandidateSkillsController{service: service}
}

// AddSkill godoc
// @Summary Add a new skill
// @Description Add a new skill for a candidate by candidate ID
// @Tags Candidates - Skills
// @Accept json
// @Produce json
// @Param candidate_id path string true "Candidate ID"
// @Param Skill body request.AddSkillRequest true "Skill request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /candidates/{candidate_id}/skills [post]
func (c *CandidateSkillsController) AddSkill(ctx *gin.Context) {
	candidateID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	var req request.AddSkillRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	skill, err := c.service.AddSkill(candidateID, req)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, response.Response{
		Code:    http.StatusCreated,
		Status:  "Created",
		Message: "Skill created successfully",
		Data:    responseCandidate.ToSkillsResponse(skill),
	})
}

// GetSkills godoc
// @Summary Get skills
// @Description Get all skills for a candidate by candidate ID
// @Tags Candidates - Skills
// @Produce json
// @Param candidate_id path string true "Candidate ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /candidates/{candidate_id}/skills [get]
func (c *CandidateSkillsController) GetSkills(ctx *gin.Context) {
	candidateID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	skills, err := c.service.GetSkills(candidateID)
	if err != nil {
		ctx.Error(err)
		return
	}
	var skillResponses []responseCandidate.SkillResponse
	for _, skl := range skills {
		skillResponses = append(skillResponses, responseCandidate.ToSkillsResponse(&skl))
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Skills information retrieved successfully",
		Data:    skillResponses,
	})
}

// DeleteSkill godoc
// @Summary Delete skill
// @Description Delete a skill by candidate ID and skill name
// @Tags Candidates - Skills
// @Produce json
// @Param candidate_id path string true "Candidate ID"
// @Param skill_name path string true "Skill name"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /candidates/{candidate_id}/skills/{skill_name} [delete]
func (c *CandidateSkillsController) DeleteSkill(ctx *gin.Context) {
	candidateID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	err = c.service.DeleteSkill(candidateID, ctx.Param("skill"))
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Skill deleted successfully",
	})
}

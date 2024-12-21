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

// CreateSkill godoc
// @Summary Create a new skill
// @Description Create a new skill for a candidate
// @Tags Candidates, Skills_1create
// @Accept json
// @Produce json
// @Param id path string true "Candidate ID"
// @Param skill body request.AddSkillRequest true "Skill request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /candidates/{id}/skills [post]
func (c *CandidateSkillsController) CreateSkill(ctx *gin.Context) {
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

// GetSkillsByID godoc
// @Summary Get skills by candidate ID
// @Description Get all skills for a candidate by candidate ID
// @Tags Candidates, Skills_2get
// @Produce json
// @Param id path string true "Candidate ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /candidates/{id}/skills [get]
func (c *CandidateSkillsController) GetSkillsByID(ctx *gin.Context) {
	candidateID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	skills, err := c.service.GetSkillsByCandidateID(candidateID)
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
// @Tags Candidates, Skills_3delete
// @Produce json
// @Param id path string true "Candidate ID"
// @Param skill path string true "Skill name"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /candidates/{id}/skills/{skill} [delete]
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

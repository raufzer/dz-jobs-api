package controllers

import (
    "dz-jobs-api/internal/dto/request"
    "dz-jobs-api/internal/dto/response"
    serviceInterfaces "dz-jobs-api/internal/services/interfaces"
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
// @Param Skill body request.AddSkillRequest true "Skill request"
// @Success 201 {object} response.Response{Data=response.SkillResponse} "Skill created successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Candidate not found"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /candidates/skills [post]
func (c *CandidateSkillsController) AddSkill(ctx *gin.Context) {
	userID := ctx.MustGet("candidate_id")
	candidateID, _ := uuid.Parse(userID.(string))

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
		Data:    response.ToSkillResponse(skill),
	})
}

// GetSkills godoc
// @Summary Get skills
// @Description Get all skills for a candidate by candidate ID
// @Tags Candidates - Skills
// @Produce json
// @Success 200 {object} response.Response{Data=response.SkillsResponseData} "Skills retrieved successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Candidate not found"
// @Failure 404 {object} response.Response "Skills Info not found"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /candidates/skills [get]
func (c *CandidateSkillsController) GetSkills(ctx *gin.Context) {
	userID := ctx.MustGet("candidate_id")
	candidateID, _ := uuid.Parse(userID.(string))

	skills, err := c.service.GetSkills(candidateID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Skills information retrieved successfully",
		Data:   response.ToSkillsResponse(skills),
	})
}

// DeleteSkill godoc
// @Summary Delete skill
// @Description Delete a skill by candidate ID and skill name
// @Tags Candidates - Skills
// @Produce json
// @Param skill_name path string true "Skill name"
// @Success 200 {object} response.Response "Skill deleted successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Candidate not found"
// @Failure 404 {object} response.Response "Skill Info not found"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /candidates/skills/{skill_name} [delete]
func (c *CandidateSkillsController) DeleteSkill(ctx *gin.Context) {
	userID := ctx.MustGet("candidate_id")
	candidateID, _ := uuid.Parse(userID.(string))

	err := c.service.DeleteSkill(candidateID, ctx.Param("skill_name"))
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

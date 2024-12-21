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

	education, err := c.service.AddSkill(candidateID, req)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, response.Response{
		Code:    http.StatusCreated,
		Status:  "Created",
		Message: "Skill created successfully",
		Data:    responseCandidate.ToSkillsResponse(education),
	})
}

func (c *CandidateSkillsController) GetSkillsByID(ctx *gin.Context) {
	candidateID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	skill, err := c.service.GetSkillsByCandidateID(candidateID)
	if err != nil {
		ctx.Error(err)
		return
	}
	var skillResponses []responseCandidate.SkillResponse
	for _, skl := range skill {
		skillResponses = append(skillResponses, responseCandidate.ToSkillsResponse(&skl))
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Skills information retrieved successfully",
		Data:    skillResponses,
	})
}

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

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

type CandidateExperienceController struct {
	service serviceInterfaces.CandidateExperienceService
}

func NewCandidateExperienceController(service serviceInterfaces.CandidateExperienceService) *CandidateExperienceController {
	return &CandidateExperienceController{service: service}
}

func (c *CandidateExperienceController) CreateExperience(ctx *gin.Context) {
	candidateID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	var req request.AddExperienceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	experience, err := c.service.AddExperience(candidateID,req)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, response.Response{
		Code:    http.StatusCreated,
		Status:  "Created",
		Message: "Experience created successfully",
		Data:    responseCandidate.ToExperienceResponse(experience),
	})
}

func (c *CandidateExperienceController) GetExperienceByID(ctx *gin.Context) {
	candidateID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	experience, err := c.service.GetExperienceByCandidateID(candidateID)
	if err != nil {
		ctx.Error(err)
		return
	}
	var experienceResponses []responseCandidate.ExperienceResponse
	for _, exp := range experience {
		experienceResponses = append(experienceResponses, responseCandidate.ToExperienceResponse(&exp))
	}
	ctx.JSON(http.StatusCreated, response.Response{
		Code:    http.StatusCreated,
		Status:  "Created",
		Message: "Experience retrieved successfully",
		Data:    experienceResponses,
	})
}

func (c *CandidateExperienceController) DeleteExperience(ctx *gin.Context) {
	candidateID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	err = c.service.DeleteExperience(candidateID)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Experience deleted successfully",
	})
}

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

type CandidateEducationController struct {
	service serviceInterfaces.CandidateEducationService
}

func NewCandidateEducationController(service serviceInterfaces.CandidateEducationService) *CandidateEducationController {
	return &CandidateEducationController{service: service}
}

func (c *CandidateEducationController) CreateEducation(ctx *gin.Context) {
	candidateID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	var req request.AddEducationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	education, err := c.service.AddEducation(candidateID, req)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, response.Response{
		Code:    http.StatusCreated,
		Status:  "Created",
		Message: "Education created successfully",
		Data:    responseCandidate.ToEducationResponse(education),
	})
}

func (c *CandidateEducationController) GetEducationByID(ctx *gin.Context) {
	candidateID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	education, err := c.service.GetEducationByCandidateID(candidateID)
	if err != nil {
		ctx.Error(err)
		return
	}
	var educationResponses []responseCandidate.EducationResponse
	for _, edu := range education {
		educationResponses = append(educationResponses, responseCandidate.ToEducationResponse(&edu))
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Education information retrieved successfully",
		Data:    educationResponses,
	})
}

func (c *CandidateEducationController) DeleteEducation(ctx *gin.Context) {
	candidateID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	err = c.service.DeleteEducation(candidateID)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Education deleted successfully",
	})
}

package controllers

import (
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/dto/response"
	serviceInterfaces "dz-jobs-api/internal/services/interfaces"
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

// AddEducation godoc
// @Summary Add a new education record
// @Description Add a new education record for a candidate by candidate ID
// @Tags Candidates - Education
// @Accept json
// @Produce json
// @Param education body request.AddEducationRequest true "Education request"
// @Success 201 {object} response.Response{Data=response.EducationResponse} "Education created successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Candidate not found"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /candidates/education [post]
func (c *CandidateEducationController) AddEducation(ctx *gin.Context) {
	userID := ctx.MustGet("candidate_id")
	candidateID, _ := uuid.Parse(userID.(string))

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
		Data:    response.ToEducationResponse(education),
	})
}

// GetEducation godoc
// @Summary Get education records
// @Description Get all education records for a candidate by candidate ID
// @Tags Candidates - Education
// @Produce json
// @Success 200 {object} response.Response{Data=response.EducationsResponseData} "Education information retrieved successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Candidate not found"
// @Failure 404 {object} response.Response "Education not found"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /candidates/education [get]
func (c *CandidateEducationController) GetEducation(ctx *gin.Context) {
	userID := ctx.MustGet("candidate_id")
	candidateID, _ := uuid.Parse(userID.(string))

	educations, err := c.service.GetEducation(candidateID)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Education information retrieved successfully",
		Data:    response.ToEducationsResponse(educations),
	})
}

// DeleteEducation godoc
// @Summary Delete education record
// @Description Delete an education record by candidate ID and education ID
// @Tags Candidates - Education
// @Produce json
// @Param education_id path string true "Education ID"
// @Success 200 {object} response.Response "Education deleted successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Candidate not found"
// @Failure 404 {object} response.Response "Education not found"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /candidates/education/{education_id} [delete]
func (c *CandidateEducationController) DeleteEducation(ctx *gin.Context) {
	userID := ctx.MustGet("candidate_id")
	candidateID, _ := uuid.Parse(userID.(string))
	educationIDstr := ctx.Param("education_id")
    educationID, _ := uuid.Parse(educationIDstr)
	err := c.service.DeleteEducation(candidateID, educationID)
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

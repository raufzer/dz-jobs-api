package controllers

import (
    "dz-jobs-api/internal/dto/request"
    "dz-jobs-api/internal/dto/response"
    serviceInterfaces "dz-jobs-api/internal/services/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CandidatePersonalInfoController struct {
	service serviceInterfaces.CandidatePersonalInfoService
}

func NewCandidatePersonalInfoController(service serviceInterfaces.CandidatePersonalInfoService) *CandidatePersonalInfoController {
	return &CandidatePersonalInfoController{service: service}
}

// AddPersonalInfo godoc
// @Summary Add personal information
// @Description Add personal information for a candidate by candidate ID
// @Tags Candidates - Personal Info
// @Accept json
// @Produce json
// @Param personal_info body request.AddPersonalInfoRequest true "Personal Info request"
// @Success 201 {object} response.Response{Data=response.PersonalInfoResponse}
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Candidate not found"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /candidates/personal-info [post]
func (c *CandidatePersonalInfoController) AddPersonalInfo(ctx *gin.Context) {
	userID := ctx.MustGet("candidate_id")
	candidateID, _ := uuid.Parse(userID.(string))
	var req request.AddPersonalInfoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	createdInfo, err := c.service.AddPersonalInfo(req, candidateID)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, response.Response{
		Code:    http.StatusCreated,
		Status:  "Created",
		Message: "Personal information created successfully",
		Data:    response.ToPersonalInfoResponse(createdInfo),
	})
}

// GetPersonal godoc
// @Summary Get personal information
// @Description Get personal information for a candidate by candidate ID
// @Tags Candidates - Personal Info
// @Produce json
// @Success 200 {object} response.Response{Data=response.PersonalInfoResponse}
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Candidate not found"
// @Failure 404 {object} response.Response "Personal Info not found"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /candidates/personal-info [get]
func (c *CandidatePersonalInfoController) GetPersonalInfo(ctx *gin.Context) {
	userID := ctx.MustGet("candidate_id")
	candidateID, _ := uuid.Parse(userID.(string))

	info, err := c.service.GetPersonalInfo(candidateID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Personal information retrieved successfully",
		Data:    response.ToPersonalInfoResponse(info),
	})
}

// UpdatePersonalInfo godoc
// @Summary Update personal information
// @Description Update personal information for a candidate by candidate ID
// @Tags Candidates - Personal Info
// @Accept json
// @Produce json
// @Param personal_info body request.UpdatePersonalInfoRequest true "Personal Info request"
// @Success 200 {object} response.Response{Data=response.PersonalInfoResponse}
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Candidate not found"
// @Failure 404 {object} response.Response "Personal Info not found"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /candidates/personal-info [patch]
func (c *CandidatePersonalInfoController) UpdatePersonalInfo(ctx *gin.Context) {
	userID := ctx.MustGet("candidate_id")
	candidateID, _ := uuid.Parse(userID.(string))
	var req request.UpdatePersonalInfoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	updatedInfo, err := c.service.UpdatePersonalInfo(candidateID, req)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Personal information updated successfully",
		Data:    response.ToPersonalInfoResponse(updatedInfo),
	})
}

// DeletePersonalInfo godoc
// @Summary Delete personal information
// @Description Delete personal information for a candidate by candidate ID
// @Tags Candidates - Personal Info
// @Produce json
// @Success 200 {object} response.Response "Personal information deleted successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Candidate not found"
// @Failure 404 {object} response.Response "Personal Info not found"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /candidates/personal-info [delete]
func (c *CandidatePersonalInfoController) DeletePersonalInfo(ctx *gin.Context) {
	userID := ctx.MustGet("candidate_id")
	candidateID, _ := uuid.Parse(userID.(string))

	err :=  c.service.DeletePersonalInfo(candidateID)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Personal information deleted successfully",
	})
}

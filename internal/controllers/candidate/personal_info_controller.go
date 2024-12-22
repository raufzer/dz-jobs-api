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

type CandidatePersonalInfoController struct {
	service serviceInterfaces.CandidatePersonalInfoService
}

func NewCandidatePersonalInfoController(service serviceInterfaces.CandidatePersonalInfoService) *CandidatePersonalInfoController {
	return &CandidatePersonalInfoController{service: service}
}

// AddPersonalInfo godoc
// @Summary Add personal information
// @Description Add personal information for a candidate
// @Tags Candidates - Personal Info
// @Accept json
// @Produce json
// @Param candidate_id path string true "Candidate ID"
// @Param personal_info body request.AddPersonalInfoRequest true "Personal Info request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /candidates/{candidate_id}/personal-info [post]
func (c *CandidatePersonalInfoController) AddPersonalInfo(ctx *gin.Context) {
	candidateID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
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
		Data:    responseCandidate.ToPersonalInfoResponse(createdInfo),
	})
}

// GetPersonal godoc
// @Summary Get personal information by candidate ID
// @Description Get personal information for a candidate by candidate ID
// @Tags Candidates - Personal Info
// @Produce json
// @Param candidate_id path string true "Candidate ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /candidates/{candidate_id}/personal-info [get]
func (c *CandidatePersonalInfoController) GetPersonalInfo(ctx *gin.Context) {
	candidateID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	info, err := c.service.GetPersonalInfo(candidateID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Personal information retrieved successfully",
		Data:    responseCandidate.ToPersonalInfoResponse(info),
	})
}

// UpdatePersonalInfo godoc
// @Summary Update personal information
// @Description Update personal information for a candidate
// @Tags Candidates - Personal Info
// @Accept json
// @Produce json
// @Param candidate_id path string true "Candidate ID"
// @Param personal_info body request.UpdatePersonalInfoRequest true "Personal Info request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /candidates/{candidate_id}/personal-info [put]
func (c *CandidatePersonalInfoController) UpdatePersonalInfo(ctx *gin.Context) {
	candidateID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
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
		Data:    responseCandidate.ToPersonalInfoResponse(updatedInfo),
	})
}

// DeletePersonalInfo godoc
// @Summary Delete personal information
// @Description Delete personal information for a candidate
// @Tags Candidates - Personal Info
// @Produce json
// @Param candidate_id path string true "Candidate ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /candidates/{candidate_id}/personal-info [delete]
func (c *CandidatePersonalInfoController) DeletePersonalInfo(ctx *gin.Context) {
	candidateID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	err = c.service.DeletePersonalInfo(candidateID)
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

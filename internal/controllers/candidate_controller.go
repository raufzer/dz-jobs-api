package controllers

import (
	"dz-jobs-api/config"
	"dz-jobs-api/internal/dto/response"
	serviceInterfaces "dz-jobs-api/internal/services/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CandidateController struct {
	service serviceInterfaces.CandidateService
	config  *config.AppConfig
}

func NewCandidateController(service serviceInterfaces.CandidateService, config *config.AppConfig) *CandidateController {
	return &CandidateController{service: service, config: config}
}

// CreateCandidate godoc
// @Summary Create a new candidate
// @Description Create a new candidate with profile picture and resume
// @Tags Candidates - Candidate
// @Accept multipart/form-data
// @Produce json
// @Param profile_picture formData file true "Profile Picture"
// @Param resume formData file true "Resume"
// @Success 201 {object} response.Response{Data=response.CandidateResponse} "Candidate created successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /candidates [post]
func (c *CandidateController) CreateCandidate(ctx *gin.Context) {
	candidateID := ctx.MustGet("candidate_id")
	userID := candidateID.(string)
	profilePictureFile, err := ctx.FormFile("profile_picture")
	if err != nil {
		_  = ctx.Error(err)
		return
	}

	resumeFile, err := ctx.FormFile("resume")
	if err != nil {
		_  = ctx.Error(err)
		return
	}
	candidate, err := c.service.CreateCandidate(ctx,userID, profilePictureFile, resumeFile)
	if err != nil {
		_  = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, response.Response{
		Code:    http.StatusCreated,
		Status:  "Created",
		Message: "Candidate created successfully",
		Data:    response.ToCandidateResponse(candidate),
	})
}

// CreateCandidate godoc
// @Summary Create a new candidate
// @Description Create a new candidate with default profile picture and resume (when user skip profile setup process)
// @Tags Candidates - Candidate
// @Accept multipart/form-data
// @Produce json
// @Success 201 {object} response.Response{Data=response.CandidateResponse} "Candidate created successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /candidates/default [post]
func (c *CandidateController) CreateDefaultCandidate(ctx *gin.Context) {
	candidateID := ctx.MustGet("candidate_id")
	userID := candidateID.(string)
	defaultProfilePictureURL := c.config.DefaultProfilePicture
	defaultResumeURL := c.config.DefaultResume
	candidate, err := c.service.CreateDefaultCandidate(ctx,userID, defaultProfilePictureURL, defaultResumeURL)
	if err != nil {
		_  = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, response.Response{
		Code:    http.StatusCreated,
		Status:  "Created",
		Message: "Candidate created successfully",
		Data:    response.ToCandidateResponse(candidate),
	})
}

// GetCandidate godoc
// @Summary Get candidate
// @Description Get candidate details by ID
// @Tags Candidates - Candidate
// @Produce json
// @Success 200 {object} response.Response{Data=response.CandidateResponse} "Candidate found"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Candidate not found"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /candidates [get]
func (c *CandidateController) GetCandidate(ctx *gin.Context) {
	userID := ctx.MustGet("candidate_id")
	candidateID, err := uuid.Parse(userID.(string))
	if err != nil {
		_  = ctx.Error(err)
		return
	}
	candidate, err := c.service.GetCandidate(ctx,candidateID)
	if err != nil {
		_  = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Candidate found",
		Data:    response.ToCandidateResponse(candidate),
	})
}

// UpdateCandidate godoc
// @Summary Update candidate
// @Description Update candidate details with profile picture and resume  by ID
// @Tags Candidates - Candidate
// @Accept multipart/form-data
// @Produce json
// @Param profile_picture formData file true "Profile Picture"
// @Param resume formData file true "Resume"
// @Success 200 {object} response.Response{Data=response.CandidateResponse} "Candidate updated successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Candidate not found"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /candidates [put]
func (c *CandidateController) UpdateCandidate(ctx *gin.Context) {
	userID := ctx.MustGet("candidate_id")
	candidateID, err := uuid.Parse(userID.(string))
	if err != nil {
		_  = ctx.Error(err)
		return
	}
	profilePictureFile, err := ctx.FormFile("profile_picture")
	if err != nil {
		_  = ctx.Error(err)
		return
	}

	resumeFile, err := ctx.FormFile("resume")
	if err != nil {
		_  = ctx.Error(err)
		return
	}

	updatedCandidate, err := c.service.UpdateCandidate(ctx,candidateID, profilePictureFile, resumeFile)
	if err != nil {
		_  = ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Candidate updated successfully",
		Data:    response.ToCandidateResponse(updatedCandidate),
	})
}

// DeleteCandidate godoc
// @Summary Delete candidate
// @Description Delete candidate by ID
// @Tags Candidates - Candidate
// @Produce json
// @Success 200 {object} response.Response "Candidate deleted successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Candidate not found"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /candidates [delete]
func (c *CandidateController) DeleteCandidate(ctx *gin.Context) {
	userID := ctx.MustGet("candidate_id")
	candidateID, err := uuid.Parse(userID.(string))
	if err != nil {
		_  = ctx.Error(err)
		return
	}

	err = c.service.DeleteCandidate(ctx,candidateID)
	if err != nil {
		_  = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Candidate deleted successfully",
	})
}

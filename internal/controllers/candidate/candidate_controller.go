package candidate

import (
	"dz-jobs-api/internal/dto/response"
	responseCandidate "dz-jobs-api/internal/dto/response/candidate"
	serviceInterfaces "dz-jobs-api/internal/services/interfaces/candidate"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CandidateController struct {
	service serviceInterfaces.CandidateService
}

func NewCandidateController(service serviceInterfaces.CandidateService) *CandidateController {
	return &CandidateController{service: service}
}

// CreateCandidate godoc
// @Summary Create a new candidate
// @Description Create a new candidate with profile picture and resume
// @Tags Candidates - Candidate
// @Accept multipart/form-data
// @Produce json
// @Param profile_picture formData file true "Profile Picture"
// @Param resume formData file true "Resume"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /candidates [post]
func (c *CandidateController) CreateCandidate(ctx *gin.Context) {
	accessToken, err := ctx.Cookie("access_token")
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	userID, err := c.service.ExtractTokenDetails(accessToken)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	profilePictureFile, err := ctx.FormFile("profile_picture")
	if err != nil {
		ctx.Error(err)
		return
	}

	resumeFile, err := ctx.FormFile("resume")
	if err != nil {
		ctx.Error(err)
		return
	}
	candidate, err := c.service.CreateCandidate(userID, profilePictureFile, resumeFile)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, response.Response{
		Code:    http.StatusCreated,
		Status:  "Created",
		Message: "Candidate created successfully",
		Data:    responseCandidate.ToCandidateResponse(candidate),
	})
}

// GetCandidate godoc
// @Summary Get candidate
// @Description Get candidate details by ID
// @Tags Candidates - Candidate
// @Produce json
// @Param candidate_id path string true "Candidate ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /candidates/{candidate_id} [get]
func (c *CandidateController) GetCandidate(ctx *gin.Context) {
	candidateID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	candidate, err := c.service.GetCandidate(candidateID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, candidate)
}

// UpdateCandidate godoc
// @Summary Update candidate
// @Description Update candidate details with profile picture and resume  by ID
// @Tags Candidates - Candidate
// @Accept multipart/form-data
// @Produce json
// @Param candidate_id path string true "Candidate ID"
// @Param profile_picture formData file true "Profile Picture"
// @Param resume formData file true "Resume"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /candidates/{candidate_id} [put]
func (c *CandidateController) UpdateCandidate(ctx *gin.Context) {
	candidateID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	profilePictureFile, err := ctx.FormFile("profile_picture")
	if err != nil {
		ctx.Error(err)
		return
	}

	resumeFile, err := ctx.FormFile("resume")
	if err != nil {
		ctx.Error(err)
		return
	}

	updatedCandidate, err := c.service.UpdateCandidate(candidateID, profilePictureFile, resumeFile)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Candidate updated successfully",
		Data:    responseCandidate.ToCandidateResponse(updatedCandidate),
	})
}

// DeleteCandidate godoc
// @Summary Delete candidate
// @Description Delete candidate by ID
// @Tags Candidates - Candidate
// @Produce json
// @Param candidate_id path string true "Candidate ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /candidates/{candidate_id} [delete]
func (c *CandidateController) DeleteCandidate(ctx *gin.Context) {
	candidateID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	err = c.service.DeleteCandidate(candidateID)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Candidate deleted successfully",
	})
}

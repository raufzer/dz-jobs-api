package candidate

import (
	"dz-jobs-api/internal/dto/response"
	responseCandidate "dz-jobs-api/internal/dto/response/candidate"
	serviceInterfaces "dz-jobs-api/internal/services/interfaces/candidate"
	"fmt"
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

func (c *CandidateController) CreateCandidate(ctx *gin.Context) {
	// Log the request headers
	fmt.Println("Request Headers: ", ctx.Request.Header)
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

	candidate, err := c.service.CreateCandidate(profilePictureFile, resumeFile)
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

func (c *CandidateController) GetCandidateByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	candidateID, err := uuid.Parse(idParam)
	if err != nil {
		ctx.Error(err)
		return
	}

	candidate, err := c.service.GetCandidateByID(candidateID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, candidate)
}

func (c *CandidateController) UpdateCandidate(ctx *gin.Context) {

	idParam := ctx.Param("id")
	candidateID, err := uuid.Parse(idParam)
	if err != nil {
		ctx.Error(err)
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

func (c *CandidateController) DeleteCandidate(ctx *gin.Context) {
	idParam := ctx.Param("id")
	candidateID, err := uuid.Parse(idParam)
	if err != nil {
		ctx.Error(err)
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

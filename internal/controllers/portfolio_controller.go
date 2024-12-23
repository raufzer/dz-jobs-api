package controllers

import (
    "dz-jobs-api/internal/dto/request"
    "dz-jobs-api/internal/dto/response"
    serviceInterfaces "dz-jobs-api/internal/services/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CandidatePortfolioController struct {
	service serviceInterfaces.CandidatePortfolioService
}

func NewCandidatePortfolioController(service serviceInterfaces.CandidatePortfolioService) *CandidatePortfolioController {
	return &CandidatePortfolioController{service: service}
}

// AddProject godoc
// @Summary Add a new project
// @Description Add a new project for a candidate by candidate ID
// @Tags Candidates - Portfolio
// @Accept json
// @Produce json
// @Param candidate_id path string true "Candidate ID"
// @Param project body request.AddProjectRequest true "Project request"
// @Success 201 {object} response.Response{Data=response.PortfolioResponse} "Project created successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Candidate not found"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /candidates/{candidate_id}/portfolio [post]
func (c *CandidatePortfolioController) AddProject(ctx *gin.Context) {
	candidateID, err := uuid.Parse(ctx.Param("candidate_id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	var req request.AddProjectRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	project, err := c.service.AddProject(candidateID, req)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, response.Response{
		Code:    http.StatusCreated,
		Status:  "Created",
		Message: "Project created successfully",
		Data:    response.ToPortfolioResponse(project),
	})
}

// GetPortfolio godoc
// @Summary Get all projects (portfolio)
// @Description Get all projects for a candidate by candidate ID
// @Tags Candidates - Portfolio
// @Produce json
// @Param candidate_id path string true "Candidate ID"
// @Success 200 {object} response.Response{Data=response.PortfoliosResponseData} "Projects retrieved successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Candidate not found"
// @Failure 404 {object} response.Response "Projects Info not found"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /candidates/{candidate_id}/portfolio [get]
func (c *CandidatePortfolioController) GetPortfolio(ctx *gin.Context) {
	candidateID, err := uuid.Parse(ctx.Param("candidate_id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	projects, err := c.service.GetPortfolio(candidateID)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Projects retrieved successfully",
		Data:    response.ToPortfoliosResponse(projects),
	})
}

// DeleteProject godoc
// @Summary Delete project
// @Description Delete a project by candidate ID and project ID
// @Tags Candidates - Portfolio
// @Produce json
// @Param candidate_id path string true "Candidate ID"
// @Param project_id path string true "Project ID"
// @Success 200 {object} response.Response "Project deleted successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Candidate not found"
// @Failure 404 {object} response.Response "Project Info not found"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /candidates/{candidate_id}/portfolio/{project_id} [delete]
func (c *CandidatePortfolioController) DeleteProject(ctx *gin.Context) {
	candidateID, err := uuid.Parse(ctx.Param("candidate_id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	err = c.service.DeleteProject(candidateID, ctx.Param("project_id"))
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Project deleted successfully",
	})
}

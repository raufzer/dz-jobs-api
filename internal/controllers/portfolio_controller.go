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
// @Param project body request.AddProjectRequest true "Project request"
// @Success 201 {object} response.Response{Data=response.PortfolioResponse} "Project created successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Candidate not found"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /candidates/portfolio [post]
func (c *CandidatePortfolioController) AddProject(ctx *gin.Context) {
	userID := ctx.MustGet("candidate_id")
	candidateID, err := uuid.Parse(userID.(string))
	if err != nil {
		_  = ctx.Error(err)
		return
	}
	var req request.AddProjectRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_  = ctx.Error(err)
		ctx.Abort()
		return
	}

	project, err := c.service.AddProject(ctx,candidateID, req)
	if err != nil {
		_  = ctx.Error(err)
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
// @Success 200 {object} response.Response{Data=response.PortfoliosResponseData} "Projects retrieved successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Candidate not found"
// @Failure 404 {object} response.Response "Projects Info not found"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /candidates/portfolio [get]
func (c *CandidatePortfolioController) GetPortfolio(ctx *gin.Context) {
	userID := ctx.MustGet("candidate_id")
	candidateID, err := uuid.Parse(userID.(string))
	if err != nil {
		_  = ctx.Error(err)
		return
	}

	projects, err := c.service.GetPortfolio(ctx,candidateID)
	if err != nil {
		_  = ctx.Error(err)
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
// @Param projectId path string true "Project ID"
// @Success 200 {object} response.Response "Project deleted successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Candidate not found"
// @Failure 404 {object} response.Response "Project Info not found"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /candidates/portfolio/{projectId} [delete]
func (c *CandidatePortfolioController) DeleteProject(ctx *gin.Context) {
	userID := ctx.MustGet("candidate_id")
	candidateID, err := uuid.Parse(userID.(string))
	if err != nil {
		_  = ctx.Error(err)
		return
	}

	err = c.service.DeleteProject(ctx,candidateID, ctx.Param("projectId"))
	if err != nil {
		_  = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Project deleted successfully",
	})
}

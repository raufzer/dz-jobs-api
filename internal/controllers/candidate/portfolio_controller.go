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

type CandidatePortfolioController struct {
	service serviceInterfaces.CandidatePortfolioService
}

func NewCandidatePortfolioController(service serviceInterfaces.CandidatePortfolioService) *CandidatePortfolioController {
	return &CandidatePortfolioController{service: service}
}

// CreateProject godoc
// @Summary Create a new project
// @Description Create a new project for a candidate
// @Tags portfolio
// @Accept json
// @Produce json
// @Param id path string true "Candidate ID"
// @Param project body request.AddProjectRequest true "Project request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /candidates/{id}/portfolio [post]
func (c *CandidatePortfolioController) CreateProject(ctx *gin.Context) {
	candidateID, err := uuid.Parse(ctx.Param("id"))
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
		Data:    responseCandidate.ToPortfolioResponse(project),
	})
}

// GetProjectsByCandidateID godoc
// @Summary Get projects by candidate ID
// @Description Get all projects for a candidate by candidate ID
// @Tags portfolio
// @Produce json
// @Param id path string true "Candidate ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /candidates/{id}/portfolio [get]
func (c *CandidatePortfolioController) GetProjectsByCandidateID(ctx *gin.Context) {
	candidateID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	projects, err := c.service.GetPortfolioByCandidateID(candidateID)
	if err != nil {
		ctx.Error(err)
		return
	}
	var projectResponses []responseCandidate.PortfolioResponse
	for _, proj := range projects {
		projectResponses = append(projectResponses, responseCandidate.ToPortfolioResponse(&proj))
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Projects retrieved successfully",
		Data:    projectResponses,
	})
}

// DeleteProject godoc
// @Summary Delete project
// @Description Delete a project by candidate ID and project ID
// @Tags portfolio
// @Produce json
// @Param id path string true "Candidate ID"
// @Param project_id path string true "Project ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /candidates/{id}/portfolio/{project_id} [delete]
func (c *CandidatePortfolioController) DeleteProject(ctx *gin.Context) {
	candidateID, err := uuid.Parse(ctx.Param("id"))
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

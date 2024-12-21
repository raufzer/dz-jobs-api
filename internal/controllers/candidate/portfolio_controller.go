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

	portfolio, err := c.service.AddProject(candidateID, req)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, response.Response{
		Code:    http.StatusCreated,
		Status:  "Created",
		Message: "Project created successfully",
		Data:    responseCandidate.ToPortfolioResponse(portfolio),
	})
}

func (c *CandidatePortfolioController) GetPortfolio(ctx *gin.Context) {
	candidateID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	experience, err := c.service.GetPortfolioByCandidateID(candidateID)
	if err != nil {
		ctx.Error(err)
		return
	}
	var portfolioResponses []responseCandidate.PortfolioResponse
	for _, por := range experience {
		portfolioResponses = append(portfolioResponses, responseCandidate.ToPortfolioResponse(&por))
	}
	ctx.JSON(http.StatusCreated, response.Response{
		Code:    http.StatusCreated,
		Status:  "Created",
		Message: "Portfolio retrieved successfully",
		Data:    portfolioResponses,	
	})
}

func (c *CandidatePortfolioController) DeleteProject(ctx *gin.Context) {
	candidateID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	err = c.service.DeleteProject(candidateID, ctx.Param("portfolio_id"))
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

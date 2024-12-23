package v1

import (
	"dz-jobs-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func PortfolioRoutes(rg *gin.RouterGroup, candidatePortfolioController *controllers.CandidatePortfolioController) {
	portfolioRoute := rg.Group("/portfolio")
	portfolioRoute.POST("/", candidatePortfolioController.AddProject)
	portfolioRoute.GET("/", candidatePortfolioController.GetPortfolio)
	portfolioRoute.DELETE("/:project_id", candidatePortfolioController.DeleteProject)

}

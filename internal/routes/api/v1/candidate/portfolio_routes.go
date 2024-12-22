package candidate

import (
	controllers "dz-jobs-api/internal/controllers/candidate"

	"github.com/gin-gonic/gin"
)

func PortfolioRoutes(rg *gin.RouterGroup, candidatePortfolioController *controllers.CandidatePortfolioController) {
	portfolioRoute := rg.Group("/:candidate_id/portfolio")

	portfolioRoute.POST("/", candidatePortfolioController.AddProject)
	portfolioRoute.GET("/", candidatePortfolioController.GetPortfolio)
	portfolioRoute.DELETE("/:project_id", candidatePortfolioController.DeleteProject)

}

package candidate

import (
	controllers "dz-jobs-api/internal/controllers/candidate"

	"github.com/gin-gonic/gin"
)

func PortfolioRoutes(rg *gin.RouterGroup, candidatePortfolioController *controllers.CandidatePortfolioController) {
	portfolioRoute := rg.Group("/:id/portfolio")

	portfolioRoute.POST("/", candidatePortfolioController.CreateProject)
	portfolioRoute.GET("/", candidatePortfolioController.GetProjectsByCandidateID)
	// educationRoute.PUT("/", candidateEducationController.UpdateEducation)
	portfolioRoute.DELETE("/:project_id", candidatePortfolioController.DeleteProject)

}

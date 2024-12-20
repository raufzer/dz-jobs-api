package candidate

import (
	controllers "dz-jobs-api/internal/controllers/candidate"

	"github.com/gin-gonic/gin"
)

func CandidateRoutes(rg *gin.RouterGroup, candidateController *controllers.CandidateController) {
	candidateRoute := rg.Group("/candidates")

	candidateRoute.POST("/", candidateController.CreateCandidate)
	candidateRoute.GET("/:id", candidateController.GetCandidateByID)
	candidateRoute.PUT("/:id", candidateController.UpdateCandidate)
	candidateRoute.DELETE("/:id", candidateController.DeleteCandidate)

}

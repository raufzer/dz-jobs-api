package candidate

import (
	controllers "dz-jobs-api/internal/controllers/candidate"

	"github.com/gin-gonic/gin"
)

func CandidateRoutes(rg *gin.RouterGroup, candidateController *controllers.CandidateController) {

	rg.POST("/", candidateController.CreateCandidate)
	rg.GET("/:id", candidateController.GetCandidateByID)
	rg.PUT("/:id", candidateController.UpdateCandidate)
	rg.DELETE("/:id", candidateController.DeleteCandidate)

}

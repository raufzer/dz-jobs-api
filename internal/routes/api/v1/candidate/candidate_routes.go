package candidate

import (
	controllers "dz-jobs-api/internal/controllers/candidate"

	"github.com/gin-gonic/gin"
)

func CandidateRoutes(rg *gin.RouterGroup, candidateController *controllers.CandidateController) {

	rg.POST("/", candidateController.CreateCandidate)
	rg.GET("/:candidate_id", candidateController.GetCandidate)
	rg.PUT("/:candidate_id", candidateController.UpdateCandidate)
	rg.DELETE("/:candidate_id", candidateController.DeleteCandidate)

}

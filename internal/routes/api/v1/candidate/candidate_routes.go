package candidate

import (
	controllers "dz-jobs-api/internal/controllers/candidate"

	"github.com/gin-gonic/gin"
)

func CandidateRoutes(rg *gin.RouterGroup, candidateController *controllers.CandidateController) {

	rg.POST("/", candidateController.CreateCandidate)
	rg.GET("/", candidateController.GetCandidate)
	rg.PUT("/", candidateController.UpdateCandidate)
	rg.DELETE("/", candidateController.DeleteCandidate)

}

package v1

import (
	"dz-jobs-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func CandidateRoutes(rg *gin.RouterGroup, candidateController *controllers.CandidateController) {

	rg.POST("/", candidateController.CreateCandidate)
	rg.GET("/", candidateController.GetCandidate)
	rg.PUT("/", candidateController.UpdateCandidate)
	rg.DELETE("/", candidateController.DeleteCandidate)

}

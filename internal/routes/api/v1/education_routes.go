package v1

import (
	"dz-jobs-api/internal/controllers"
	"github.com/gin-gonic/gin"
)

func EducationRoutes(rg *gin.RouterGroup, candidateEducationController *controllers.CandidateEducationController) {
	educationRoute := rg.Group("/education")
	educationRoute.POST("/", candidateEducationController.AddEducation)
	educationRoute.GET("/", candidateEducationController.GetEducation)
	educationRoute.DELETE("/", candidateEducationController.DeleteEducation)

}

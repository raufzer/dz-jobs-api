package v1

import (
	"dz-jobs-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func PersonalInfoRoutes(rg *gin.RouterGroup, candidatePersonalInfoController *controllers.CandidatePersonalInfoController) {
	personalInfoRoute := rg.Group("/personal-info")

	personalInfoRoute.POST("/", candidatePersonalInfoController.AddPersonalInfo)
	personalInfoRoute.GET("/", candidatePersonalInfoController.GetPersonalInfo)
	personalInfoRoute.PUT("/", candidatePersonalInfoController.UpdatePersonalInfo)
	personalInfoRoute.DELETE("/", candidatePersonalInfoController.DeletePersonalInfo)

}

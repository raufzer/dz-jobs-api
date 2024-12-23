package candidate

import (
	controllers "dz-jobs-api/internal/controllers/candidate"

	"github.com/gin-gonic/gin"
)

func PersonalInfoRoutes(rg *gin.RouterGroup, candidatePersonalInfoController *controllers.CandidatePersonalInfoController) {
	personalInfoRoute := rg.Group("/personal-info")

	personalInfoRoute.POST("/", candidatePersonalInfoController.AddPersonalInfo)
	personalInfoRoute.GET("/", candidatePersonalInfoController.GetPersonalInfo)
	personalInfoRoute.PUT("/", candidatePersonalInfoController.UpdatePersonalInfo)
	personalInfoRoute.DELETE("/", candidatePersonalInfoController.DeletePersonalInfo)

}

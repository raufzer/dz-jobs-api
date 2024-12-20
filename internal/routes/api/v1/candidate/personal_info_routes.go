package candidate

import (
	controllers "dz-jobs-api/internal/controllers/candidate"
	"github.com/gin-gonic/gin"
)

func PersonalInfoRoutes(rg *gin.RouterGroup, candidatePersonalInfoController *controllers.CandidatePersonalInfoController) {
	personalInfoRoute := rg.Group("/:id/personal_info")
	
		// personalInfoRoute.POST("/", candidatePersonalInfoControlller.CreatePersonalInfo)
		personalInfoRoute.GET("/", candidatePersonalInfoController.GetPersonalInfoByID)
		personalInfoRoute.PUT("/", candidatePersonalInfoController.UpdatePersonalInfo)
		// personalInfoRoute.DELETE("/", candidatePersonalInfoController.DeletePersonalInfo)
	
}

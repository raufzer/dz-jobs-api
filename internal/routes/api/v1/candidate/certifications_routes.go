package candidate

import (
	controllers "dz-jobs-api/internal/controllers/candidate"

	"github.com/gin-gonic/gin"
)

func CertificationsRoutes(rg *gin.RouterGroup, candidateCertificationsController *controllers.CandidateCertificationsController) {
	certificationsRoute := rg.Group("/:id/certifications")

	certificationsRoute.POST("/", candidateCertificationsController.CreateCertification)
	certificationsRoute.GET("/", candidateCertificationsController.GetCertfications)
	// educationRoute.PUT("/", candidateEducationController.UpdateEducation)
	certificationsRoute.DELETE("/:certification", candidateCertificationsController.DeleteCertification)

}

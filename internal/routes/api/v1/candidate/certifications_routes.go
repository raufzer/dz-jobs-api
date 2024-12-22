package candidate

import (
	controllers "dz-jobs-api/internal/controllers/candidate"

	"github.com/gin-gonic/gin"
)

func CertificationsRoutes(rg *gin.RouterGroup, candidateCertificationsController *controllers.CandidateCertificationsController) {
	certificationsRoute := rg.Group("/:candidate_id/certifications")

	certificationsRoute.POST("/", candidateCertificationsController.AddCertification)
	certificationsRoute.GET("/", candidateCertificationsController.GetCertifications)
	certificationsRoute.DELETE("/:certification_id", candidateCertificationsController.DeleteCertification)

}

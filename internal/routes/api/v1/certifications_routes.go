package v1

import (
	"dz-jobs-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func CertificationsRoutes(rg *gin.RouterGroup, candidateCertificationsController *controllers.CandidateCertificationsController) {
	certificationsRoute := rg.Group("/certifications")
	certificationsRoute.POST("/", candidateCertificationsController.AddCertification)
	certificationsRoute.GET("/", candidateCertificationsController.GetCertifications)
	certificationsRoute.DELETE("/:certificationName", candidateCertificationsController.DeleteCertification)

}

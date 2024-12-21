package candidate

import (
	request "dz-jobs-api/internal/dto/request/candidate"
	"dz-jobs-api/internal/dto/response"
	responseCandidate "dz-jobs-api/internal/dto/response/candidate"
	serviceInterfaces "dz-jobs-api/internal/services/interfaces/candidate"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CandidateCertificationsController struct {
	service serviceInterfaces.CandidateCertificationsService
}

func NewCandidateCertificationsController(service serviceInterfaces.CandidateCertificationsService) *CandidateCertificationsController {
	return &CandidateCertificationsController{service: service}
}

func (c *CandidateCertificationsController) CreateCertification(ctx *gin.Context) {
	candidateID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	var req request.AddCertificationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	certification, err := c.service.AddCertification(candidateID, req)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, response.Response{
		Code:    http.StatusCreated,
		Status:  "Created",
		Message: "Certification created successfully",
		Data:    responseCandidate.ToCertificationResponse(certification),
	})
}

func (c *CandidateCertificationsController) GetCertfications(ctx *gin.Context) {
	candidateID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	experience, err := c.service.GetCertificationsByCandidateID(candidateID)
	if err != nil {
		ctx.Error(err)
		return
	}
	var certificationResponses []responseCandidate.CertificationResponse
	for _, cer := range experience {
		certificationResponses = append(certificationResponses, responseCandidate.ToCertificationResponse(&cer))
	}
	ctx.JSON(http.StatusCreated, response.Response{
		Code:    http.StatusCreated,
		Status:  "Created",
		Message: "Certifications retrieved successfully",
		Data:    certificationResponses,
	})
}

func (c *CandidateCertificationsController) DeleteCertification(ctx *gin.Context) {
	candidateID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	err = c.service.DeleteCertification(candidateID, ctx.Param("certification_id"))
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Certification deleted successfully",
	})
}

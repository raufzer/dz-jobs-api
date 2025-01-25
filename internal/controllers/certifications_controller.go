package controllers

import (
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/dto/response"
	serviceInterfaces "dz-jobs-api/internal/services/interfaces"
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

// AddCertification godoc
// @Summary Add a new certification
// @Description Add a new certification for a candidate by candidate ID
// @Tags Candidates - Certifications
// @Accept json
// @Produce json
// @Param certification body request.AddCertificationRequest true "Certification request"
// @Success 201 {object} response.Response{Data=response.CertificationResponse} "Certification created successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 404 {object} response.Response "Candidate not found"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /candidates/certifications [post]
func (c *CandidateCertificationsController) AddCertification(ctx *gin.Context) {
	userID := ctx.MustGet("candidate_id")
	candidateID, err := uuid.Parse(userID.(string))
	if err != nil {
		_  = ctx.Error(err)
		return
	}
	var req request.AddCertificationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_  = ctx.Error(err)
		ctx.Abort()
		return
	}

	certification, err := c.service.AddCertification(candidateID, req)
	if err != nil {
		_  = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, response.Response{
		Code:    http.StatusCreated,
		Status:  "Created",
		Message: "Certification created successfully",
		Data:    response.ToCertificationResponse(certification),
	})
}

// GetCertifications godoc
// @Summary Get certifications
// @Description Get all certifications for a candidate by candidate ID
// @Tags Candidates - Certifications
// @Produce json
// @Success 200 {object} response.Response{Data=response.CertificationsResponseData} "Certifications retrieved successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Candidate not found"
// @Failure 404 {object} response.Response "Certifications not found"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /candidates/certifications [get]
func (c *CandidateCertificationsController) GetCertifications(ctx *gin.Context) {
	userID := ctx.MustGet("candidate_id")
	candidateID, err := uuid.Parse(userID.(string))
	if err != nil {
		_  = ctx.Error(err)
		return
	}

	certifications, err := c.service.GetCertifications(candidateID)
	if err != nil {
		_  = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Certifications retrieved successfully",
		Data:    response.ToCertificationsResponse(certifications),
	})
}

// DeleteCertification godoc
// @Summary Delete certification
// @Description Delete a certification by candidate ID and certification ID
// @Tags Candidates - Certifications
// @Produce json
// @Param certificationName path string true "Certification Name"
// @Success 200 {object} response.Response "Certification deleted successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "You do not own this certification"
// @Failure 404 {object} response.Response "Candidate not found"
// @Failure 404 {object} response.Response "Certification not found"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /candidates/certifications/{certificationName} [delete]
func (c *CandidateCertificationsController) DeleteCertification(ctx *gin.Context) {
	userID := ctx.MustGet("candidate_id")
	candidateID, err := uuid.Parse(userID.(string))
	if err != nil {
		_  = ctx.Error(err)
		return
	}

	err = c.service.DeleteCertification(candidateID, ctx.Param("certificationName"))
	if err != nil {
		_  = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Certification deleted successfully",
	})
}

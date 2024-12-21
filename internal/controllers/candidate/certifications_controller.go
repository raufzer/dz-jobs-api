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

// CreateCertification godoc
// @Summary Create a new certification
// @Description Create a new certification for a candidate
// @Tags certifications_1create
// @Accept json
// @Produce json
// @Param id path string true "Candidate ID"
// @Param certification body request.AddCertificationRequest true "Certification request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /candidates/{id}/certifications [post]
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

// GetCertifications godoc
// @Summary Get certifications by candidate ID
// @Description Get all certifications for a candidate by candidate ID
// @Tags certifications_2get
// @Produce json
// @Param id path string true "Candidate ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /candidates/{id}/certifications [get]
func (c *CandidateCertificationsController) GetCertifications(ctx *gin.Context) {
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
    ctx.JSON(http.StatusOK, response.Response{
        Code:    http.StatusOK,
        Status:  "OK",
        Message: "Certifications retrieved successfully",
        Data:    certificationResponses,
    })
}

// DeleteCertification godoc
// @Summary Delete certification
// @Description Delete a certification by candidate ID and certification ID
// @Tags certifications_3delete
// @Produce json
// @Param id path string true "Candidate ID"
// @Param certification_id path string true "Certification ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /candidates/{id}/certifications/{certification_id} [delete]
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
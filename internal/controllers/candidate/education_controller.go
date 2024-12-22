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

type CandidateEducationController struct {
    service serviceInterfaces.CandidateEducationService
}

func NewCandidateEducationController(service serviceInterfaces.CandidateEducationService) *CandidateEducationController {
    return &CandidateEducationController{service: service}
}

// AddEducation godoc
// @Summary Add a new education record
// @Description Add a new education record for a candidate by candidate ID
// @Tags Candidates - Education
// @Accept json
// @Produce json
// @Param candidate_id path string true "Candidate ID"
// @Param education body request.AddEducationRequest true "Education request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /candidates/{candidate_id}/education [post]
func (c *CandidateEducationController) AddEducation(ctx *gin.Context) {
    candidateID, err := uuid.Parse(ctx.Param("id"))
    if err != nil {
        ctx.Error(err)
        ctx.Abort()
        return
    }

    var req request.AddEducationRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.Error(err)
        ctx.Abort()
        return
    }

    education, err := c.service.AddEducation(candidateID, req)
    if err != nil {
        ctx.Error(err)
        return
    }
    ctx.JSON(http.StatusCreated, response.Response{
        Code:    http.StatusCreated,
        Status:  "Created",
        Message: "Education created successfully",
        Data:    responseCandidate.ToEducationResponse(education),
    })
}

// GetEducation godoc
// @Summary Get education records
// @Description Get all education records for a candidate by candidate ID
// @Tags Candidates - Education
// @Produce json
// @Param candidate_id path string true "Candidate ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /candidates/{candidate_id}/education [get]
func (c *CandidateEducationController) GetEducation(ctx *gin.Context) {
    candidateID, err := uuid.Parse(ctx.Param("id"))
    if err != nil {
        ctx.Error(err)
        ctx.Abort()
        return
    }

    education, err := c.service.GetEducation(candidateID)
    if err != nil {
        ctx.Error(err)
        return
    }
    var educationResponses []responseCandidate.EducationResponse
    for _, edu := range education {
        educationResponses = append(educationResponses, responseCandidate.ToEducationResponse(&edu))
    }
    ctx.JSON(http.StatusOK, response.Response{
        Code:    http.StatusOK,
        Status:  "OK",
        Message: "Education information retrieved successfully",
        Data:    educationResponses,
    })
}

// DeleteEducation godoc
// @Summary Delete education record
// @Description Delete an education record by candidate ID
// @Tags Candidates - Education
// @Produce json
// @Param candidate_id path string true "Candidate ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /candidates/{candidate_id}/education [delete]
func (c *CandidateEducationController) DeleteEducation(ctx *gin.Context) {
    candidateID, err := uuid.Parse(ctx.Param("id"))
    if err != nil {
        ctx.Error(err)
        ctx.Abort()
        return
    }

    err = c.service.DeleteEducation(candidateID)
    if err != nil {
        ctx.Error(err)
        return
    }
    ctx.JSON(http.StatusOK, response.Response{
        Code:    http.StatusOK,
        Status:  "OK",
        Message: "Education deleted successfully",
    })
}
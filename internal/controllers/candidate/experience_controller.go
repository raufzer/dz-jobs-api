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

type CandidateExperienceController struct {
    service serviceInterfaces.CandidateExperienceService
}

func NewCandidateExperienceController(service serviceInterfaces.CandidateExperienceService) *CandidateExperienceController {
    return &CandidateExperienceController{service: service}
}

// AddExperience godoc
// @Summary Add a new experience record
// @Description Add a new experience record for a candidate by candidate ID
// @Tags Candidates - Experience
// @Accept json
// @Produce json
// @Param candidate_id path string true "Candidate ID"
// @Param experience body request.AddExperienceRequest true "Experience request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /candidates/{candidate_id}/experience [post]
func (c *CandidateExperienceController) AddExperience(ctx *gin.Context) {
    candidateID, err := uuid.Parse(ctx.Param("candidate_id"))
    if err != nil {
        ctx.Error(err)
        ctx.Abort()
        return
    }
    var req request.AddExperienceRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.Error(err)
        ctx.Abort()
        return
    }

    experience, err := c.service.AddExperience(candidateID, req)
    if err != nil {
        ctx.Error(err)
        return
    }
    ctx.JSON(http.StatusCreated, response.Response{
        Code:    http.StatusCreated,
        Status:  "Created",
        Message: "Experience created successfully",
        Data:    responseCandidate.ToExperienceResponse(experience),
    })
}

// GetExperience godoc
// @Summary Get experience records
// @Description Get all experience records for a candidate by candidate ID
// @Tags Candidates - Experience
// @Produce json
// @Param candidate_id path string true "Candidate ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /candidates/{candidate_id}/experience [get]
func (c *CandidateExperienceController) GetExperience(ctx *gin.Context) {
    candidateID, err := uuid.Parse(ctx.Param("candidate_id"))
    if err != nil {
        ctx.Error(err)
        ctx.Abort()
        return
    }

    experience, err := c.service.GetExperience(candidateID)
    if err != nil {
        ctx.Error(err)
        return
    }
    var experienceResponses []responseCandidate.ExperienceResponse
    for _, exp := range experience {
        experienceResponses = append(experienceResponses, responseCandidate.ToExperienceResponse(&exp))
    }
    ctx.JSON(http.StatusOK, response.Response{
        Code:    http.StatusOK,
        Status:  "OK",
        Message: "Experience retrieved successfully",
        Data:    experienceResponses,
    })
}

// DeleteExperience godoc
// @Summary Delete experience record
// @Description Delete an experience record by candidate ID
// @Tags Candidates - Experience
// @Produce json
// @Param candidate_id path string true "Candidate ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /candidates/{candidate_id}/experience [delete]
func (c *CandidateExperienceController) DeleteExperience(ctx *gin.Context) {
    candidateID, err := uuid.Parse(ctx.Param("candidate_id"))
    if err != nil {
        ctx.Error(err)
        ctx.Abort()
        return
    }

    err = c.service.DeleteExperience(candidateID)
    if err != nil {
        ctx.Error(err)
        return
    }
    ctx.JSON(http.StatusOK, response.Response{
        Code:    http.StatusOK,
        Status:  "OK",
        Message: "Experience deleted successfully",
    })
}
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

type CandidatePersonalInfoController struct {
	service serviceInterfaces.CandidatePersonalInfoService
}

func NewCandidatePersonalInfoController(service serviceInterfaces.CandidatePersonalInfoService) *CandidatePersonalInfoController {
	return &CandidatePersonalInfoController{service: service}
}

func (c *CandidatePersonalInfoController) GetPersonalInfoByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	candidateID, err := uuid.Parse(idParam)
	if err != nil {
		ctx.Error(err)
		return
	}

	info, err := c.service.GetPersonalInfo(candidateID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Personal information retrieved successfully",
		Data:    responseCandidate.ToPersonalInfoResponse(info),
	})
	ctx.JSON(http.StatusOK, responseCandidate.ToPersonalInfoResponse(info))
}

func (c *CandidatePersonalInfoController) UpdatePersonalInfo(ctx *gin.Context) {
	var req request.UpdateCandidatePersonalInfoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	idParam := ctx.Param("id")
	candidateID, err := uuid.Parse(idParam)
	if err != nil {
		ctx.Error(err)
		return
	}

	updatedInfo, err := c.service.UpdatePersonalInfo(candidateID, req)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Personal information updated successfully",
		Data:    responseCandidate.ToPersonalInfoResponse(updatedInfo),
	})
}


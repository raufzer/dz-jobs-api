package controllers

import (
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/dto/response"
	serviceInterfaces "dz-jobs-api/internal/services/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RecruiterController struct {
	recruiterService serviceInterfaces.RecruiterService
}

func NewRecruiterController(service serviceInterfaces.RecruiterService) *RecruiterController {
	return &RecruiterController{
		recruiterService: service,
	}
}

// CreateRecruiter godoc
// @Summary Create a new recruiter
// @Description Create a new recruiter with company logo
// @Tags Recruiters
// @Accept multipart/form-data
// @Produce json
// @Param company_logo formData file true "Company Logo"
// @Param recruiter body request.CreateRecruiterRequest true "Recruiter request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /recruiters [post]
func (rc *RecruiterController) CreateRecruiter(ctx *gin.Context) {
	accessToken, err := ctx.Cookie("access_token")
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	userID, err := rc.recruiterService.ExtractTokenDetails(accessToken)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	companyLogoFile, err := ctx.FormFile("company_logo")
	if err != nil {
		ctx.Error(err)
		return
	}
	var req request.CreateRecruiterRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	recruiter, err := rc.recruiterService.CreateRecruiter(userID, req, companyLogoFile)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, response.Response{
		Code:    http.StatusCreated,
		Status:  "Created",
		Message: "Recruiter created successfully",
		Data:    response.ToRecruiterResponse(recruiter),
	})
}

// GetRecruiterByID godoc
// @Summary Get recruiter by ID
// @Description Get recruiter details by ID
// @Tags Recruiters
// @Produce json
// @Param id path string true "Recruiter ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /recruiters/{id} [get]
func (rc *RecruiterController) GetRecruiterByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	recruiter, err := rc.recruiterService.GetRecruiterByID(id)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Recruiter found",
		Data:    response.ToRecruiterResponse(recruiter),
	})
}

// UpdateRecruiter godoc
// @Summary Update recruiter
// @Description Update recruiter details with company logo
// @Tags Recruiters
// @Accept multipart/form-data
// @Produce json
// @Param id path string true "Recruiter ID"
// @Param company_logo formData file true "Company Logo"
// @Param recruiter body request.UpdateRecruiterRequest true "Recruiter request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /recruiters/{id} [put]
func (rc *RecruiterController) UpdateRecruiter(ctx *gin.Context) {
	var req request.UpdateRecruiterRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		return
	}
	companyLogoFile, err := ctx.FormFile("company_logo")
	if err != nil {
		ctx.Error(err)
		return
	}
	updatedRecruiter, err := rc.recruiterService.UpdateRecruiter(id, req, companyLogoFile)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Recruiter updated successfully",
		Data:    response.ToRecruiterResponse(updatedRecruiter),
	})
}

// DeleteRecruiter godoc
// @Summary Delete recruiter
// @Description Delete recruiter by ID
// @Tags Recruiters
// @Produce json
// @Param id path string true "Recruiter ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /recruiters/{id} [delete]
func (rc *RecruiterController) DeleteRecruiter(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	err = rc.recruiterService.DeleteRecruiter(id)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Recruiter deleted successfully",
	})
}

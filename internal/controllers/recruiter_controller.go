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
// @Tags Recruiters - Recruiter
// @Accept multipart/form-data
// @Produce json
// @Param company_logo formData file true "Company Logo"
// @Param recruiter body request.CreateRecruiterRequest true "Recruiter request"
// @Success 201 {object} response.Response{Data=response.RecruiterResponse} "Recruiter created successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /recruiters [post]
func (c *RecruiterController) CreateRecruiter(ctx *gin.Context) {

	recuiterID := ctx.MustGet("recruiter_id")
	userID := recuiterID.(string)

	companyLogoFile, err := ctx.FormFile("company_logo")
	if err != nil {
		_  = ctx.Error(err)
		return
	}
	var req request.CreateRecruiterRequest
	if err := ctx.ShouldBind(&req); err != nil {
		_  = ctx.Error(err)
		ctx.Abort()
		return
	}
	recruiter, err := c.recruiterService.CreateRecruiter(userID, req, companyLogoFile)
	if err != nil {
		_  = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, response.Response{
		Code:    http.StatusCreated,
		Status:  "Created",
		Message: "Recruiter created successfully",
		Data:    response.ToRecruiterResponse(recruiter),
	})
}

// GetRecruiter godoc
// @Summary Get recruiter
// @Description Get recruiter details by recruiter_id
// @Tags Recruiters - Recruiter
// @Produce json
// @Success 200 {object} response.Response{Data=response.RecruiterResponse} "Recruiter found"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Recruiter not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /recruiters [get]
func (c *RecruiterController) GetRecruiter(ctx *gin.Context) {
	userID := ctx.MustGet("recruiter_id")
	recruiterID, err := uuid.Parse(userID.(string))
	if err != nil {
		_  = ctx.Error(err)
		return
	}
	recruiter, err := c.recruiterService.GetRecruiter(recruiterID)
	if err != nil {
		_  = ctx.Error(err)
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
// @Description Update recruiter details with company logo by recruiter_id
// @Tags Recruiters - Recruiter
// @Accept multipart/form-data
// @Produce json
// @Param company_logo formData file true "Company Logo"
// @Param recruiter body request.UpdateRecruiterRequest true "Recruiter request"
// @Success 200 {object} response.Response{Data=response.RecruiterResponse} "Recruiter updated successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Recruiter not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /recruiters [put]
func (c *RecruiterController) UpdateRecruiter(ctx *gin.Context) {
	var req request.UpdateRecruiterRequest
	if err := ctx.ShouldBind(&req); err != nil {
		_  = ctx.Error(err)
		ctx.Abort()
		return
	}
	userID := ctx.MustGet("recruiter_id")
	recruiterID, err := uuid.Parse(userID.(string))
	if err != nil {
		_  = ctx.Error(err)
		return
	}
	companyLogoFile, err := ctx.FormFile("company_logo")
	if err != nil {
		_  = ctx.Error(err)
		return
	}
	updatedRecruiter, err := c.recruiterService.UpdateRecruiter(recruiterID, req, companyLogoFile)
	if err != nil {
		_  = ctx.Error(err)
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
// @Description Delete recruiter by recruiter_id
// @Tags Recruiters - Recruiter
// @Produce json
// @Success 200 {object} response.Response "Recruiter deleted successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Recruiter not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /recruiters [delete]
func (c *RecruiterController) DeleteRecruiter(ctx *gin.Context) {
	userID := ctx.MustGet("recruiter_id")
	recruiterID, err := uuid.Parse(userID.(string))
	if err != nil {
		_  = ctx.Error(err)
		return
	}
	err = c.recruiterService.DeleteRecruiter(recruiterID)
	if err != nil {
		_  = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Recruiter deleted successfully",
	})
}

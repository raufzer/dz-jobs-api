package controllers

import (
	"dz-jobs-api/internal/dto/response"
	serviceInterfaces "dz-jobs-api/internal/services/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookmarksController struct {
	service serviceInterfaces.BookmarksService
}

func NewBookmarksController(service serviceInterfaces.BookmarksService) *BookmarksController {
	return &BookmarksController{service: service}
}

// AddBookmark godoc
// @Summary Add a job to bookmarks
// @Description Add a job to bookmarks for a candidate by candidate ID
// @Tags Candidates - Bookmarks
// @Accept json
// @Produce json
// @Param JobId path int true "Job ID"
// @Success 201 {object} response.Response "Job added successfully to bookmarks"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Candidate not found"
// @Failure 404 {object} response.Response "Job not found"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /candidates/bookmarks/{JobId} [post]
func (c *BookmarksController) AddBookmark(ctx *gin.Context) {
	userID := ctx.MustGet("candidate_id")
	candidateID, err := uuid.Parse(userID.(string))
	if err != nil {
		_  = ctx.Error(err)
		return
	}
	jobIDStr := ctx.Param("JobId")
	jobID, err := strconv.ParseInt(jobIDStr, 10, 64)
	if err != nil {
		_  = ctx.Error(err)
		ctx.Abort()
		return
	}

	if err := c.service.AddBookmark(candidateID, jobID); err != nil {
		_  = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, response.Response{
		Code:    http.StatusCreated,
		Status:  "Created",
		Message: "Job added successfully to bookmarks",
	})
}

// GetBookmarks godoc
// @Summary Get jobs bookmarked
// @Description Retrieve a list of jobs bookmarked by a candidate by candidate ID
// @Tags Candidates - Bookmarks
// @Produce json
// @Success 200 {object} response.Response{Data=response.JobsResponseData} "Jobs bookmarked retrieved successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Candidate not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /candidates/bookmarks/ [get]
func (c *BookmarksController) GetBookmarks(ctx *gin.Context) {
	userID := ctx.MustGet("candidate_id")
	candidateID, _ := uuid.Parse(userID.(string))
	jobs, err := c.service.GetBookmarks(candidateID)
	if err != nil {
		_  = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Jobs bookmarked retrieved successfully",
		Data:    response.ToJobsResponse(jobs),
	})
}

// RemoveBookmark godoc
// @Summary Remove a job from bookmarks
// @Description Remove a job from bookmarks for a candidate by candidate ID
// @Tags Candidates - Bookmarks
// @Accept json
// @Produce json
// @Param JobId path int true "Job ID"
// @Success 201 {object} response.Response "Job removed successfully from bookmarks"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Candidate not found"
// @Failure 404 {object} response.Response "Job not found"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /candidates/bookmarks/{JobId} [delete]
func (c *BookmarksController) RemoveBookmark(ctx *gin.Context) {
	userID := ctx.MustGet("candidate_id")
	candidateID, _ := uuid.Parse(userID.(string))
	jobIDStr := ctx.Param("JobId")
	jobID, err := strconv.ParseInt(jobIDStr, 10, 64)
	if err != nil {
		_  = ctx.Error(err)
		ctx.Abort()
		return
	}
	if err := c.service.RemoveBookmark(candidateID, jobID); err != nil {
		_  = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Job removed successfully from bookmarks",
	})
}

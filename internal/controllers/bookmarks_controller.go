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
// @Param candidate_id path string true "Candidate ID"
// @Param job_id path int true "Job ID"
// @Success 201 {object} response.Response "Job added successfully to bookmarks"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Candidate not found"
// @Failure 404 {object} response.Response "Job not found"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /candidates/{candidate_id}/bookmarks/job_id [post]
func (c *BookmarksController) AddBookmark(ctx *gin.Context) {
	candidateID, err := uuid.Parse(ctx.Param("candidate_id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	jobIDStr := ctx.Param("job_id")
	jobID, err := strconv.ParseInt(jobIDStr, 10, 64)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	if err := c.service.AddBookmark(candidateID, jobID); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, response.Response{
		Code:    http.StatusCreated,
		Status:  "Created",
		Message: "Job added successfully to bookmarks",
	})
}

// GetJobListingsByStatus godoc
// @Summary Get job listings by status
// @Description Retrieve a list of jobs bookmarked by a candidate "{total: int, jobs: []response.JobResponse}"}
// @Tags Candidates - Bookmarks
// @Produce json
// @Param status query string true "Job status (e.g., open, closed)"
// @Success 200 {object} response.Response{Data=response.JobsResponseData} "Jobs bookmarked retrieved successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Candidate not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /candidates/{candidate_id}/bookmarks/ [get]
func (c *BookmarksController) GetBookmarks(ctx *gin.Context) {
	candidateID, err := uuid.Parse(ctx.Param("candidate_id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	jobs, err := c.service.GetBookmarks(candidateID)
	if err != nil {
		ctx.Error(err)
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
// @Param candidate_id path string true "Candidate ID"
// @Param job_id path int true "Job ID"
// @Success 201 {object} response.Response "Job removed successfully from bookmarks"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 404 {object} response.Response "Candidate not found"
// @Failure 404 {object} response.Response "Job not found"
// @Failure 500 {object} response.Response "An unexpected error occurred"
// @Router /candidates/{candidate_id}/bookmarks/job_id [delete]
func (c *BookmarksController) RemoveBookmark(ctx *gin.Context) {
	candidateID, err := uuid.Parse(ctx.Param("candidate_id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	jobIDStr := ctx.Param("job_id")
	jobID, err := strconv.ParseInt(jobIDStr, 10, 64)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	if err := c.service.RemoveBookmark(candidateID, jobID); err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Job removed successfully from bookmarks",
	})
}

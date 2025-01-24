package controllers

import (
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/dto/response"
	serviceInterfaces "dz-jobs-api/internal/services/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// JobController handles job-related API requests
type JobController struct {
	jobService serviceInterfaces.JobService
}

// NewJobController creates a new instance of JobController
func NewJobController(service serviceInterfaces.JobService) *JobController {
	return &JobController{
		jobService: service,
	}
}

// PostNewJob godoc
// @Summary Post a new job
// @Description Allows recruiters to post a new job (NOTE: salary range must be in this format: 'min - max DZD')
// @Tags Recruiters - Jobs
// @Accept json
// @Produce json
// @Param job body request.PostNewJobRequest true "Job details"
// @Success 201 {object} response.Response{Data=response.JobResponse} "Job posted successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /recruiters/jobs [post]
func (c *JobController) PostNewJob(ctx *gin.Context) {
	userID := ctx.MustGet("recruiter_id")
	recruiterID, _ := uuid.Parse(userID.(string))
	var req request.PostNewJobRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	job, err := c.jobService.PostNewJob(recruiterID, req)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, response.Response{
		Code:    http.StatusCreated,
		Status:  "Created",
		Message: "Job posted successfully",
		Data:    response.ToJobResponse(job),
	})
}

// GetJobDetails godoc
// @Summary Get job details
// @Description Retrieve the details of a specific job by jobId
// @Tags Recruiters - Jobs
// @Produce json
// @Param jobId path int true "Job ID"
// @Success 200 {object} response.Response{Data=response.JobResponse} "Job details found"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "You do not own this Job"
// @Failure 404 {object} response.Response "Jobs not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /recruiters/jobs/{jobId} [get]
func (c *JobController) GetJobDetails(ctx *gin.Context) {
	jobIDStr := ctx.Param("jobId")
	jobID, err := strconv.ParseInt(jobIDStr, 10, 64)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	userID := ctx.MustGet("recruiter_id")
	recruiterID, _ := uuid.Parse(userID.(string))
	job, err := c.jobService.GetJobDetails(jobID, recruiterID)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Job details found",
		Data:    response.ToJobResponse(job),
	})
}

// GetJobListingsByStatus godoc
// @Summary Get job listings by status
// @Description Retrieve a list of jobs filtered by their status (e.g., open, closed)
// @Tags Recruiters - Jobs
// @Produce json
// @Param status query string true "Job status (e.g., open, closed)"
// @Success 200 {object} response.Response{Data=response.JobsResponseData} "Jobs retrieved successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "You do not own these Jobs"
// @Failure 404 {object} response.Response "Jobs not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /recruiters/jobs [get]
func (c *JobController) GetJobListingsByStatus(ctx *gin.Context) {
	userID := ctx.MustGet("recruiter_id")
	recruiterID, _ := uuid.Parse(userID.(string))
	jobs, err := c.jobService.GetJobListingsByStatus(ctx.Query("status"), recruiterID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Jobs retrieved successfully",
		Data:    response.ToJobsResponse(jobs),
	})
}

// Edit Jobgodoc
// @Summary Edit an existing job
// @Description Update the details of a specific job by jobId
// @Tags Recruiters - Jobs
// @Accept json
// @Produce json
// @Param jobId path int true "Job ID"
// @Param job body request.EditJobRequest true "Updated job details"
// @Success 200 {object} response.Response{Data=response.JobResponse} "Job updated successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "You do not own this Job"
// @Failure 404 {object} response.Response "Job notfound"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /recruiters/jobs/{jobId} [put]
func (c *JobController) EditJob(ctx *gin.Context) {
	userID := ctx.MustGet("recruiter_id")
	recruiterID, _ := uuid.Parse(userID.(string))
	var req request.EditJobRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	jobIDStr := ctx.Param("jobId")
	jobID, err := strconv.ParseInt(jobIDStr, 10, 64)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	updatedJob, err := c.jobService.EditJob(jobID, req, recruiterID)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Job updated successfully",
		Data:    response.ToJobResponse(updatedJob),
	})
}

// DeactivateJob deactivates a job
// @Summary Deactivate a job
// @Description Disable a specific job by jobId
// @Tags Recruiters - Jobs
// @Produce json
// @Param jobId path int true "Job ID"
// @Success 200 {object} response.Response{Data=response.JobResponse} "Job deactivated successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "You do not own this Job"
// @Failure 404 {object} response.Response "Job not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /recruiters/jobs/{jobId}/deactivate [put]
func (c *JobController) DeactivateJob(ctx *gin.Context) {
	jobIDStr := ctx.Param("jobId")
	jobID, err := strconv.ParseInt(jobIDStr, 10, 64)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	userID := ctx.MustGet("recruiter_id")
	recruiterID, _ := uuid.Parse(userID.(string))
	updatedJob, err := c.jobService.DeactivateJob(jobID, recruiterID)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Job deactivated successfully",
		Data:    response.ToJobResponse(updatedJob),
	})
}

// RepostJob godoc
// @Summary Repost a job
// @Description Repost a deactivated job by jobId
// @Tags Recruiters - Jobs
// @Produce json
// @Param jobId path int true "Job ID"
// @Success 200 {object} response.Response{Data=response.JobResponse} "Job reposted successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "You do not own this Job"
// @Failure 404 {object} response.Response "Job not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /recruiters/jobs/{jobId}/repost [put]
func (c *JobController) RepostJob(ctx *gin.Context) {
	jobIDStr := ctx.Param("jobId")
	jobID, err := strconv.ParseInt(jobIDStr, 10, 64)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	userID := ctx.MustGet("recruiter_id")
	recruiterID, _ := uuid.Parse(userID.(string))
	updatedJob, err := c.jobService.RepostJob(jobID, recruiterID)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Job reposted successfully",
		Data:    response.ToJobResponse(updatedJob),
	})
}

// DeleteJob godoc
// @Summary Delete a job
// @Description Remove a specific job from the system by its ID
// @Tags Recruiters - Jobs
// @Produce json
// @Param jobId path int true "Job ID"
// @Success 200 {object} response.Response "Job deleted successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "You do not own this Job"
// @Failure 404 {object} response.Response "Job not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /recruiters/jobs/{jobId} [delete]
func (c *JobController) DeleteJob(ctx *gin.Context) {
	jobIDStr := ctx.Param("jobId")
	jobID, err := strconv.ParseInt(jobIDStr, 10, 64)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	userID := ctx.MustGet("recruiter_id")
	recruiterID, _ := uuid.Parse(userID.(string))
	err = c.jobService.DeleteJob(jobID, recruiterID)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Job deleted successfully",
	})
}

// GetAllJobs godoc
// @Summary Get all jobs
// @Description Retrieve all jobs in the system
// @Tags Jobs
// @Produce json
// @Success 200 {object} response.Response{Data=response.JobsResponseData} "Jobs retrieved successfully"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /jobs [get]
func (c *JobController) GetAllJobs(ctx *gin.Context) {
	jobs, err := c.jobService.GetAllJobs()
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Jobs retrieved successfully",
		Data:    response.ToJobsResponse(jobs),
	})
}

// SearchJobs godoc
// @Summary Search for jobs
// @Description Search for jobs using various filters
// @Tags Jobs
// @Accept json
// @Produce json
// @Param filters query request.JobFilters true "Job search filters"
// @Success 200 {object} response.Response{Data=response.JobsResponseData} "Jobs found successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 404 {object} response.Response "Jobs not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /jobs/search [get]
func (c *JobController) SearchJobs(ctx *gin.Context) {
	var filters request.JobFilters
	if err := ctx.ShouldBindQuery(&filters); err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	jobs, err := c.jobService.SearchJobs(filters)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Jobs found successfully",
		Data:    response.ToJobsResponse(jobs),
	})
}

// GetJobDetailsPublic godoc
// @Summary Get job details
// @Description Retrieve the details of a specific job by jobId
// @Tags Jobs
// @Produce json
// @Param jobId path int true "Job ID"
// @Success 200 {object} response.Response{Data=response.JobResponse} "Job details found"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 404 {object} response.Response "Job not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /jobs/{jobId} [get]
func (c *JobController) GetJobDetailsPublic(ctx *gin.Context) {
	jobIDStr := ctx.Param("jobId")
	jobID, err := strconv.ParseInt(jobIDStr, 10, 64)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	job, err := c.jobService.GetJobDetailsPublic(jobID)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Job details found",
		Data:    response.ToJobResponse(job),
	})
}

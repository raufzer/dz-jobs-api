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
// @Description Allows recruiters to post a new job
// @Tags Recruiters - Jobs 
// @Accept json
// @Produce json
// @Param recruiter_id path string true "Recruiter ID"
// @Param job body request.PostNewJobRequest true "Job details"
// @Success 201 {object} response.Response{Data=response.JobResponse} "Job posted successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /recruiters/{recruiter_id}/jobs [post]
func (c *JobController) PostNewJob(ctx *gin.Context) {
	recruiterID, err := uuid.Parse(ctx.Param("recruiter_id"))
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
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
// @Description Retrieve the details of a specific job by job_id
// @Tags Recruiters - Jobs 
// @Produce json
// @Param job_id path int true "Job ID"
// @Success 200 {object} response.Response{Data=response.JobResponse} "Job details found"
// @Failure 400 {object} response.Response "Invalid job ID"
// @Failure 404 {object} response.Response "Job not found"
// @Router /recruiters/{recruiter_id}/jobs/{job_id} [get]
func (c *JobController) GetJobDetails(ctx *gin.Context) {
	jobIDStr := ctx.Param("job_id")
	jobID, err := strconv.ParseInt(jobIDStr, 10, 64)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	job, err := c.jobService.GetJobDetails(jobID)
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
// @Success 200 {object} response.Response{Data=object} "Job listings retrieved successfully"
// @Failure 400 {object} response.Response "Invalid status query"
// @Router /recruiters/{recruiter_id}/jobs [get]
func (c *JobController) GetJobListingsByStatus(ctx *gin.Context) {
	jobs, err := c.jobService.GetJobListingsByStatus(ctx.Query("status"))
	if err != nil {
		ctx.Error(err)
		return
	}
	jobResponses := make([]response.JobResponse, len(jobs))
	for i, job := range jobs {
		jobResponses[i] = response.ToJobResponse(job)
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Jobs retrieved successfully",
		Data: gin.H{
			"total": len(jobs),
			"jobs":  jobResponses,
		},
	})
}

// Edit Jobgodoc
// @Summary Edit an existing job
// @Description Update the details of a specific job by job_id
// @Tags Recruiters - Jobs 
// @Accept json
// @Produce json
// @Param job_id path int true "Job ID"
// @Param job body request.EditJobRequest true "Updated job details"
// @Success 200 {object} response.Response{Data=response.JobResponse} "Job updated successfully"
// @Failure 400 {object} response.Response "Invalid input"
// @Failure 404 {object} response.Response "Job not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /recruiters/{recruiter_id}/jobs{job_id} [put]
func (c *JobController) EditJob(ctx *gin.Context) {
	var req request.EditJobRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
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
	updatedJob, err := c.jobService.EditJob(jobID, req)
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
// @Description Disable a specific job by job_id
// @Tags Recruiters - Jobs 
// @Produce json
// @Param job_id path int true "Job ID"
// @Success 200 {object} response.Response{Data=response.JobResponse} "Job deactivated successfully"
// @Failure 404 {object} response.Response "Job not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /recruiters/{recruiter_id}/jobs/{job_id}/deactivate [put]
func (c *JobController) DeactivateJob(ctx *gin.Context) {
	jobIDStr := ctx.Param("job_id")
	jobID, err := strconv.ParseInt(jobIDStr, 10, 64)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	updatedJob, err := c.jobService.DeactivateJob(jobID)
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
// @Description Repost a deactivated job by job_id
// @Tags Recruiters - Jobs 
// @Produce json
// @Param job_id path int true "Job ID"
// @Success 200 {object} response.Response{Data=response.JobResponse} "Job reposted successfully"
// @Failure 404 {object} response.Response "Job not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /recruiters/{recruiter_id}/jobs/{job_id}/repost [put]
func (c *JobController) RepostJob(ctx *gin.Context) {
	jobIDStr := ctx.Param("job_id")
	jobID, err := strconv.ParseInt(jobIDStr, 10, 64)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	updatedJob, err := c.jobService.RepostJob(jobID)
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
// @Param job_id path int true "Job ID"
// @Success 200 {object} response.Response "Job deleted successfully"
// @Failure 404 {object} response.Response "Job not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /recruiters/{recruiter_id}/jobs/{job_id} [delete]
func (c *JobController) DeleteJob(ctx *gin.Context) {
	jobIDStr := ctx.Param("job_id")
	jobID, err := strconv.ParseInt(jobIDStr, 10, 64)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	err = c.jobService.DeleteJob(jobID)
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

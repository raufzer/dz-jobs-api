package response

import (
    "dz-jobs-api/internal/models"
    "time"

    "github.com/google/uuid"
)

type JobResponse struct {
    JobID          int64     `json:"job_id"`
    Title          string    `json:"title"`
    Description    string    `json:"description"`
    Location       string    `json:"location,omitempty"`
    SalaryRange    string    `json:"salary_range,omitempty"`
    RequiredSkills string    `json:"required_skills,omitempty"`
    RecruiterID    uuid.UUID `json:"recruiter_id"`
    CreatedAt      time.Time `json:"created_at"`
    UpdatedAt      time.Time `json:"updated_at"`
    Status         string    `json:"status"`
}

func ToJobResponse(job *models.Job) JobResponse {
    return JobResponse{
        JobID:          job.JobID,
        Title:          job.Title,
        Description:    job.Description,
        Location:       job.Location,
        SalaryRange:    job.SalaryRange,
        RequiredSkills: job.RequiredSkills,
        RecruiterID:    job.RecruiterID,
        CreatedAt:      job.CreatedAt,
        UpdatedAt:      job.UpdatedAt,
        Status:         job.Status,
    }
}

type JobsResponseData struct {
    Total int           `json:"total"`
    Jobs  []JobResponse `json:"jobs"`
}

func ToJobsResponse(jobs []*models.Job) JobsResponseData {
    var jobResponses []JobResponse
    for _, job := range jobs {
        jobResponses = append(jobResponses, ToJobResponse(job))
    }
    return JobsResponseData{
        Total: len(jobs),
        Jobs:  jobResponses,
    }
}

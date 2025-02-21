package services

import (
    "context" // Add this import
    "database/sql"
    "dz-jobs-api/internal/dto/request"
    "dz-jobs-api/internal/models"
    "dz-jobs-api/internal/repositories/interfaces"
    "dz-jobs-api/pkg/utils"
    "net/http"
    "time"

    "github.com/google/uuid"
)

type JobService struct {
    jobRepository interfaces.JobRepository
}

func NewJobService(jobRepo interfaces.JobRepository) *JobService {
    return &JobService{jobRepository: jobRepo}
}

func (s *JobService) PostNewJob(ctx context.Context, recruiterID uuid.UUID, req request.PostNewJobRequest) (*models.Job, error) {
    job := &models.Job{
        Title:          req.Title,
        Description:    req.Description,
        Location:       req.Location,
        SalaryRange:    req.SalaryRange,
        RequiredSkills: req.RequiredSkills,
        RecruiterID:    recruiterID,
        CreatedAt:      time.Now(),
        UpdatedAt:      time.Now(),
        Status:         req.Status,
        JobType:        req.JobType,
    }

    err := s.jobRepository.CreateJob(ctx, job) // Pass context
    if err != nil {
        return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to create job")
    }

    return job, nil
}

func (s *JobService) GetJobDetails(ctx context.Context, jobID int64, recruiterID uuid.UUID) (*models.Job, error) {
    err := s.jobRepository.ValidateJobOwnership(ctx, jobID, recruiterID) // Pass context
    if err != nil {
        return nil, utils.NewCustomError(http.StatusForbidden, "You do not own this job")
    }
    job, err := s.jobRepository.GetJobDetails(ctx, jobID, recruiterID) // Pass context
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, utils.NewCustomError(http.StatusNotFound, "Job not found")
        }

        return nil, utils.NewCustomError(http.StatusInternalServerError, "Error fetching job details")
    }

    if job.RecruiterID != recruiterID {
        return nil, utils.NewCustomError(http.StatusForbidden, "You do not own this job")
    }

    return job, nil
}

func (s *JobService) GetJobListingsByStatus(ctx context.Context, status string, recruiterID uuid.UUID) ([]*models.Job, error) {
    jobs, err := s.jobRepository.GetJobListingsByStatus(ctx, status, recruiterID) // Pass context
    if err != nil {
        return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to fetch jobs by status: "+status)
    }
    return jobs, nil
}

func (s *JobService) EditJob(ctx context.Context, jobID int64, req request.EditJobRequest, recruiterID uuid.UUID) (*models.Job, error) {
    err := s.jobRepository.ValidateJobOwnership(ctx, jobID, recruiterID) // Pass context
    if err != nil {
        return nil, utils.NewCustomError(http.StatusForbidden, "You do not own this job")
    }
    job, err := s.jobRepository.GetJobDetails(ctx, jobID, recruiterID) // Pass context
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, utils.NewCustomError(http.StatusNotFound, "Job not found")
        }
        return nil, utils.NewCustomError(http.StatusInternalServerError, "Error fetching job details")
    }

    if job.RecruiterID != recruiterID {
        return nil, utils.NewCustomError(http.StatusForbidden, "You do not own this job")
    }

    updatedJob := &models.Job{
        Title:          req.Title,
        Description:    req.Description,
        Location:       req.Location,
        SalaryRange:    req.SalaryRange,
        RequiredSkills: req.RequiredSkills,
        UpdatedAt:      time.Now(),
        JobType:        req.JobType,
    }

    if err := s.jobRepository.UpdateJob(ctx, jobID, recruiterID, updatedJob); err != nil { // Pass context
        if err == sql.ErrNoRows {
            return nil, utils.NewCustomError(http.StatusNotFound, "Job not found")
        }
        return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to update job")
    }

    return s.jobRepository.GetJobDetails(ctx, jobID, recruiterID) // Pass context
}

func (s *JobService) DeactivateJob(ctx context.Context, jobID int64, recruiterID uuid.UUID) (*models.Job, error) {
    err := s.jobRepository.ValidateJobOwnership(ctx, jobID, recruiterID) // Pass context
    if err != nil {
        return nil, utils.NewCustomError(http.StatusForbidden, "You do not own this job")
    }
    job, err := s.jobRepository.GetJobDetails(ctx, jobID, recruiterID) // Pass context
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, utils.NewCustomError(http.StatusNotFound, "Job not found")
        }
        return nil, utils.NewCustomError(http.StatusInternalServerError, "Error fetching job details")
    }

    if job.RecruiterID != recruiterID {
        return nil, utils.NewCustomError(http.StatusForbidden, "You do not own this job")
    }

    if err := s.jobRepository.DeactivateJob(ctx, jobID, recruiterID); err != nil { // Pass context
        if err == sql.ErrNoRows {
            return nil, utils.NewCustomError(http.StatusNotFound, "Job not found")
        }
        return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to deactivate job")
    }

    return s.jobRepository.GetJobDetails(ctx, jobID, recruiterID) // Pass context
}

func (s *JobService) RepostJob(ctx context.Context, jobID int64, recruiterID uuid.UUID) (*models.Job, error) {
    err := s.jobRepository.ValidateJobOwnership(ctx, jobID, recruiterID) // Pass context
    if err != nil {
        return nil, utils.NewCustomError(http.StatusForbidden, "You do not own this job")
    }
    job, err := s.jobRepository.GetJobDetails(ctx, jobID, recruiterID) // Pass context
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, utils.NewCustomError(http.StatusNotFound, "Job not found")
        }
        return nil, utils.NewCustomError(http.StatusInternalServerError, "Error fetching job details")
    }

    if job.RecruiterID != recruiterID {
        return nil, utils.NewCustomError(http.StatusForbidden, "You do not own this job")
    }

    if err := s.jobRepository.RepostJob(ctx, jobID, recruiterID); err != nil { // Pass context
        if err == sql.ErrNoRows {
            return nil, utils.NewCustomError(http.StatusNotFound, "Job not found")
        }
        return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to repost job")
    }

    return s.jobRepository.GetJobDetails(ctx, jobID, recruiterID) // Pass context
}

func (s *JobService) DeleteJob(ctx context.Context, jobID int64, recruiterID uuid.UUID) error {
    err := s.jobRepository.ValidateJobOwnership(ctx, jobID, recruiterID) // Pass context
    if err != nil {
        return utils.NewCustomError(http.StatusForbidden, "You do not own this job")
    }
    job, err := s.jobRepository.GetJobDetails(ctx, jobID, recruiterID) // Pass context
    if err != nil {
        if err == sql.ErrNoRows {
            return utils.NewCustomError(http.StatusNotFound, "Job not found")
        }
        return utils.NewCustomError(http.StatusInternalServerError, "Error fetching job details")
    }

    if job.RecruiterID != recruiterID {
        return utils.NewCustomError(http.StatusForbidden, "You do not own this job")
    }

    err = s.jobRepository.DeleteJob(ctx, jobID, recruiterID) // Pass context
    if err != nil {
        if err == sql.ErrNoRows {
            return utils.NewCustomError(http.StatusNotFound, "Job not found")
        }
        return utils.NewCustomError(http.StatusInternalServerError, "Failed to delete job")
    }
    return nil
}

func (s *JobService) GetAllJobs(ctx context.Context) ([]*models.Job, error) {
    jobs, err := s.jobRepository.GetAllJobs(ctx) // Pass context
    if err != nil {
        return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to fetch all jobs")
    }
    return jobs, nil
}

func (s *JobService) SearchJobs(ctx context.Context, filters request.JobFilters) ([]*models.Job, error) {
    jobs, err := s.jobRepository.GetJobListings(ctx, filters) // Pass context
    if err != nil {
        return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to search jobs")
    }
    return jobs, nil
}

func (s *JobService) GetJobDetailsPublic(ctx context.Context, jobID int64) (*models.Job, error) {
    job, err := s.jobRepository.GetJobDetailsPublic(ctx, jobID) // Pass context
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, utils.NewCustomError(http.StatusNotFound, "Job not found")
        }
        return nil, utils.NewCustomError(http.StatusInternalServerError, "Error fetching job details")
    }
    return job, nil
}
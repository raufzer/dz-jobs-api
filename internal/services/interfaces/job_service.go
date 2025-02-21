package interfaces

import (
    "context"
    "dz-jobs-api/internal/dto/request"
    "dz-jobs-api/internal/models"

    "github.com/google/uuid"
)

type JobService interface {
    PostNewJob(ctx context.Context, recruiterID uuid.UUID, req request.PostNewJobRequest) (*models.Job, error)
    GetJobDetails(ctx context.Context, jobID int64, recruiterID uuid.UUID) (*models.Job, error)
    GetJobListingsByStatus(ctx context.Context, status string, recruiterID uuid.UUID) ([]*models.Job, error)
    EditJob(ctx context.Context, jobID int64, req request.EditJobRequest, recruiterID uuid.UUID) (*models.Job, error)
    DeactivateJob(ctx context.Context, jobID int64, recruiterID uuid.UUID) (*models.Job, error)
    RepostJob(ctx context.Context, jobID int64, recruiterID uuid.UUID) (*models.Job, error)
    DeleteJob(ctx context.Context, jobID int64, recruiterID uuid.UUID) error
    GetAllJobs(ctx context.Context) ([]*models.Job, error)
    SearchJobs(ctx context.Context, filters request.JobFilters) ([]*models.Job, error)
    GetJobDetailsPublic(ctx context.Context, jobID int64) (*models.Job, error)
}
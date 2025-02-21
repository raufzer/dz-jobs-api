package interfaces

import (
	"context"
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/models"

	"github.com/google/uuid"
)

type JobRepository interface {
	CreateJob(ctx context.Context, job *models.Job) error
	GetJobDetails(ctx context.Context, jobID int64, recruiterID uuid.UUID) (*models.Job, error)
	GetJobListingsByStatus(ctx context.Context, status string, recruiterID uuid.UUID) ([]*models.Job, error)
	UpdateJob(ctx context.Context, jobID int64, recruiterID uuid.UUID, job *models.Job) error
	DeactivateJob(ctx context.Context, jobID int64, recruiterID uuid.UUID) error
	RepostJob(ctx context.Context, jobID int64, recruiterID uuid.UUID) error
	DeleteJob(ctx context.Context, jobID int64, recruiterID uuid.UUID) error
	ValidateJobOwnership(ctx context.Context, jobID int64, recruiterID uuid.UUID) error
	GetAllJobs(ctx context.Context, ) ([]*models.Job, error)
	GetJobListings(ctx context.Context, filters request.JobFilters) ([]*models.Job, error)
	GetJobDetailsPublic(ctx context.Context, jobID int64) (*models.Job, error)
}

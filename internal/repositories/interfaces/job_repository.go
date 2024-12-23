package interfaces

import (
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/models"

	"github.com/google/uuid"
)

type JobRepository interface {
	CreateJob(job *models.Job) error
	GetJobDetails(jobID int64, recruiterID uuid.UUID) (*models.Job, error)
	GetJobListingsByStatus(status string, recruiterID uuid.UUID) ([]*models.Job, error)
	UpdateJob(jobID int64, recruiterID uuid.UUID, job *models.Job) error
	DeactivateJob(jobID int64, recruiterID uuid.UUID) error
    RepostJob(jobID int64, recruiterID uuid.UUID) error
	DeleteJob(jobID int64, recruiterID uuid.UUID) error
	ValidateJobOwnership(jobID int64, recruiterID uuid.UUID) error
	GetAllJobs() ([]*models.Job, error)
	GetJobListings(filters request.JobFilters) ([]*models.Job, error)
	GetJobDetailsPublic(jobID int64) (*models.Job, error)
}

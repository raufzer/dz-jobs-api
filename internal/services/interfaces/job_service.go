package interfaces

import (
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/models"

	"github.com/google/uuid"
)

type JobService interface {
	PostNewJob(recruiterID uuid.UUID, req request.PostNewJobRequest) (*models.Job, error)
	GetJobDetails(jobID int64, recruiterID uuid.UUID) (*models.Job, error)
	GetJobListingsByStatus(status string, recruiterID uuid.UUID) ([]*models.Job, error)
	EditJob(jobID int64, req request.EditJobRequest, recruiterID uuid.UUID) (*models.Job, error)
	DeactivateJob(jobID int64, recruiterID uuid.UUID) (*models.Job, error)
	RepostJob(jobID int64, recruiterID uuid.UUID) (*models.Job, error)
	DeleteJob(jobID int64, recruiterID uuid.UUID) error
	GetAllJobs() ([]*models.Job, error)
	SearchJobs(filters request.JobFilters) ([]*models.Job, error)
	GetJobDetailsPublic(jobID int64) (*models.Job, error)
}

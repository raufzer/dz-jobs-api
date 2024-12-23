package interfaces

import (
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/models"

	"github.com/google/uuid"
)

type JobService interface {
	PostNewJob(recruiterID uuid.UUID, req request.PostNewJobRequest) (*models.Job, error)
	GetJobDetails(jobID int64) (*models.Job, error)
	GetJobListingsByStatus(status string) ([]*models.Job, error)
	EditJob(jobID int64, req request.EditJobRequest) (*models.Job, error)
	DeactivateJob(jobID int64) (*models.Job, error)
    RepostJob(jobID int64) (*models.Job, error)
	DeleteJob(jobID int64) error 
}

package interfaces

import (
	"dz-jobs-api/internal/models"

)

type JobRepository interface {
	CreateJob(job *models.Job) error
	GetJobDetails(jobID int64) (*models.Job, error)
	GetJobListingsByStatus(status string) ([]*models.Job, error)
	UpdateJob(jobID int64, job *models.Job) error
	DeactivateJob(jobID int64) error
    RepostJob(jobID int64) error
	DeleteJob(jobID int64) error
}

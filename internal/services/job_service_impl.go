package services

import (
	"database/sql"
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/models"
	"dz-jobs-api/internal/repositories/interfaces"
	"dz-jobs-api/pkg/utils"
	"log"
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

func (s *JobService) PostNewJob(recruiterID uuid.UUID, req request.PostNewJobRequest) (*models.Job, error) {
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
	}

	if err := s.jobRepository.CreateJob(job); err != nil {
		log.Println(err)
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Job posting failed")
	}

	return job, nil
}

func (s *JobService) GetJobDetails(jobID int64, recruiterID uuid.UUID) (*models.Job, error) {
	// Validate ownership
	job, err := s.jobRepository.GetJobDetails(jobID,recruiterID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NewCustomError(http.StatusNotFound, "Job not found")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Error fetching job details")
	}

	// Check if the recruiter is the owner of the job
	if job.RecruiterID != recruiterID {
		return nil, utils.NewCustomError(http.StatusForbidden, "You do not own this job")
	}

	return job, nil
}

func (s *JobService) GetJobListingsByStatus(status string, recruiterID uuid.UUID) ([]*models.Job, error) {
	// Fetch only jobs that the recruiter owns and match the status
	jobs, err := s.jobRepository.GetJobListingsByStatus(status, recruiterID)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to fetch jobs by status: "+status)
	}
	return jobs, nil
}

func (s *JobService) EditJob(jobID int64, req request.EditJobRequest, recruiterID uuid.UUID) (*models.Job, error) {
	// Validate ownership
	job, err := s.jobRepository.GetJobDetails(jobID,recruiterID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NewCustomError(http.StatusNotFound, "Job not found")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Error fetching job details")
	}

	// Check if the recruiter is the owner of the job
	if job.RecruiterID != recruiterID {
		return nil, utils.NewCustomError(http.StatusForbidden, "You do not own this job")
	}

	// Update the job details
	updatedJob := &models.Job{
		Title:          req.Title,
		Description:    req.Description,
		Location:       req.Location,
		SalaryRange:    req.SalaryRange,
		RequiredSkills: req.RequiredSkills,
		UpdatedAt:      time.Now(),
	}

	if err := s.jobRepository.UpdateJob(jobID, recruiterID, updatedJob); err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NewCustomError(http.StatusNotFound, "Job not found")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to update job")
	}

	return s.jobRepository.GetJobDetails(jobID,recruiterID)
}

func (s *JobService) DeactivateJob(jobID int64, recruiterID uuid.UUID) (*models.Job, error) {
	// Validate ownership
	job, err := s.jobRepository.GetJobDetails(jobID,recruiterID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NewCustomError(http.StatusNotFound, "Job not found")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Error fetching job details")
	}

	// Check if the recruiter is the owner of the job
	if job.RecruiterID != recruiterID {
		return nil, utils.NewCustomError(http.StatusForbidden, "You do not own this job")
	}

	// Deactivate the job
	if err := s.jobRepository.DeactivateJob(jobID,recruiterID); err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NewCustomError(http.StatusNotFound, "Job not found")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to deactivate job")
	}

	return s.jobRepository.GetJobDetails(jobID,recruiterID)
}

func (s *JobService) RepostJob(jobID int64, recruiterID uuid.UUID) (*models.Job, error) {
	// Validate ownership
	job, err := s.jobRepository.GetJobDetails(jobID,recruiterID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NewCustomError(http.StatusNotFound, "Job not found")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Error fetching job details")
	}

	// Check if the recruiter is the owner of the job
	if job.RecruiterID != recruiterID {
		return nil, utils.NewCustomError(http.StatusForbidden, "You do not own this job")
	}

	// Repost the job
	if err := s.jobRepository.RepostJob(jobID,recruiterID); err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NewCustomError(http.StatusNotFound, "Job not found")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to repost job")
	}

	return s.jobRepository.GetJobDetails(jobID,recruiterID)
}

func (s *JobService) DeleteJob(jobID int64, recruiterID uuid.UUID) error {
	// Validate ownership
	job, err := s.jobRepository.GetJobDetails(jobID, recruiterID)
	if err != nil {
		if err == sql.ErrNoRows {
			return utils.NewCustomError(http.StatusNotFound, "Job not found")
		}
		return utils.NewCustomError(http.StatusInternalServerError, "Error fetching job details")
	}

	// Check if the recruiter is the owner of the job
	if job.RecruiterID != recruiterID {
		return utils.NewCustomError(http.StatusForbidden, "You do not own this job")
	}

	// Delete the job
	err = s.jobRepository.DeleteJob(jobID,recruiterID)
	if err != nil {
		if err == sql.ErrNoRows {
			return utils.NewCustomError(http.StatusNotFound, "Job not found")
		}
		return utils.NewCustomError(http.StatusInternalServerError, "Failed to delete job")
	}
	return nil
}

func (s *JobService) GetAllJobs() ([]*models.Job, error) {
    jobs, err := s.jobRepository.GetAllJobs()
    if err != nil {
        return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to fetch all jobs")
    }
    return jobs, nil
}


func (s *JobService) SearchJobs(filters request.JobFilters) ([]*models.Job, error) {
    jobs, err := s.jobRepository.GetJobListings(filters)
    if err != nil {
        return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to search jobs")
    }
    return jobs, nil
}

func (s *JobService) GetJobDetailsPublic(jobID int64) (*models.Job, error) {
	job, err := s.jobRepository.GetJobDetailsPublic(jobID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NewCustomError(http.StatusNotFound, "Job not found")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Error fetching job details")
	}
	return job, nil
}
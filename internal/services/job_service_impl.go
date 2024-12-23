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

func (s *JobService) GetJobDetails(jobID int64) (*models.Job, error) {
	job, err := s.jobRepository.GetJobDetails(jobID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NewCustomError(http.StatusNotFound, "Job not found")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Error fetching job details")
	}
	return job, nil
}
func (s *JobService) GetJobListingsByStatus(status string) ([]*models.Job, error) {
	jobs, err := s.jobRepository.GetJobListingsByStatus(status)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to fetch jobs by status: "+status)
	}
	return jobs, nil
}
func (s *JobService) EditJob(jobID int64, req request.EditJobRequest) (*models.Job, error) {
	updatedJob := &models.Job{
		Title:          req.Title,
		Description:    req.Description,
		Location:       req.Location,
		SalaryRange:    req.SalaryRange,
		RequiredSkills: req.RequiredSkills,
		UpdatedAt:      time.Now(),
	}

	if err := s.jobRepository.UpdateJob(jobID, updatedJob); err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NewCustomError(http.StatusNotFound, "Job not found")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to update job")
	}

	return s.jobRepository.GetJobDetails(jobID)
}

func (s *JobService) DeactivateJob(jobID int64) (*models.Job, error) {


	if err := s.jobRepository.DeactivateJob(jobID); err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NewCustomError(http.StatusNotFound, "Job not found")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to desactivate job")
	}

	return s.jobRepository.GetJobDetails(jobID)
}

func (s *JobService) RepostJob(jobID int64) (*models.Job, error) {


	if err := s.jobRepository.RepostJob(jobID); err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NewCustomError(http.StatusNotFound, "Job not found")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to desactivate job")
	}

	return s.jobRepository.GetJobDetails(jobID)
}

func (s *JobService) DeleteJob(jobID int64) error {
	err := s.jobRepository.DeleteJob(jobID)
	if err != nil {
		if err == sql.ErrNoRows {
			return utils.NewCustomError(http.StatusNotFound, "Job not found")
		}
		return utils.NewCustomError(http.StatusInternalServerError, "Failed to delete job")
	}
	return nil
}
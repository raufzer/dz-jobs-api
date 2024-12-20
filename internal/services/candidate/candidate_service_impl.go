package candidate

import (
	"database/sql"
	request "dz-jobs-api/internal/dto/request/candidate"
	"dz-jobs-api/internal/helpers"
	models "dz-jobs-api/internal/models/candidate"
	interfaces "dz-jobs-api/internal/repositories/interfaces/candidate"
	"net/http"

	"github.com/google/uuid"
)

type CandidateService struct {
	candidateRepo interfaces.CandidateRepository
}

func NewCandidateService(repo interfaces.CandidateRepository) *CandidateService {
	return &CandidateService{candidateRepo: repo}
}

func (s *CandidateService) CreateCandidate(request request.CreateCandidateRequest) (*models.Candidate, error) {
	newCandidate := &models.Candidate{
		Resume:         request.Resume,
		ProfilePicture: request.ProfilePicture,
	}

	_, err := s.candidateRepo.CreateCandidate(*newCandidate)
	if err != nil {
		return nil, helpers.NewCustomError(http.StatusInternalServerError, "Failed to create candidate")
	}

	return newCandidate, nil
}

func (s *CandidateService) GetCandidateByID(candidateID uuid.UUID) (*models.Candidate, error) {
	candidate, err := s.candidateRepo.GetCandidateByID(candidateID)
	if err != nil {
		return nil, helpers.NewCustomError(http.StatusNotFound, "Candidate not found")
	}

	return &candidate, nil
}

func (s *CandidateService) UpdateCandidate(candidateID uuid.UUID, req request.UpdateCandidateRequest) (*models.Candidate, error) {
	updatedCandidate := &models.Candidate{
		Resume:         req.Resume,
		ProfilePicture: req.ProfilePicture,
	}

	if err := s.candidateRepo.UpdateCandidate(candidateID, *updatedCandidate); err != nil {
		if err == sql.ErrNoRows {
			return nil, helpers.NewCustomError(http.StatusNotFound, "Candidate not found")
		}
		return nil, helpers.NewCustomError(http.StatusInternalServerError, "Failed to update candidate")
	}

	return s.GetCandidateByID(candidateID)
}

func (s *CandidateService) DeleteCandidate(candidateID uuid.UUID) error {
	if err := s.candidateRepo.DeleteCandidate(candidateID); err != nil {
		return helpers.NewCustomError(http.StatusInternalServerError, "Failed to delete candidate")
	}

	return nil
}

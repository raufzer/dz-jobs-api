package services

import (
	"database/sql"
	"dz-jobs-api/config"
	"dz-jobs-api/internal/integrations"
	"dz-jobs-api/internal/models"
	"dz-jobs-api/internal/repositories/interfaces"
	"dz-jobs-api/pkg/utils"
	"mime/multipart"
	"net/http"

	"github.com/google/uuid"
)

type CandidateService struct {
	candidateRepo interfaces.CandidateRepository
	config        *config.AppConfig
}

func NewCandidateService(repo interfaces.CandidateRepository, config *config.AppConfig) *CandidateService {
	return &CandidateService{candidateRepo: repo,
		config: config}
}

func (s *CandidateService) CreateCandidate(userID string, profilePictureFile, resumeFile *multipart.FileHeader) (*models.Candidate, error) {
	existingCandidate, _ := s.candidateRepo.GetCandidate(uuid.MustParse(userID))
	if existingCandidate != nil {
		return nil, utils.NewCustomError(http.StatusBadRequest, "Candidate already exists")
	}
	if profilePictureFile == nil {
		return nil, utils.NewCustomError(http.StatusBadRequest, "Profile picture is required")
	}
	if resumeFile == nil {
		return nil, utils.NewCustomError(http.StatusBadRequest, "Resume is required")
	}
	profilePictureURL, err := integrations.UploadImage(profilePictureFile)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to upload profile picture")
	}
	resumeURL, err := integrations.UploadPDF(resumeFile)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to upload resume")
	}
	newCandidate := &models.Candidate{
		CandidateID:    uuid.MustParse(userID),
		Resume:         resumeURL,
		ProfilePicture: profilePictureURL,
	}

	_, err = s.candidateRepo.CreateCandidate(newCandidate)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to create candidate")
	}
	return newCandidate, nil
}

func (s *CandidateService) CreateDefaultCandidate(userID , resumeURL, profilePictureURL string) (*models.Candidate, error) {
	existingCandidate, _ := s.candidateRepo.GetCandidate(uuid.MustParse(userID))
	if existingCandidate != nil {
		return nil, utils.NewCustomError(http.StatusBadRequest, "Candidate already exists")
	}

	newCandidate := &models.Candidate{
		CandidateID:    uuid.MustParse(userID),
		Resume:         resumeURL,
		ProfilePicture: profilePictureURL,
	}

	_, err := s.candidateRepo.CreateCandidate(newCandidate)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to create candidate")
	}
	return newCandidate, nil
}

func (s *CandidateService) GetCandidate(candidateID uuid.UUID) (*models.Candidate, error) {
	candidate, err := s.candidateRepo.GetCandidate(candidateID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NewCustomError(http.StatusNotFound, "User not found")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Error fetching user")
	}

	return candidate, nil
}

func (s *CandidateService) UpdateCandidate(candidateID uuid.UUID, profilePictureFile, resumeFile *multipart.FileHeader) (*models.Candidate, error) {

	profilePictureURL, err := integrations.UploadImage(profilePictureFile)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to upload profile picture")
	}
	if profilePictureFile == nil {
		return nil, utils.NewCustomError(http.StatusBadRequest, "Profile picture is required")
	}

	resumeURL, err := integrations.UploadPDF(resumeFile)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to upload resume")
	}
	if resumeFile == nil {
		return nil, utils.NewCustomError(http.StatusBadRequest, "Resume is required")
	}

	updatedCandidate := &models.Candidate{
		Resume:         resumeURL,
		ProfilePicture: profilePictureURL,
	}

	if err := s.candidateRepo.UpdateCandidate(candidateID, updatedCandidate); err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NewCustomError(http.StatusNotFound, "Candidate not found")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to update candidate")
	}

	return s.candidateRepo.GetCandidate(candidateID)
}

func (s *CandidateService) DeleteCandidate(candidateID uuid.UUID) error {
	if err := s.candidateRepo.DeleteCandidate(candidateID); err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "Failed to delete candidate")
	}

	return nil
}


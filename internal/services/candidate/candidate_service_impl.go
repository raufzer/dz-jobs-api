package candidate

import (
	"database/sql"
	"dz-jobs-api/internal/integrations"
	models "dz-jobs-api/internal/models/candidate"
	interfaces "dz-jobs-api/internal/repositories/interfaces/candidate"
	"dz-jobs-api/pkg/utils"
	"mime/multipart"
	"net/http"

	"github.com/google/uuid"
)

type CandidateService struct {
	candidateRepo interfaces.CandidateRepository
}

func NewCandidateService(repo interfaces.CandidateRepository) *CandidateService {
	return &CandidateService{candidateRepo: repo}
}

func (s *CandidateService) CreateCandidate(profilePictureFile, resumeFile *multipart.FileHeader) (*models.Candidate, error) {
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
		Resume:         resumeURL,
		ProfilePicture: profilePictureURL,
	}

	_, err = s.candidateRepo.CreateCandidate(*newCandidate)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to create candidate")
	}

	return newCandidate, nil
}

func (s *CandidateService) GetCandidateByID(candidateID uuid.UUID) (*models.Candidate, error) {
	candidate, err := s.candidateRepo.GetCandidateByID(candidateID)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusNotFound, "Candidate not found")
	}

	return &candidate, nil
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

	if err := s.candidateRepo.UpdateCandidate(candidateID, *updatedCandidate); err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NewCustomError(http.StatusNotFound, "Candidate not found")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to update candidate")
	}

	return s.GetCandidateByID(candidateID)
}

func (s *CandidateService) DeleteCandidate(candidateID uuid.UUID) error {
	if err := s.candidateRepo.DeleteCandidate(candidateID); err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "Failed to delete candidate")
	}

	return nil
}

package services

import (
	"database/sql"
	"dz-jobs-api/config"
	"dz-jobs-api/internal/integrations"
	"dz-jobs-api/internal/models"
	"dz-jobs-api/internal/repositories/interfaces"
	"dz-jobs-api/pkg/utils"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type CandidateService struct {
	candidateRepo   interfaces.CandidateRepository
	redisRepository interfaces.RedisRepository
	config          *config.AppConfig
}

func NewCandidateService(repo interfaces.CandidateRepository, redisRepo interfaces.RedisRepository, config *config.AppConfig) *CandidateService {
	return &CandidateService{
		candidateRepo:   repo,
		redisRepository: redisRepo,
		config:          config,
	}
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

	profilePictureURL, err := s.uploadAndCacheFile(profilePictureFile, "image")
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to upload profile picture")
	}

	resumeURL, err := s.uploadAndCacheFile(resumeFile, "pdf")
	if err != nil {

		s.redisRepository.InvalidateAssetCache(profilePictureURL, "image")
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to upload resume")
	}

	newCandidate := &models.Candidate{
		CandidateID:    uuid.MustParse(userID),
		Resume:         resumeURL,
		ProfilePicture: profilePictureURL,
	}

	_, err = s.candidateRepo.CreateCandidate(newCandidate)
	if err != nil {

		s.redisRepository.InvalidateAssetCache(profilePictureURL, "image")
		s.redisRepository.InvalidateAssetCache(resumeURL, "pdf")
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to create candidate")
	}

	return newCandidate, nil
}

func (s *CandidateService) CreateDefaultCandidate(userID, resumeURL, profilePictureURL string) (*models.Candidate, error) {
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

func (s *CandidateService) uploadAndCacheFile(file *multipart.FileHeader, fileType string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	_, err = io.ReadAll(src)
	if err != nil {
		return "", err
	}

	var uploadURL string
	if fileType == "image" {
		uploadURL, err = integrations.UploadImage(file)
	} else {
		uploadURL, err = integrations.UploadPDF(file)
	}
	if err != nil {
		return "", err
	}

	assetCache := &utils.AssetCache{
		URL: uploadURL,
		Metadata: map[string]interface{}{
			"filename":   file.Filename,
			"size":       file.Size,
			"uploadedAt": time.Now(),
			"type":       fileType,
		},
		UpdatedAt: time.Now(),
	}

	err = s.redisRepository.StoreAssetCache(uploadURL, fileType, assetCache, 24*time.Hour)
	if err != nil {
		log.Printf("Failed to store file in cache: %v", err)
	}

	return uploadURL, nil
}

func (s *CandidateService) UpdateCandidate(candidateID uuid.UUID, profilePictureFile, resumeFile *multipart.FileHeader) (*models.Candidate, error) {

	existingCandidate, err := s.candidateRepo.GetCandidate(candidateID)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusNotFound, "Candidate not found")
	}

	if profilePictureFile == nil {
		return nil, utils.NewCustomError(http.StatusBadRequest, "Profile picture is required")
	}
	if resumeFile == nil {
		return nil, utils.NewCustomError(http.StatusBadRequest, "Resume is required")
	}

	profilePictureURL, err := s.uploadAndCacheFile(profilePictureFile, "image")
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to upload profile picture")
	}

	resumeURL, err := s.uploadAndCacheFile(resumeFile, "pdf")
	if err != nil {

		s.redisRepository.InvalidateAssetCache(profilePictureURL, "image")
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to upload resume")
	}

	updatedCandidate := &models.Candidate{
		Resume:         resumeURL,
		ProfilePicture: profilePictureURL,
	}

	if err := s.candidateRepo.UpdateCandidate(candidateID, updatedCandidate); err != nil {

		s.redisRepository.InvalidateAssetCache(profilePictureURL, "image")
		s.redisRepository.InvalidateAssetCache(resumeURL, "pdf")
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to update candidate")
	}

	s.redisRepository.InvalidateAssetCache(existingCandidate.ProfilePicture, "image")
	s.redisRepository.InvalidateAssetCache(existingCandidate.Resume, "pdf")

	return s.candidateRepo.GetCandidate(candidateID)
}

func (s *CandidateService) DeleteCandidate(candidateID uuid.UUID) error {

	candidate, err := s.candidateRepo.GetCandidate(candidateID)
	if err != nil {
		return utils.NewCustomError(http.StatusNotFound, "Candidate not found")
	}

	if err := s.candidateRepo.DeleteCandidate(candidateID); err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "Failed to delete candidate")
	}

	s.redisRepository.InvalidateAssetCache(candidate.ProfilePicture, "image")
	s.redisRepository.InvalidateAssetCache(candidate.Resume, "pdf")

	return nil
}

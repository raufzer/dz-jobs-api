package services

import (
	"database/sql"
	"dz-jobs-api/config"
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/integrations"
	"dz-jobs-api/internal/models"
	"dz-jobs-api/internal/repositories/interfaces"
	"dz-jobs-api/pkg/utils"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type RecruiterService struct {
	recruiterRepository interfaces.RecruiterRepository
	redisRepository     interfaces.RedisRepository
	config              *config.AppConfig
}

func NewRecruiterService(recruiterRepo interfaces.RecruiterRepository, redisRepo interfaces.RedisRepository, config *config.AppConfig) *RecruiterService {
	return &RecruiterService{
		recruiterRepository: recruiterRepo,
		redisRepository:     redisRepo,
		config:              config,
	}
}

func (s *RecruiterService) CreateRecruiter(userID string, req request.CreateRecruiterRequest, companyLogo *multipart.FileHeader) (*models.Recruiter, error) {
	existingRecruiter, err := s.recruiterRepository.GetRecruiter(uuid.MustParse(userID))
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to check recruiter existence")
	}
	if existingRecruiter != nil {
		return nil, utils.NewCustomError(http.StatusBadRequest, "Recruiter already exists")
	}
	if companyLogo == nil {
		return nil, utils.NewCustomError(http.StatusBadRequest, "Company Logo is required")
	}

	companyLogoURL, err := s.uploadAndCacheFile(companyLogo, "image")
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to upload company logo")
	}

	recruiter := &models.Recruiter{
		ID:                 uuid.MustParse(userID),
		CompanyName:        req.CompanyName,
		CompanyLogo:        companyLogoURL,
		CompanyDescription: req.CompanyDescription,
		CompanyWebsite:     req.CompanyWebsite,
		CompanyLocation:    req.CompanyLocation,
		CompanyContact:     req.CompanyContact,
		SocialLinks:        req.SocialLinks,
		VerifiedStatus:     req.VerifiedStatus,
	}

	if err := s.recruiterRepository.CreateRecruiter(recruiter); err != nil {
		if err := s.redisRepository.InvalidateAssetCache(companyLogoURL, "image"); err != nil {
			return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to invalidate asset cache")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Recruiter creation failed")
	}

	return recruiter, nil
}

func (s *RecruiterService) GetRecruiter(recruiterID uuid.UUID) (*models.Recruiter, error) {
	recruiter, err := s.recruiterRepository.GetRecruiter(recruiterID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NewCustomError(http.StatusNotFound, "REcruiter not found")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Error fetching recuiter")
	}
	return recruiter, nil
}

func (s *RecruiterService) UpdateRecruiter(recruiterID uuid.UUID, req request.UpdateRecruiterRequest, companyLogo *multipart.FileHeader) (*models.Recruiter, error) {
	existingRecruiter, err := s.recruiterRepository.GetRecruiter(recruiterID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NewCustomError(http.StatusNotFound, "Recruiter not found")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Error fetching recruiter")
	}

	if companyLogo == nil {
		return nil, utils.NewCustomError(http.StatusBadRequest, "Company Logo is required")
	}

	companyLogoURL, err := s.uploadAndCacheFile(companyLogo, "image")
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to upload company logo")
	}

	updatedRecruiter := &models.Recruiter{
		CompanyName:        req.CompanyName,
		CompanyLogo:        companyLogoURL,
		CompanyDescription: req.CompanyDescription,
		CompanyWebsite:     req.CompanyWebsite,
		CompanyLocation:    req.CompanyLocation,
		CompanyContact:     req.CompanyContact,
		SocialLinks:        req.SocialLinks,
		VerifiedStatus:     req.VerifiedStatus,
	}

	if err := s.recruiterRepository.UpdateRecruiter(recruiterID, updatedRecruiter); err != nil {
		if err := s.redisRepository.InvalidateAssetCache(companyLogoURL, "image"); err != nil {
			return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to invalidate asset cache")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to update Recruiter")
	}

	if err = s.redisRepository.InvalidateAssetCache(existingRecruiter.CompanyLogo, "image"); err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to invalidate asset cache")
	}
	return s.recruiterRepository.GetRecruiter(recruiterID)
}

func (s *RecruiterService) DeleteRecruiter(recruiterID uuid.UUID) error {
	recruiter, err := s.recruiterRepository.GetRecruiter(recruiterID)
	if err != nil {
		if err == sql.ErrNoRows {
			return utils.NewCustomError(http.StatusNotFound, "Recruiter not found")
		}
		return utils.NewCustomError(http.StatusInternalServerError, "Error fetching recruiter")
	}

	if err := s.recruiterRepository.DeleteRecruiter(recruiterID); err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "Failed to delete Recruiter")
	}

	if err = s.redisRepository.InvalidateAssetCache(recruiter.CompanyLogo, "image"); err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "Failed to invalidate asset cache")
	}
	return nil
}

func (s *RecruiterService) uploadAndCacheFile(file *multipart.FileHeader, fileType string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	uploadURL, err := integrations.UploadImage(file)
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
		return "", utils.NewCustomError(http.StatusInternalServerError, "Failed to cache asset")
	}

	return uploadURL, nil
}

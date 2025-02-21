package services

import (
	"context"
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

func (s *RecruiterService) CreateRecruiter(ctx context.Context, userID string, req request.CreateRecruiterRequest, companyLogo *multipart.FileHeader) (*models.Recruiter, error) {
	existingRecruiter, err := s.recruiterRepository.GetRecruiter(ctx,uuid.MustParse(userID))
	if err != nil {
		if err == sql.ErrNoRows {
			existingRecruiter = nil
		}
	}
	if existingRecruiter != nil {
		return nil, utils.NewCustomError(http.StatusBadRequest, "Recruiter already exists")
	} else {
		if companyLogo == nil {
			return nil, utils.NewCustomError(http.StatusBadRequest, "Company Logo is required")
		}

		companyLogoURL, err := s.uploadAndCacheFile(ctx,companyLogo, "image")
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

		if err := s.recruiterRepository.CreateRecruiter(ctx,recruiter); err != nil {
			if err := s.redisRepository.InvalidateAssetCache(ctx,companyLogoURL, "image"); err != nil {
				return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to invalidate asset cache")
			}
			return nil, utils.NewCustomError(http.StatusInternalServerError, "Recruiter creation failed")
		}

		return recruiter, nil
	}
}

func (s *RecruiterService) GetRecruiter(ctx context.Context, recruiterID uuid.UUID) (*models.Recruiter, error) {
	recruiter, err := s.recruiterRepository.GetRecruiter(ctx,recruiterID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NewCustomError(http.StatusNotFound, "REcruiter not found")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Error fetching recuiter")
	}
	return recruiter, nil
}

func (s *RecruiterService) UpdateRecruiter(ctx context.Context, recruiterID uuid.UUID, req request.UpdateRecruiterRequest, companyLogo *multipart.FileHeader) (*models.Recruiter, error) {
	existingRecruiter, err := s.recruiterRepository.GetRecruiter(ctx,recruiterID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NewCustomError(http.StatusNotFound, "Recruiter not found")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Error fetching recruiter")
	}

	if companyLogo == nil {
		return nil, utils.NewCustomError(http.StatusBadRequest, "Company Logo is required")
	}

	companyLogoURL, err := s.uploadAndCacheFile(ctx,companyLogo, "image")
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

	if err := s.recruiterRepository.UpdateRecruiter(ctx,recruiterID, updatedRecruiter); err != nil {
		if err := s.redisRepository.InvalidateAssetCache(ctx,companyLogoURL, "image"); err != nil {
			return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to invalidate asset cache")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to update Recruiter")
	}

	if err = s.redisRepository.InvalidateAssetCache(ctx,existingRecruiter.CompanyLogo, "image"); err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to invalidate asset cache")
	}
	return s.recruiterRepository.GetRecruiter(ctx,recruiterID)
}

func (s *RecruiterService) DeleteRecruiter(ctx context.Context, recruiterID uuid.UUID) error {
	recruiter, err := s.recruiterRepository.GetRecruiter(ctx,recruiterID)
	if err != nil {
		if err == sql.ErrNoRows {
			return utils.NewCustomError(http.StatusNotFound, "Recruiter not found")
		}
		return utils.NewCustomError(http.StatusInternalServerError, "Error fetching recruiter")
	}

	if err := s.recruiterRepository.DeleteRecruiter(ctx,recruiterID); err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "Failed to delete Recruiter")
	}

	if err = s.redisRepository.InvalidateAssetCache(ctx,recruiter.CompanyLogo, "image"); err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "Failed to invalidate asset cache")
	}
	return nil
}

func (s *RecruiterService) uploadAndCacheFile(ctx context.Context, file *multipart.FileHeader, fileType string) (string, error) {
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

	err = s.redisRepository.StoreAssetCache(ctx,uploadURL, fileType, assetCache, 24*time.Hour)
	if err != nil {
		return "", utils.NewCustomError(http.StatusInternalServerError, "Failed to cache asset")
	}

	return uploadURL, nil
}

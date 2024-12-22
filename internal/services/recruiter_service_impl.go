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

	"github.com/google/uuid"
)

type RecruiterService struct {
	recruiterRepository interfaces.RecruiterRepository
	config              *config.AppConfig
}

func NewRecruiterService(recruiterRepo interfaces.RecruiterRepository, config *config.AppConfig) *RecruiterService {
	return &RecruiterService{recruiterRepository: recruiterRepo,
		config: config}
}

func (rs *RecruiterService) CreateRecruiter(userID string, req request.CreateRecruiterRequest, companyLogo *multipart.FileHeader) (*models.Recruiter, error) {
	existingRecruiter, _ := rs.recruiterRepository.GetRecruiterByID(uuid.MustParse(userID))
	if existingRecruiter != nil {
		return nil, utils.NewCustomError(http.StatusBadRequest, "Recruiter already exists")
	}
	if companyLogo == nil {
		return nil, utils.NewCustomError(http.StatusBadRequest, "Company Logo is required")
	}

	companyLogoURL, err := integrations.UploadImage(companyLogo)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to upload profile picture")
	}
	recruiter := &models.Recruiter{
		RecruiterID:        uuid.MustParse(userID),
		CompanyName:        req.CompanyName,
		CompanyLogo:        companyLogoURL,
		CompanyDescription: req.CompanyDescription,
		CompanyWebsite:     req.CompanyWebsite,
		CompanyLocation:    req.CompanyLocation,
		CompanyContact:     req.CompanyContact,
		SocialLinks:        req.SocialLinks,
		VerifiedStatus:     req.VerifiedStatus,
	}

	if err := rs.recruiterRepository.CreateRecruiter(recruiter); err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Recruiter creation failed")
	}

	return recruiter, nil
}

func (rs *RecruiterService) GetRecruiterByID(recruiter_id uuid.UUID) (*models.Recruiter, error) {
	recruiter, err := rs.recruiterRepository.GetRecruiterByID(recruiter_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NewCustomError(http.StatusNotFound, "REcruiter not found")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Error fetching recuiter")
	}
	return recruiter, nil
}

func (rs *RecruiterService) UpdateRecruiter(recruiter_id uuid.UUID, req request.UpdateRecruiterRequest, companyLogo *multipart.FileHeader) (*models.Recruiter, error) {
	if companyLogo == nil {
		return nil, utils.NewCustomError(http.StatusBadRequest, "Company Logo is required")
	}

	companyLogoURL, err := integrations.UploadImage(companyLogo)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to upload profile picture")
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

	if err := rs.recruiterRepository.UpdateRecruiter(recruiter_id, updatedRecruiter); err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NewCustomError(http.StatusNotFound, "Recruiter not found")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to update Recruiter")
	}

	return rs.recruiterRepository.GetRecruiterByID(recruiter_id)
}

func (rs *RecruiterService) DeleteRecruiter(recruiter_id uuid.UUID) error {
	err := rs.recruiterRepository.DeleteRecruiter(recruiter_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return utils.NewCustomError(http.StatusNotFound, "Recruiter not found")
		}
		return utils.NewCustomError(http.StatusInternalServerError, "Failed to delete Recruiter")
	}
	return nil
}

func (rs *RecruiterService) ExtractTokenDetails(token string) (string, error) {

	claims, err := utils.ExtractTokenDetails(token, rs.config.AccessTokenSecret)
	if err != nil {
		return "", utils.NewCustomError(http.StatusUnauthorized, "Invalid or expired token")
	}
	return claims.UserID, nil
}

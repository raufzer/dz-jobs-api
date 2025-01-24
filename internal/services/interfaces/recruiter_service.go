package interfaces

import (
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/models"
	"mime/multipart"

	"github.com/google/uuid"
)

type RecruiterService interface {
	CreateRecruiter(userID string, req request.CreateRecruiterRequest, companyLogo *multipart.FileHeader) (*models.Recruiter, error)
	UpdateRecruiter(recruiterID uuid.UUID, req request.UpdateRecruiterRequest, companyLogo *multipart.FileHeader) (*models.Recruiter, error)
	GetRecruiter(recruiterID uuid.UUID) (*models.Recruiter, error)
	DeleteRecruiter(recruiterID uuid.UUID) error
}

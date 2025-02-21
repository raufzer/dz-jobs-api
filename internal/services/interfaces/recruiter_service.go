package interfaces

import (
    "context"
    "dz-jobs-api/internal/dto/request"
    "dz-jobs-api/internal/models"
    "mime/multipart"

    "github.com/google/uuid"
)

type RecruiterService interface {
    CreateRecruiter(ctx context.Context, userID string, req request.CreateRecruiterRequest, companyLogo *multipart.FileHeader) (*models.Recruiter, error)
    UpdateRecruiter(ctx context.Context, recruiterID uuid.UUID, req request.UpdateRecruiterRequest, companyLogo *multipart.FileHeader) (*models.Recruiter, error)
    GetRecruiter(ctx context.Context, recruiterID uuid.UUID) (*models.Recruiter, error)
    DeleteRecruiter(ctx context.Context, recruiterID uuid.UUID) error
}
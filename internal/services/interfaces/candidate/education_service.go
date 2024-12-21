package candidate

import (
	request "dz-jobs-api/internal/dto/request/candidate"
	models "dz-jobs-api/internal/models/candidate"

	"github.com/google/uuid"
)

type CandidateEducationService interface {
	AddEducation(userID string, request request.AddEducationRequest) (*models.CandidateEducation, error)
	GetEducationByCandidateID(candidateID uuid.UUID) ([]models.CandidateEducation, error)
	DeleteEducation(educationID uuid.UUID) error
	ExtractTokenDetails(token string) (string, error)
}

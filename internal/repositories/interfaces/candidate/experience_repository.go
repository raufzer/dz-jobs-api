package candidate
import (
	models"dz-jobs-api/internal/models/candidate"
	"github.com/google/uuid"
)


type CandidateExperienceRepository interface {
	CreateExperience(experience *models.CandidateExperience) error
	GetExperienceByCandidateID(id uuid.UUID) ([]models.CandidateExperience, error)
	UpdateExperience(experience *models.CandidateExperience) error
	DeleteExperience(id uuid.UUID) error
}

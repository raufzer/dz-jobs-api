package candidate
import (
	models"dz-jobs-api/internal/models/candidate"
	"github.com/google/uuid"
)


type CandidateExperienceRepository interface {
	CreateExperience(experience *models.CandidateExperience) error
	GetExperience(id uuid.UUID) ([]models.CandidateExperience, error)
	DeleteExperience(id uuid.UUID) error
}

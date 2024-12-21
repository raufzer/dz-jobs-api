package candidate

import (
	"dz-jobs-api/config"
	request "dz-jobs-api/internal/dto/request/candidate"
	models "dz-jobs-api/internal/models/candidate"
	interfaces "dz-jobs-api/internal/repositories/interfaces/candidate"
	"dz-jobs-api/pkg/utils"
	"net/http"


	"github.com/google/uuid"
)

type candidateEducationService struct {
	educationRepo interfaces.CandidateEducationRepository
	config        *config.AppConfig
}

func NewCandidateEducationService(repo interfaces.CandidateEducationRepository, config *config.AppConfig) *candidateEducationService {
	return &candidateEducationService{educationRepo: repo,
		config: config}	
}

func (s *candidateEducationService) AddEducation(userID string, request request.AddEducationRequest) (*models.CandidateEducation, error) {

	education := &models.CandidateEducation{
		EducationID: uuid.New(),
		CandidateID: uuid.MustParse(userID),
		Degree:      request.Degree,
		Institution: request.Institution,
		StartDate:   request.StartDate,
		EndDate:     request.EndDate,
		Description: request.Description,
	}

	err := s.educationRepo.CreateEducation(education)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to add education")
	}

	return education, nil
}

func (s *candidateEducationService) GetEducationByCandidateID(candidateID uuid.UUID) ([]models.CandidateEducation, error) {
	educations, err := s.educationRepo.GetEducationByCandidateID(candidateID)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusNotFound, "No education records found")
	}

	return educations, nil
}

func (s *candidateEducationService) DeleteEducation(educationID uuid.UUID) error {
	err := s.educationRepo.DeleteEducation(educationID)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "Failed to delete education")
	}

	return nil
}

func (s *candidateEducationService) ExtractTokenDetails(token string) (string, error) {

	claims, err := utils.ExtractTokenDetails(token, s.config.AccessTokenSecret)
	if err != nil {
		return "", utils.NewCustomError(http.StatusUnauthorized, "Invalid or expired token")
	}
	return claims.UserID, nil
}

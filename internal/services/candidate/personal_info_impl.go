package candidate

import (
	request "dz-jobs-api/internal/dto/request/candidate"
	"dz-jobs-api/internal/helpers"
	models "dz-jobs-api/internal/models/candidate"
	interfaces "dz-jobs-api/internal/repositories/interfaces/candidate"
	"net/http"

	"github.com/google/uuid"
)

type candidatePersonalInfoService struct {
	personalInfoRepo interfaces.CandidatePersonalInfoRepository
}

func NewCandidatePersonalInfoService(repo interfaces.CandidatePersonalInfoRepository) *candidatePersonalInfoService {
	return &candidatePersonalInfoService{personalInfoRepo: repo}
}

func (s *candidatePersonalInfoService) UpdatePersonalInfo(request request.UpdateCandidatePersonalInfoRequest) error {
	info := &models.CandidatePersonalInfo{
		CandidateID: request.CandidateID,
		Name:        request.Name,
		Email:       request.Email,
		Phone:       request.Phone,
		Address:     request.Address,
	}

	err := s.personalInfoRepo.UpdatePersonalInfo(*info)
	if err != nil {
		return helpers.NewCustomError(http.StatusInternalServerError, "Failed to update personal info")
	}

	return nil
}

func (s *candidatePersonalInfoService) GetPersonalInfo(candidateID uuid.UUID) (*models.CandidatePersonalInfo, error) {
	info, err := s.personalInfoRepo.GetPersonalInfoByCandidateID(candidateID)
	if err != nil {
		return nil, helpers.NewCustomError(http.StatusNotFound, "Personal info not found")
	}

	return &info, nil
}

package candidate

import (
	request "dz-jobs-api/internal/dto/request/candidate"
	models "dz-jobs-api/internal/models/candidate"
	interfaces "dz-jobs-api/internal/repositories/interfaces/candidate"
	"dz-jobs-api/pkg/utils"
	"net/http"

	"github.com/google/uuid"
)

type candidatePersonalInfoService struct {
	personalInfoRepo interfaces.CandidatePersonalInfoRepository
}

func NewCandidatePersonalInfoService(repo interfaces.CandidatePersonalInfoRepository) *candidatePersonalInfoService {
	return &candidatePersonalInfoService{personalInfoRepo: repo}
}

func (s *candidatePersonalInfoService) UpdatePersonalInfo(id uuid.UUID, request request.UpdateCandidatePersonalInfoRequest) (*models.CandidatePersonalInfo, error) {
	info := &models.CandidatePersonalInfo{
		Name:    request.Name,
		Email:   request.Email,
		Phone:   request.Phone,
		Address: request.Address,
	}

	err := s.personalInfoRepo.UpdatePersonalInfo(info)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to update personal info")
	}
	return s.GetPersonalInfo(id)
}

func (s *candidatePersonalInfoService) GetPersonalInfo(candidateID uuid.UUID) (*models.CandidatePersonalInfo, error) {
	info, err := s.personalInfoRepo.GetPersonalInfoByCandidateID(candidateID)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusNotFound, "Personal info not found")
	}

	return info, nil
}

func (s *candidatePersonalInfoService) CreatePersonalInfo(request request.CreateCandidatePersonalInfoRequest, candidateID uuid.UUID) (*models.CandidatePersonalInfo, error) {
	existingPersonalInfo, _ := s.personalInfoRepo.GetPersonalInfoByCandidateID(candidateID)

	if existingPersonalInfo != nil {
		return nil, utils.NewCustomError(http.StatusBadRequest, "Personal info already exists")
	}
	info := &models.CandidatePersonalInfo{
		CandidateID: candidateID,
		Name:        request.Name,
		Email:       request.Email,
		Phone:       request.Phone,
		Address:     request.Address,
	}

	err := s.personalInfoRepo.CreatePersonalInfo(info)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to create personal info")
	}
	return s.GetPersonalInfo(candidateID)
}

func (s *candidatePersonalInfoService) DeletePersonalInfo(candidateID uuid.UUID) error {
	err := s.personalInfoRepo.DeletePersonalInfo(candidateID)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "Failed to delete personal info")
	}
	return nil
}

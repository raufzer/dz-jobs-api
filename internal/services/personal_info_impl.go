package services

import (
	"database/sql"
	"dz-jobs-api/internal/dto/request"
	"dz-jobs-api/internal/models"
	"dz-jobs-api/internal/repositories/interfaces"
	"dz-jobs-api/pkg/utils"
	"net/http"

	"github.com/google/uuid"
)

type CandidatePersonalInfoService struct {
	candidatePersonalInfoRepo interfaces.CandidatePersonalInfoRepository
}

func NewCandidatePersonalInfoService(repo interfaces.CandidatePersonalInfoRepository) *CandidatePersonalInfoService {
	return &CandidatePersonalInfoService{candidatePersonalInfoRepo: repo}
}

func (s *CandidatePersonalInfoService) UpdatePersonalInfo(candidateID uuid.UUID, request request.UpdatePersonalInfoRequest) (*models.CandidatePersonalInfo, error) {
	info := &models.CandidatePersonalInfo{
		ID:          candidateID,
		Name:        request.Name,
		Email:       request.Email,
		Phone:       request.Phone,
		Address:     request.Address,
		DateOfBirth: request.DateOfBirth,
		Gender:      request.Gender,
		Bio:         request.Bio,
	}

	err := s.candidatePersonalInfoRepo.UpdatePersonalInfo(info)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to update personal info")
	}
	return s.GetPersonalInfo(candidateID)
}

func (s *CandidatePersonalInfoService) GetPersonalInfo(candidateID uuid.UUID) (*models.CandidatePersonalInfo, error) {
	info, err := s.candidatePersonalInfoRepo.GetPersonalInfo(candidateID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NewCustomError(http.StatusNotFound, "Personal info not found")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to fetch personal info")
	}

	return info, nil
}

func (s *CandidatePersonalInfoService) AddPersonalInfo(request request.AddPersonalInfoRequest, candidateID uuid.UUID) (*models.CandidatePersonalInfo, error) {
	info := &models.CandidatePersonalInfo{
		ID:          candidateID,
		Name:        request.Name,
		Email:       request.Email,
		Phone:       request.Phone,
		Address:     request.Address,
		DateOfBirth: request.DateOfBirth,
		Gender:      request.Gender,
		Bio:         request.Bio,
	}

	err := s.candidatePersonalInfoRepo.CreatePersonalInfo(info)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to create personal info")
	}
	return s.GetPersonalInfo(candidateID)
}

func (s *CandidatePersonalInfoService) DeletePersonalInfo(candidateID uuid.UUID) error {
	err := s.candidatePersonalInfoRepo.DeletePersonalInfo(candidateID)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "Failed to delete personal info")
	}
	return nil
}

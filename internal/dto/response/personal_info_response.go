package response

import (
	"dz-jobs-api/internal/models"

	"github.com/google/uuid"
)

type PersonalInfoResponse struct {
	CandidateID uuid.UUID `json:"candidate_id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Address     string    `json:"address"`
	DateOfBirth string    `json:"date_of_birth"`
	Gender      string    `json:"gender"`
	Bio         string    `json:"bio"`
}

func ToPersonalInfoResponse(info *models.CandidatePersonalInfo) PersonalInfoResponse {
	return PersonalInfoResponse{
		CandidateID: info.CandidateID,
		Name:        info.Name,
		Email:       info.Email,
		Phone:       info.Phone,
		Address:     info.Address,
		DateOfBirth: info.DateOfBirth,
		Gender:      info.Gender,
		Bio:         info.Bio,
	}
}

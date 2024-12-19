package candidate

import (
	models "dz-jobs-api/internal/models/candidate"

	"github.com/google/uuid"
)

type PersonalInfoResponse struct {
	CandidateID uuid.UUID `json:"candidate_id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Address     string    `json:"address"`
}


func ToPersonalInfoResponse(info *models.CandidatePersonalInfo) PersonalInfoResponse {
    return PersonalInfoResponse{
        CandidateID: info.CandidateID,
        Name:        info.Name,
        Email:       info.Email,
        Phone:       info.Phone,
        Address:     info.Address,
    }
}

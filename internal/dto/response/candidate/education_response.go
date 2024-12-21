package candidate

import (
	models "dz-jobs-api/internal/models/candidate"

	"github.com/google/uuid"
)

type EducationResponse struct {
	EducationID uuid.UUID `json:"education_id"`
	CandidateID uuid.UUID `json:"candidate_id"`
	Degree      string    `json:"degree"`
	Institution string    `json:"institution"`
	StartDate   string    `json:"start_date"`
	EndDate     string    `json:"end_date"`
	Description string    `json:"description"`
}

func ToEducationResponse(education *models.CandidateEducation) EducationResponse {
	return EducationResponse{
		EducationID: education.EducationID,
		CandidateID: education.CandidateID,
		Degree:      education.Degree,
		Institution: education.Institution,
		StartDate:   education.StartDate,
		EndDate:     education.EndDate,
		Description: education.Description,
	}
}

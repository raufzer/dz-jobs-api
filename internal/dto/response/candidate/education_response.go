package response

import (
	models "dz-jobs-api/internal/models/candidate"
	"time"

	"github.com/google/uuid"
)

type EducationResponse struct {
	EducationID uuid.UUID `json:"education_id"`
	CandidateID uuid.UUID `json:"candidate_id"`
	Degree      string    `json:"degree"`
	Institution string    `json:"institution"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
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

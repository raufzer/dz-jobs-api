package candidate

import (
	models "dz-jobs-api/internal/models/candidate"

	"github.com/google/uuid"
)

type CertificationResponse struct {
	CertificationID   uuid.UUID `json:"certification_id"`
	CandidateID       uuid.UUID `json:"candidate_id"`
	CertificationName string    `json:"certification_name"`
	IssuedBy          string    `json:"issued_by"`
	IssueDate         string    `json:"issue_date"`
	ExpirationDate    string    `json:"expiration_date"`
}

func ToCertificationResponse(certification *models.CandidateCertification) CertificationResponse {
	return CertificationResponse{
		CertificationID:   certification.CertificationID,
		CandidateID:       certification.CandidateID,
		CertificationName: certification.CertificationName,
		IssuedBy:          certification.IssuedBy,
		IssueDate:         certification.IssueDate,
		ExpirationDate:    certification.ExpirationDate,
	}
}

type CertificationsResponseData struct {
	Total          int                     `json:"total"`
	Certifications []CertificationResponse `json:"certifications"`
}

func ToCertificationsResponse(certifications []models.CandidateCertification) CertificationsResponseData {
	var certificationResponses []CertificationResponse
	for _, cer := range certifications {
		certificationResponses = append(certificationResponses, ToCertificationResponse(&cer))
	}
	return CertificationsResponseData{
		Total:          len(certifications),
		Certifications: certificationResponses,
	}
}

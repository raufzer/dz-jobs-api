package interfaces

import (
	"dz-jobs-api/internal/models"

	"github.com/google/uuid"
)

type BookmarksRepository interface {
	AddBookmark(candidateID uuid.UUID, jobID int64) error
	RemoveBookmark(candidateID uuid.UUID, jobID int64) error
	GetBookmarks(candidateID uuid.UUID) ([]*models.Job, error)
}

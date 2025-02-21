package interfaces

import (
	"context"
	"dz-jobs-api/internal/models"

	"github.com/google/uuid"
)

type BookmarksRepository interface {
	AddBookmark(ctx context.Context, candidateID uuid.UUID, jobID int64) error
	RemoveBookmark(ctx context.Context, candidateID uuid.UUID, jobID int64) error
	GetBookmarks(ctx context.Context, candidateID uuid.UUID) ([]*models.Job, error)
}

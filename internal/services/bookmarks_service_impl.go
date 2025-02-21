package services

import (
	"context"
	"database/sql"
	"dz-jobs-api/internal/models"
	"dz-jobs-api/internal/repositories/interfaces"
	"dz-jobs-api/pkg/utils"
	"net/http"

	"github.com/google/uuid"
)

type BookmarksService struct {
	bookmarksRepository interfaces.BookmarksRepository
}

func NewBookmarksService(bookmarksRepo interfaces.BookmarksRepository) *BookmarksService {
	return &BookmarksService{bookmarksRepository: bookmarksRepo}
}

func (s *BookmarksService) AddBookmark(ctx context.Context, candidateID uuid.UUID, jobID int64) error {
	err := s.bookmarksRepository.AddBookmark(ctx, candidateID, jobID)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "Error adding Bookmark")
	}
	return nil
}

func (s *BookmarksService) RemoveBookmark(ctx context.Context, candidateID uuid.UUID, jobID int64) error {
	err := s.bookmarksRepository.RemoveBookmark(ctx, candidateID, jobID)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "Error removing Bookmark")
	}
	return nil
}

func (s *BookmarksService) GetBookmarks(ctx context.Context, candidateID uuid.UUID) ([]*models.Job, error) {
	bookmarks, err := s.bookmarksRepository.GetBookmarks(ctx, candidateID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return bookmarks, nil
}

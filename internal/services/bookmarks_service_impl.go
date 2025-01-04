package services

import (
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
func (s *BookmarksService) AddBookmark(candidateID uuid.UUID, jobID int64) error {
	err := s.bookmarksRepository.AddBookmark(candidateID, jobID)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "Error adding Bookmark")
	}
	return nil
}

func (s *BookmarksService) RemoveBookmark(candidateID uuid.UUID, jobID int64) error {
	err := s.bookmarksRepository.RemoveBookmark(candidateID, jobID)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "Error removing Bookmark")
	}
	return nil
}

func (s *BookmarksService) GetBookmarks(candidateID uuid.UUID) ([]*models.Job, error) {
	bookmarks, err := s.bookmarksRepository.GetBookmarks(candidateID)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Error fetching Bookmarks")
	}
	return bookmarks, nil
}

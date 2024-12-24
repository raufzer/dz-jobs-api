package services

import (
	"dz-jobs-api/internal/models"
	"dz-jobs-api/internal/repositories/interfaces"

	"github.com/google/uuid"
)

type BookmarksService struct {
	bookmarksRepository interfaces.BookmarksRepository
}

func NewBookmarksService(bookmarksRepo interfaces.BookmarksRepository) *BookmarksService {
	return &BookmarksService{bookmarksRepository: bookmarksRepo}
}
func (s *BookmarksService) AddBookmark(candidateID uuid.UUID, jobID int64) error {
	return s.bookmarksRepository.AddBookmark(candidateID, jobID)
}

func (s *BookmarksService) RemoveBookmark(candidateID uuid.UUID, jobID int64) error {
	return s.bookmarksRepository.RemoveBookmark(candidateID, jobID)
}

func (s *BookmarksService) GetBookmarks(candidateID uuid.UUID) ([]*models.Job, error) {
	{
		return s.bookmarksRepository.GetBookmarks(candidateID)
	}
}

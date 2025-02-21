package services

import (
    "context"
    "database/sql"
    "dz-jobs-api/config"
    "dz-jobs-api/internal/integrations"
    "dz-jobs-api/internal/models"
    "dz-jobs-api/internal/repositories/interfaces"
    "dz-jobs-api/pkg/utils"
    "io"
    "mime/multipart"
    "net/http"
    "time"

    "github.com/google/uuid"

)

type CandidateService struct {
    candidateRepo   interfaces.CandidateRepository
    redisRepository interfaces.RedisRepository
    config          *config.AppConfig
}

func NewCandidateService(repo interfaces.CandidateRepository, redisRepo interfaces.RedisRepository, config *config.AppConfig) *CandidateService {
    return &CandidateService{
        candidateRepo:   repo,
        redisRepository: redisRepo,
        config:          config,
    }
}

func (s *CandidateService) CreateCandidate(ctx context.Context, userID string, profilePictureFile, resumeFile *multipart.FileHeader) (*models.Candidate, error) {
    existingCandidate, err := s.candidateRepo.GetCandidate(ctx, uuid.MustParse(userID))
    if err != nil {
        if err == sql.ErrNoRows {
            existingCandidate = nil
        }
    }
    if existingCandidate != nil {
        return nil, utils.NewCustomError(http.StatusBadRequest, "Candidate already exists")
    } else {

        if profilePictureFile == nil {
            return nil, utils.NewCustomError(http.StatusBadRequest, "Profile picture is required")
        }
        if resumeFile == nil {
            return nil, utils.NewCustomError(http.StatusBadRequest, "Resume is required")
        }

        profilePictureURL, err := s.uploadAndCacheFile(ctx, profilePictureFile, "image")
        if err != nil {
            return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to upload profile picture")
        }

        resumeURL, err := s.uploadAndCacheFile(ctx, resumeFile, "pdf")
        if err != nil {

            if err := s.redisRepository.InvalidateAssetCache(ctx, profilePictureURL, "image"); err != nil {
                return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to invalidate asset cache")
            }
            return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to upload resume")
        }

        newCandidate := &models.Candidate{
            ID:             uuid.MustParse(userID),
            Resume:         resumeURL,
            ProfilePicture: profilePictureURL,
        }

        _, err = s.candidateRepo.CreateCandidate(ctx, newCandidate)
        if err != nil {

            if err := s.redisRepository.InvalidateAssetCache(ctx, profilePictureURL, "image"); err != nil {
                return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to invalidate asset cache")
            }
            if err := s.redisRepository.InvalidateAssetCache(ctx, resumeURL, "pdf"); err != nil {
                return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to invalidate asset cache")
            }
            return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to create candidate")
        }

        return newCandidate, nil
    }
}

func (s *CandidateService) CreateDefaultCandidate(ctx context.Context, userID, resumeURL, profilePictureURL string) (*models.Candidate, error) {
    existingCandidate, err := s.candidateRepo.GetCandidate(ctx, uuid.MustParse(userID))
    if err != nil {
        if err == sql.ErrNoRows {
            existingCandidate = nil
        }
    }
    if existingCandidate != nil {
        return nil, utils.NewCustomError(http.StatusBadRequest, "Candidate already exists")
    } else {

        newCandidate := &models.Candidate{
            ID:             uuid.MustParse(userID),
            Resume:         resumeURL,
            ProfilePicture: profilePictureURL,
        }

        _, err = s.candidateRepo.CreateCandidate(ctx, newCandidate)
        if err != nil {
            return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to create candidate")
        }
        return newCandidate, nil
    }
}

func (s *CandidateService) GetCandidate(ctx context.Context, candidateID uuid.UUID) (*models.Candidate, error) {
    candidate, err := s.candidateRepo.GetCandidate(ctx, candidateID)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, utils.NewCustomError(http.StatusNotFound, "Candidate not found")
        }
        return nil, utils.NewCustomError(http.StatusInternalServerError, "Error fetching Candidate")
    }

    return candidate, nil
}

func (s *CandidateService) uploadAndCacheFile(ctx context.Context, file *multipart.FileHeader, fileType string) (string, error) {
    src, err := file.Open()
    if err != nil {
        return "", err
    }
    defer src.Close()

    _, err = io.ReadAll(src)
    if err != nil {
        return "", err
    }

    var uploadURL string
    if fileType == "image" {
        uploadURL, err = integrations.UploadImage(file)
    } else {
        uploadURL, err = integrations.UploadPDF(file)
    }
    if err != nil {
        return "", err
    }

    assetCache := &utils.AssetCache{
        URL: uploadURL,
        Metadata: map[string]interface{}{
            "filename":   file.Filename,
            "size":       file.Size,
            "uploadedAt": time.Now(),
            "type":       fileType,
        },
        UpdatedAt: time.Now(),
    }

    err = s.redisRepository.StoreAssetCache(ctx, uploadURL, fileType, assetCache, 24*time.Hour)
    if err != nil {
        return "", utils.NewCustomError(http.StatusInternalServerError, "Failed to cache asset")
    }

    return uploadURL, nil
}

func (s *CandidateService) UpdateCandidate(ctx context.Context, candidateID uuid.UUID, profilePictureFile, resumeFile *multipart.FileHeader) (*models.Candidate, error) {

    existingCandidate, err := s.candidateRepo.GetCandidate(ctx, candidateID)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, utils.NewCustomError(http.StatusNotFound, "Candidate not found")
        }
        return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to fetch candidate")
    }

    if profilePictureFile == nil {
        return nil, utils.NewCustomError(http.StatusBadRequest, "Profile picture is required")
    }
    if resumeFile == nil {
        return nil, utils.NewCustomError(http.StatusBadRequest, "Resume is required")
    }

    profilePictureURL, err := s.uploadAndCacheFile(ctx, profilePictureFile, "image")
    if err != nil {
        return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to upload profile picture")
    }

    resumeURL, err := s.uploadAndCacheFile(ctx, resumeFile, "pdf")
    if err != nil {

        if err := s.redisRepository.InvalidateAssetCache(ctx, profilePictureURL, "image"); err != nil {
            return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to invalidate asset cache")
        }
        return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to upload resume")
    }

    updatedCandidate := &models.Candidate{
        Resume:         resumeURL,
        ProfilePicture: profilePictureURL,
    }

    if err := s.candidateRepo.UpdateCandidate(ctx, candidateID, updatedCandidate); err != nil {

        if err := s.redisRepository.InvalidateAssetCache(ctx, profilePictureURL, "image"); err != nil {
            return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to invalidate asset cache")
        }
        if err := s.redisRepository.InvalidateAssetCache(ctx, resumeURL, "pdf"); err != nil {
            return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to invalidate asset cache")
        }
        return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to update candidate")
    }

    if err := s.redisRepository.InvalidateAssetCache(ctx, existingCandidate.ProfilePicture, "image"); err != nil {
        return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to invalidate asset cache")
    }
    if err := s.redisRepository.InvalidateAssetCache(ctx, existingCandidate.Resume, "pdf"); err != nil {
        return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to invalidate asset cache")
    }

    return s.candidateRepo.GetCandidate(ctx, candidateID)
}

func (s *CandidateService) DeleteCandidate(ctx context.Context, candidateID uuid.UUID) error {

    candidate, err := s.candidateRepo.GetCandidate(ctx, candidateID)
    if err != nil {
        if err == sql.ErrNoRows {
            return utils.NewCustomError(http.StatusNotFound, "Candidate not found")
        }
        return utils.NewCustomError(http.StatusInternalServerError, "Failed to fetch candidate")
    }

    if err := s.candidateRepo.DeleteCandidate(ctx, candidateID); err != nil {
        return utils.NewCustomError(http.StatusInternalServerError, "Failed to delete candidate")
    }

    if err := s.redisRepository.InvalidateAssetCache(ctx, candidate.ProfilePicture, "image"); err != nil {
        return utils.NewCustomError(http.StatusInternalServerError, "Failed to invalidate asset cache")
    }
    if err := s.redisRepository.InvalidateAssetCache(ctx, candidate.Resume, "pdf"); err != nil {
        return utils.NewCustomError(http.StatusInternalServerError, "Failed to invalidate asset cache")
    }

    return nil
}
package integrations

import (
	"context"
	"dz-jobs-api/config"
	"fmt"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

var cld *cloudinary.Cloudinary

func Init(cfg *config.AppConfig) {
	var err error
	cld, err = cloudinary.NewFromParams(cfg.CloudinaryCloudName, cfg.CloudinaryAPIKey, cfg.CloudinaryAPISecret)
	if err != nil {
		panic(fmt.Errorf("failed to initialize Cloudinary: %v", err))
	}
}

func UploadImage(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open image file: %w", err)
	}
	if file == nil {
		return "", fmt.Errorf("file is nil")
	}
	defer file.Close()

	uniqueID := uuid.New().String()
	publicID := fmt.Sprintf("pdfs/%s", uniqueID)

	ctx := context.Background()
	result, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		ResourceType: "image",
		PublicID:     publicID,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload image to Cloudinary: %w", err)
	}

	if result == nil {
		return "", fmt.Errorf("received nil result from Cloudinary upload")
	}

	return result.URL, nil
}

func UploadPDF(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open PDF file: %w", err)
	}
	defer file.Close()

	uniqueID := uuid.New().String()
	publicID := fmt.Sprintf("pdfs/%s", uniqueID)

	ctx := context.Background()
	result, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		ResourceType: "raw",
		PublicID:     publicID,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload PDF: %w", err)
	}

	return result.URL, nil
}

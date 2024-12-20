package integrations

import (
	"context"
	"dz-jobs-api/config"
	"fmt"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

var cld *cloudinary.Cloudinary

func Init(cfg *config.AppConfig) {
	var err error
	cloudinaryURL := "cloudinary://" + cfg.CloudinaryAPIKey + ":" + cfg.CloudinaryAPISecret + "@" + cfg.CloudinaryCloudName
	cld, err = cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		panic(fmt.Errorf("failed to initialize Cloudinary: %w", err))
	}
}

func UploadImage(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open image file: %w", err)
	}
	defer file.Close()

	ctx := context.Background()
	result, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		ResourceType: "image",
		PublicID:     fileHeader.Filename,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload image: %w", err)
	}

	return result.URL, nil
}

func UploadPDF(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open PDF file: %w", err)
	}
	defer file.Close()

	ctx := context.Background()
	result, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		ResourceType: "raw",
		PublicID:     fileHeader.Filename,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload PDF: %w", err)
	}

	return result.URL, nil
}

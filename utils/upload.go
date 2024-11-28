package utils

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/config"
)

func UploadToGCS(fileHeader *multipart.FileHeader, objectName string) (string, error) {
	cfg := config.GetConfig()
	ctx := context.Background()

	client := GCSClient
	if client == nil {
		return "", fmt.Errorf("GCS client is not initialized")
	}

	bucketName := cfg.GCS_BUCKET_NAME
	bucket := client.Bucket(bucketName)
	object := bucket.Object(objectName)

	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	writer := object.NewWriter(ctx)
	defer writer.Close()

	if _, err := io.Copy(writer, file); err != nil {
		return "", fmt.Errorf("failed to write file to GCS: %w", err)
	}

	return fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, objectName), nil
}

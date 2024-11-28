package utils

import (
	"context"
	"log"

	"cloud.google.com/go/storage"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/config"
	"google.golang.org/api/option"
)

var (
	GCSClient *storage.Client
)

func InitGCS() {
	cfg := config.GetConfig()
	ctx := context.Background()

	client, err := storage.NewClient(ctx, option.WithCredentialsFile(cfg.GOOGLE_APPLICATION_CREDENTIALS))
	if err != nil {
		log.Fatalf("Failed to initialize GCS client: %v", err)
	}

	GCSClient = client
	log.Println("GCS client initialized successfully with explicit credentials")
}

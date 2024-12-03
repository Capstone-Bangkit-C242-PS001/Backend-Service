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

	var client *storage.Client
	var err error

	// Check if the environment is "development"
	if cfg.APP_ENV == "development" {
		client, err = storage.NewClient(ctx, option.WithCredentialsFile(cfg.GOOGLE_APPLICATION_CREDENTIALS))
		if err != nil {
			log.Fatalf("Failed to initialize GCS client with credentials file: %v", err)
		}
		log.Println("GCS client initialized successfully with explicit credentials (development mode)")
	} else {
		client, err = storage.NewClient(ctx)
		if err != nil {
			log.Fatalf("Failed to initialize GCS client with default credentials: %v", err)
		}
		log.Println("GCS client initialized successfully with default credentials (production mode)")
	}

	GCSClient = client
}

package config

import (
	"log"
	"os"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	APP_PORT                       string
	DB_USER                        string
	DB_PASSWORD                    string
	DB_NAME                        string
	DB_HOST                        string
	APP_ENV                        string
	DB_PORT                        string
	GOOGLE_APPLICATION_CREDENTIALS string
	GCS_BUCKET_NAME                string
	ML_SERVICE_BASE_URL            string
}

var (
	appConfig *Config
	once      sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		viper.AutomaticEnv() // Prioritize environment variables

		// Optional: Load .env file for local development
		if _, err := os.Stat(".env"); err == nil {
			viper.SetConfigFile(".env")
			viper.SetConfigType("env")

			if err := viper.ReadInConfig(); err != nil {
				log.Printf("Error reading .env file: %v", err)
			}
		}

		// Bind environment variables (fallback to Viper's env)
		_ = viper.BindEnv("APP_PORT")
		_ = viper.BindEnv("DB_USER")
		_ = viper.BindEnv("DB_PASSWORD")
		_ = viper.BindEnv("DB_NAME")
		_ = viper.BindEnv("DB_HOST")
		_ = viper.BindEnv("APP_ENV")
		_ = viper.BindEnv("DB_PORT")
		_ = viper.BindEnv("GOOGLE_APPLICATION_CREDENTIALS")
		_ = viper.BindEnv("GCS_BUCKET_NAME")
		_ = viper.BindEnv("ML_SERVICE_BASE_URL")

		// Unmarshal configuration into struct
		appConfig = &Config{
			APP_PORT:                       viper.GetString("APP_PORT"),
			DB_USER:                        viper.GetString("DB_USER"),
			DB_PASSWORD:                    viper.GetString("DB_PASSWORD"),
			DB_NAME:                        viper.GetString("DB_NAME"),
			DB_HOST:                        viper.GetString("DB_HOST"),
			APP_ENV:                        viper.GetString("APP_ENV"),
			DB_PORT:                        viper.GetString("DB_PORT"),
			GOOGLE_APPLICATION_CREDENTIALS: viper.GetString("GOOGLE_APPLICATION_CREDENTIALS"),
			GCS_BUCKET_NAME:                viper.GetString("GCS_BUCKET_NAME"),
			ML_SERVICE_BASE_URL:            viper.GetString("ML_SERVICE_BASE_URL"),
		}

		// Set default port if not provided
		if appConfig.APP_PORT == "" {
			appConfig.APP_PORT = "9000"
		}
	})

	return appConfig
}

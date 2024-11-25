package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	APP_PORT    string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	APP_ENV     string
	DB_PORT     string
}

var (
	appConfig *Config
	once      sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		viper.AddConfigPath(".")
		viper.SetConfigName(".env")
		viper.SetConfigType("env")

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("No .env file found: %v", err)
		}

		if err := viper.Unmarshal(&appConfig); err != nil {
			log.Fatalf("Failed to unmarshall configuration: %v", err)
		}
	})

	return appConfig
}

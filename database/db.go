package database

import (
	"fmt"
	"log"

	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadDB() {
	cfg := config.GetConfig()

	var dsn string
	if cfg.APP_ENV == "development" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)
	} else {
		SOCKET_DIR := "/cloudsql"
		dsn = fmt.Sprintf("%s:%s@unix(/%s/%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.DB_USER, cfg.DB_PASSWORD, SOCKET_DIR, cfg.DB_HOST, cfg.DB_NAME)
	}

	log.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to Connect Database: %v", err)
	}

	DB = db
	fmt.Println("Database connection established successfully.")
}

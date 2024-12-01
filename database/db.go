package database

import (
	"errors"
	"fmt"
	"log"

	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
		dsn = fmt.Sprintf("%s:%s@unix(%s/%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.DB_USER,     // Database username
			cfg.DB_PASSWORD, // Database password
			SOCKET_DIR,      // Correct Cloud SQL Unix socket directory
			cfg.DB_HOST,     // Cloud SQL instance connection name
			cfg.DB_NAME,     // Database name
		)
	}

	log.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to Connect Database: %v", err)
	}

	DB = db
	fmt.Println("Database connection established successfully.")

	runMigrations(dsn)
}

func runMigrations(dsn string) {
	log.Println("Running database migrations...")

	// Format DSN for golang-migrate
	migrationDSN := fmt.Sprintf("mysql://%s", dsn)

	// Path to migration files
	migrationPath := "file://./database/migration"

	m, err := migrate.New(migrationPath, migrationDSN)
	if err != nil {
		log.Fatalf("Failed to initialize migration: %v", err)
	}

	// Apply migrations
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("No new migrations to apply.")
		} else {
			log.Fatalf("Failed to apply migrations: %v", err)
		}
		return
	}

	log.Println("Database migrations applied successfully.")
}

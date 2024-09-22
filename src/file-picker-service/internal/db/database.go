package db

import (
	"fmt"
	"log"
	"os"

	"file-picker-service/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDatabase initializes the database connection and runs migrations.
func InitDatabase() (*gorm.DB, error) {
	// Load database configuration from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Build the DSN (Data Source Name) for the PostgreSQL connection
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	// Connect to the database using GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return nil, err
	}

	// Assign the database connection to the global DB variable
	DB = db

	// Run migrations
	err = runMigrations()
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
		return nil, err
	}

	log.Println("Database connection established successfully.")
	return DB, nil
}

// runMigrations runs the GORM auto-migrations for all models.
func runMigrations() error {
	log.Println("Running database migrations...")

	// Migrate the schema for User and File models
	err := DB.AutoMigrate(&models.User{}, &models.File{})
	if err != nil {
		return fmt.Errorf("failed to run migrations: %v", err)
	}

	log.Println("Database migrations completed successfully.")
	return nil
}
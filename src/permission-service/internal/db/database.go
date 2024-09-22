package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConn *gorm.DB // Global variable for the database connection

// InitDB initializes the database connection using the provided DB URL.
func InitDB(dbUrl string) error {
	var err error

	// Attempt to open a connection to the database
	DBConn, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to the database: %v", err)
		return err
	}

	log.Println("Database connection initialized successfully")
	return nil
}

// CloseDB closes the database connection. Use this when shutting down the service gracefully.
func CloseDB() error {
	sqlDB, err := DBConn.DB()
	if err != nil {
		log.Printf("Failed to get the raw database connection for closing: %v", err)
		return err
	}

	// Close the connection
	err = sqlDB.Close()
	if err != nil {
		log.Printf("Error closing the database connection: %v", err)
		return err
	}

	log.Println("Database connection closed successfully")
	return nil
}

// Migrate runs the database migrations for the given models. This should be run when the service starts.
func Migrate(models ...interface{}) error {
	err := DBConn.AutoMigrate(models...)
	if err != nil {
		log.Printf("Failed to run migrations: %v", err)
		return err
	}

	log.Println("Database migrations completed successfully")
	return nil
}
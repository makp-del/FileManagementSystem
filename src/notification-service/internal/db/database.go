package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DBConn *sql.DB

// InitDB initializes the database connection
func InitDB() error {
	// Get the database URL from environment variables
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	// Open the PostgreSQL database connection
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %v", err)
	}

	// Verify the connection is successful
	err = conn.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	DBConn = conn
	log.Println("Database connection established successfully")
	return nil
}

// CloseDB closes the database connection
func CloseDB() {
	if DBConn != nil {
		DBConn.Close()
		log.Println("Database connection closed")
	}
}
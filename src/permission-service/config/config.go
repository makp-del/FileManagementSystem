package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	Port string
	DBUrl string
}

// LoadConfig loads environment variables from .env file (if present) or system envs
func LoadConfig() *Config {
	// Load .env file if it exists (optional)
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, loading system environment variables")
	}

	// Set default values if environment variables are not set
	port := os.Getenv("PERMISSION_SERVICE_PORT")
	if port == "" {
		port = "50053"
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		dbUrl = "postgres://admin:admin@localhost:5432/permission_service_db?sslmode=disable"
	}

	return &Config{
		Port: port,
		DBUrl: dbUrl,
	}
}
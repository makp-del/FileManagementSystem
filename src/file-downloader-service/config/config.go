package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                   string
	NotificationServiceAddr string
}

var AppConfig Config

// LoadConfig loads the configuration from the environment variables or .env file.
func LoadConfig() error {
	// Load .env file if present
	err := godotenv.Load()
	if err != nil {
		log.Printf("No .env file found, using system environment variables.")
	}

	// Set the configuration values
	AppConfig = Config{
		Port:                   getEnv("FILE_DOWNLOADER_SERVICE_PORT", "50052"),
		NotificationServiceAddr: getEnv("NOTIFICATION_SERVICE_ADDRESS", "localhost:50054"),
	}

	return nil
}

// Helper function to get environment variable values or default values.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
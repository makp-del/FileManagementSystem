package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort         string
	LogLevel           string
	LogFormat          string
	TransformationTimeout int // Timeout in seconds for file transformations
}

// LoadConfig loads configuration from environment variables or a .env file
func LoadConfig() *Config {
	// Load environment variables from .env file for local development
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	return &Config{
		ServerPort:         getEnv("SERVER_PORT", "50051"),                         // Default to port 50051 if not set
		LogLevel:           getEnv("LOG_LEVEL", "info"),                            // Default to info level logging
		LogFormat:          getEnv("LOG_FORMAT", "text"),                           // Default to text format logging
		TransformationTimeout: getEnvAsInt("TRANSFORMATION_TIMEOUT", 30),           // Default to 30 seconds for transformation timeout
	}
}

// Helper function to fetch environment variables or provide a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// Helper function to fetch environment variables as integers or provide a default value
func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		var intValue int
		_, err := fmt.Sscanf(value, "%d", &intValue)
		if err != nil {
			log.Printf("Invalid value for %s, defaulting to %d", key, defaultValue)
			return defaultValue
		}
		return intValue
	}
	return defaultValue
}
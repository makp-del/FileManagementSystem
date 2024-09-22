package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all the configuration values required by the service
type Config struct {
	ServerPort    string
	DatabaseURL   string
	GrpcAddresses GrpcConfig
}

// GrpcConfig holds the addresses for gRPC communication with other services
type GrpcConfig struct {
	FileDownloaderAddress string
	PermissionsAddress    string
	NotificationAddress   string
	TransformationsAddress string
}

// LoadConfig loads the configuration from environment variables or a .env file (for local development)
func LoadConfig() *Config {
	// Load environment variables from .env file (if exists) for local development
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Load essential environment variables
	cfg := &Config{
		ServerPort:  getEnv("SERVER_PORT", "8080"),            // Default to port 8080 if not specified
		DatabaseURL: getEnv("DATABASE_URL", ""),               // Required database connection string
		GrpcAddresses: GrpcConfig{
			FileDownloaderAddress: getEnv("GRPC_FILE_DOWNLOADER_ADDRESS", ""),
			PermissionsAddress:    getEnv("GRPC_PERMISSIONS_ADDRESS", ""),
			NotificationAddress:   getEnv("GRPC_NOTIFICATION_ADDRESS", ""),
			TransformationsAddress: getEnv("GRPC_TRANSFORMATIONS_ADDRESS", ""),
		},
	}

	// Ensure all necessary environment variables are set
	validateConfig(cfg)

	fmt.Println("Configuration loaded successfully")

	return cfg
}

// getEnv retrieves an environment variable or returns a default value if not set
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// validateConfig ensures that all required environment variables are present
func validateConfig(cfg *Config) {
	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}
	if cfg.GrpcAddresses.FileDownloaderAddress == "" {
		log.Fatal("GRPC_FILE_DOWNLOADER_ADDRESS environment variable is required")
	}
	if cfg.GrpcAddresses.PermissionsAddress == "" {
		log.Fatal("GRPC_PERMISSIONS_ADDRESS environment variable is required")
	}
	if cfg.GrpcAddresses.NotificationAddress == "" {
		log.Fatal("GRPC_NOTIFICATION_ADDRESS environment variable is required")
	}
}
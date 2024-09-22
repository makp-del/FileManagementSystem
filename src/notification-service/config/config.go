package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	Port     string // WebSocket server port
	GRPCPort string // gRPC server port
}

// LoadConfig loads environment variables from .env file (if present) or system envs
func LoadConfig() *Config {
	// Load .env file if it exists
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, loading system environment variables")
	}

	// Set default values if environment variables are not set
	port := os.Getenv("NOTIFICATION_SERVICE_PORT")
	if port == "" {
		port = "50054" // Default WebSocket port
	}

	grpcPort := os.Getenv("NOTIFICATION_SERVICE_GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50055" // Default gRPC port
	}

	return &Config{
		Port:     port,
		GRPCPort: grpcPort,
	}
}
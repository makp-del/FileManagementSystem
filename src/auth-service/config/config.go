package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config stores all the configuration needed for the application
type Config struct {
	JWTIssuer        string
	PrivateKeyPath   string
	AdminUsername    string
	AdminPassword    string
	AdminEmail       string
}

// LoadConfig loads environment variables from the .env file and validates them
func LoadConfig() *Config {
	// Load environment variables from .env file if present
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Load and validate environment variables
	config := &Config{
		JWTIssuer:      getEnv("JWT_ISSUER", "default-issuer"),
		PrivateKeyPath: getEnv("PRIVATE_KEY_FILEPATH", ""),
		AdminUsername:  getEnv("ADMIN_USERNAME", "admin"),
		AdminPassword:  getEnv("ADMIN_PASSWORD", ""),
		AdminEmail:     getEnv("ADMIN_EMAIL", "admin@example.com"),
	}

	if config.PrivateKeyPath == "" {
		log.Fatal("PRIVATE_KEY_FILEPATH is required")
	}

	return config
}

// getEnv retrieves environment variables with an optional fallback value
func getEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
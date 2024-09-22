package pkg

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger is the global logger for the application
var Logger *logrus.Logger

// InitLogger initializes the logger
func InitLogger() {
	Logger = logrus.New()

	// Set the output to stdout
	Logger.SetOutput(os.Stdout)

	// Set the log level from the environment (default: info)
	logLevel, exists := os.LookupEnv("LOG_LEVEL")
	if !exists {
		logLevel = "info"
	}

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		Logger.Warn("Invalid log level provided. Defaulting to 'info'.")
		level = logrus.InfoLevel
	}
	Logger.SetLevel(level)

	// Set formatter to JSON if needed (can be set via an env variable)
	logFormat, exists := os.LookupEnv("LOG_FORMAT")
	if exists && logFormat == "json" {
		Logger.SetFormatter(&logrus.JSONFormatter{})
	}
}
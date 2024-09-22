package pkg

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger is the global logger instance used throughout the service
var Logger *logrus.Logger

// InitLogger initializes the global logger with standard settings
func InitLogger() {
	Logger = logrus.New()

	// Set output to stdout (for containerized environments)
	Logger.SetOutput(os.Stdout)

	// Set log level from the environment, default to Info level
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		Logger.Warn("Invalid log level provided. Defaulting to 'info'.")
		level = logrus.InfoLevel
	}
	Logger.SetLevel(level)

	// Set the log format to JSON for structured logging
	logFormat := os.Getenv("LOG_FORMAT")
	if logFormat == "json" {
		Logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		Logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}
}
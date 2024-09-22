package pkg

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger is a global logger instance
var Logger *logrus.Logger

// InitLogger initializes the global logger with specified settings
func InitLogger() {
	Logger = logrus.New()

	// Set the log output to stdout
	Logger.SetOutput(os.Stdout)

	// Set the log level from environment variables (default to "info")
	level, exists := os.LookupEnv("LOG_LEVEL")
	if !exists {
		level = "info" // Default log level
	}

	// Parse and set the log level
	parsedLevel, err := logrus.ParseLevel(level)
	if err != nil {
		Logger.Warnf("Invalid LOG_LEVEL '%s', defaulting to info", level)
		parsedLevel = logrus.InfoLevel
	}
	Logger.SetLevel(parsedLevel)

	// Set the log format from environment variables (default to "text")
	format, exists := os.LookupEnv("LOG_FORMAT")
	if !exists {
		format = "text" // Default format
	}

	// Set the logger format (text or JSON)
	if format == "json" {
		Logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		Logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}

	Logger.Infof("Logger initialized with level: %s and format: %s", level, format)
}
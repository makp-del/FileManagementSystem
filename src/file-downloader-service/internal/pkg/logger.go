package pkg

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger is the global logger instance
var Logger *logrus.Logger

// InitLogger initializes the global logger
func InitLogger() {
	Logger = logrus.New()

	// Set the output to stdout
	Logger.SetOutput(os.Stdout)

	// Set the log level based on environment variable (or default to "info")
	level, exists := os.LookupEnv("LOG_LEVEL")
	if !exists {
		level = "info" // default to info level logging
	}

	// Parse and set log level
	parsedLevel, err := logrus.ParseLevel(level)
	if err != nil {
		Logger.Warnf("Invalid LOG_LEVEL '%s', defaulting to info", level)
		parsedLevel = logrus.InfoLevel
	}
	Logger.SetLevel(parsedLevel)

	// Set the log format based on environment variable (or default to text)
	format, exists := os.LookupEnv("LOG_FORMAT")
	if !exists {
		format = "text" // default to text format logging
	}

	// Set logger format (text or JSON)
	if format == "json" {
		Logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		Logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}

	Logger.Infof("Logger initialized with level: %s and format: %s", level, format)
}
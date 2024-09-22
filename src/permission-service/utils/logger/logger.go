package logger

import (
	"log"
	"os"
)

var (
	// Info logs general information
	Info *log.Logger

	// Warning logs warning information
	Warning *log.Logger

	// Error logs error information
	Error *log.Logger
)

// Init initializes the loggers with custom formatting.
func Init() {
	// Create a file to log the output (or change this to log to console)
	file, err := os.OpenFile("permission-service.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}

	Info = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
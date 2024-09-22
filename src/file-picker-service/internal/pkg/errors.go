package pkg

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIError is a custom error type for the API
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error implements the error interface for APIError
func (e *APIError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}

// NewAPIError creates a new instance of APIError
func NewAPIError(code int, message string) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
	}
}

// HandleError is a utility function to handle errors and send the appropriate response
func HandleError(c *gin.Context, err error) {
	// Check if the error is of type APIError
	if apiErr, ok := err.(*APIError); ok {
		// Log the error and send the response with the proper status code
		Logger.WithField("error", apiErr.Message).Errorf("API Error: %s", apiErr.Error())
		c.JSON(apiErr.Code, gin.H{
			"error": apiErr.Message,
		})
		return
	}

	// For all other errors, send a 500 Internal Server Error
	Logger.WithField("error", err.Error()).Error("Internal Server Error")
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "Internal Server Error",
	})
}

// NotFoundError creates a 404 Not Found error
func NotFoundError(message string) *APIError {
	return NewAPIError(http.StatusNotFound, message)
}

// BadRequestError creates a 400 Bad Request error
func BadRequestError(message string) *APIError {
	return NewAPIError(http.StatusBadRequest, message)
}

// InternalServerError creates a 500 Internal Server Error
func InternalServerError(message string) *APIError {
	return NewAPIError(http.StatusInternalServerError, message)
}
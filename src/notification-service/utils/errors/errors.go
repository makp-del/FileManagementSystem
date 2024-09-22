package errors

import "errors"

// Define custom errors for the notification service
var (
	ErrNotificationFailed = errors.New("failed to send notification")
	ErrInvalidUserID      = errors.New("invalid user ID provided")
	ErrInvalidMessage     = errors.New("invalid or empty message provided")
)
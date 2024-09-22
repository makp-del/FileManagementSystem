package errors

import "errors"

// Common error messages
var (
	ErrPermissionNotFound = errors.New("permission not found")
	ErrInvalidPermission  = errors.New("invalid permission type")
	ErrUserNotAuthorized  = errors.New("user not authorized to access this file")
	ErrDatabaseError      = errors.New("database error occurred")
)

// WrapDatabaseError provides a wrapper for logging and returning database errors
func WrapDatabaseError(err error) error {
	if err != nil {
		// Optionally log the database error here
		return ErrDatabaseError
	}
	return nil
}
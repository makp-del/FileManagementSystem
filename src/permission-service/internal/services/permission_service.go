package services

import (
	"permission-service/internal/models"
	"permission-service/utils/logger"
	"permission-service/utils/errors"
	"gorm.io/gorm"
)

type PermissionService struct {
	DB *gorm.DB
}

// NewPermissionService creates a new PermissionService
func NewPermissionService(db *gorm.DB) *PermissionService {
	return &PermissionService{DB: db}
}

// CheckPermission checks if a user has a specific permission for a file.
func (s *PermissionService) CheckPermission(userID uint64, fileID string, permissionType string) (bool, error) {
	logger.Info.Println("Checking permission for user:", userID, "on file:", fileID)

	// Check the permission in the database
	hasPermission, err := models.GetPermission(s.DB, userID, fileID, permissionType)
	if err != nil {
		logger.Error.Println("Error checking permission for user:", userID, "Error:", err)
		return false, errors.WrapDatabaseError(err)
	}

	logger.Info.Println("Permission check result for user:", userID, "on file:", fileID, ":", hasPermission)
	return hasPermission, nil
}

// UpdatePermission updates permissions for a file or folder for a shared user.
func (s *PermissionService) UpdatePermission(ownerID, sharedUserID uint64, fileIDs []string, permissions []string) error {
	logger.Info.Println("Updating permissions for owner:", ownerID, "to share with user:", sharedUserID)

	for _, fileID := range fileIDs {
		for _, perm := range permissions {
			// Create a new permission entry or update an existing one
			err := models.CreatePermission(s.DB, sharedUserID, fileID, ownerID, perm)
			if err != nil {
				logger.Error.Println("Error updating permission for file:", fileID, "and user:", sharedUserID, "Error:", err)
				return errors.WrapDatabaseError(err)
			}
		}
	}

	logger.Info.Println("Permissions successfully updated for user:", sharedUserID, "on files:", fileIDs)
	return nil
}
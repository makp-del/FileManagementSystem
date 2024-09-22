package models

import (
	"permission-service/utils/errors"
	"permission-service/utils/logger"
	"time"

	"gorm.io/gorm"
)

// Permission represents the permissions granted to a user for a file.
type Permission struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement"`
	UserID     uint64    `gorm:"not null"`
	FileID     string    `gorm:"not null"`
	Permission string    `gorm:"type:varchar(255)"`
	OwnerID    uint64    `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

// CreatePermission creates a new permission entry in the database.
func CreatePermission(db *gorm.DB, userID uint64, fileID string, ownerID uint64, permission string) error {
	newPermission := Permission{
		UserID:     userID,
		FileID:     fileID,
		OwnerID:    ownerID,
		Permission: permission,
	}

	// Create the permission in the database
	err := db.Create(&newPermission).Error
	if err != nil {
		logger.Error.Println("Failed to create permission:", err)
		return errors.WrapDatabaseError(err)
	}

	logger.Info.Println("Permission created successfully for user:", userID)
	return nil
}

// GetPermissionsByUser retrieves all permissions for a specific user.
func GetPermissionsByUser(db *gorm.DB, userID uint64) ([]Permission, error) {
	var permissions []Permission
	err := db.Where("user_id = ?", userID).Find(&permissions).Error
	if err != nil {
		logger.Error.Println("Failed to retrieve permissions for user:", userID, "Error:", err)
		return nil, errors.WrapDatabaseError(err)
	}
	return permissions, nil
}

// GetPermission checks if a specific user has a specific permission for a file.
func GetPermission(db *gorm.DB, userID uint64, fileID string, permissionType string) (bool, error) {
	var permission Permission
	err := db.Where("user_id = ? AND file_id = ? AND permission = ?", userID, fileID, permissionType).First(&permission).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Warning.Println("Permission not found for user:", userID, "on file:", fileID)
			return false, errors.ErrPermissionNotFound
		}
		logger.Error.Println("Failed to check permission for user:", userID, "Error:", err)
		return false, errors.WrapDatabaseError(err)
	}

	logger.Info.Println("Permission found for user:", userID, "on file:", fileID)
	return true, nil
}

// UpdatePermission updates the permission for a user on a specific file.
func UpdatePermission(db *gorm.DB, userID, fileID uint64, newPermission string) error {
	var permission Permission
	err := db.Where("user_id = ? AND file_id = ?", userID, fileID).First(&permission).Error
	if err != nil {
		logger.Error.Println("Failed to find permission to update for user:", userID, "on file:", fileID)
		return errors.WrapDatabaseError(err)
	}

	permission.Permission = newPermission
	if err := db.Save(&permission).Error; err != nil {
		logger.Error.Println("Failed to update permission for user:", userID, "on file:", fileID, "Error:", err)
		return errors.WrapDatabaseError(err)
	}

	logger.Info.Println("Permission updated successfully for user:", userID, "on file:", fileID)
	return nil
}

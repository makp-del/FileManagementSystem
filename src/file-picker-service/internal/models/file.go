package models

import (
	"time"

	"gorm.io/gorm"
)

// File represents the metadata of an uploaded file.
type File struct {
	ID        string      `gorm:"primaryKey"`
	FileName  string    `gorm:"not null"`
	FilePath  string    `gorm:"not null"`
	OwnerID   uint      `gorm:"not null"`  // ID of the user who owns the file
	IsShared  bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// SaveFileMetadata stores the metadata of a newly uploaded file in the database.
func SaveFileMetadata(db *gorm.DB, fileId string, fileName string, ownerID uint) error {
	file := File{
		ID: fileId,
		FileName: fileName,
		FilePath: "./uploads/" + fileName,  // Assuming this is where files are saved
		OwnerID:  ownerID,
		IsShared: false,  // Initially, the file is not shared
	}

	return db.Create(&file).Error
}

// ListFilesByUser lists all files that belong to or are shared with the user.
func ListFilesByUser(db *gorm.DB, userID uint) ([]File, error) {
	var files []File
	err := db.Where("owner_id = ? OR is_shared = true", userID).Find(&files).Error
	return files, err
}

func GetFilePath(db *gorm.DB, fileID string) (string, error) {
	var file File
	err := db.First(&file, fileID).Error
	return file.FilePath, err
}
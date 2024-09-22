package models

import (
	"time"

	"gorm.io/gorm"
)

// File represents metadata for a downloaded file
type File struct {
	// Primary Key
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	FileName  string         `gorm:"not null" json:"file_name"`
	FilePath  string         `gorm:"not null" json:"file_path"`
	FileSize  int64          `gorm:"not null" json:"file_size"`
	Status    string         `gorm:"not null" json:"status"` // e.g., downloading, completed, failed
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`         // Soft delete support
}

// TableName sets the name of the table in the database
func (File) TableName() string {
	return "files"
}
package models

import "time"
import "gorm.io/gorm"

// User represents the user in the system.
type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"uniqueIndex;not null"`
	Email     string    `gorm:"uniqueIndex;not null"`
	Password  string    `gorm:"not null"`
	Role      string    `gorm:"not null"`  // User's role (e.g., admin, viewer)
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// GetUserByUsername retrieves a user by their username.
func GetUserByUsername(db *gorm.DB, username string, user *User) error {
	return db.Where("username = ?", username).First(user).Error
}

// CreateUser adds a new user to the database.
func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}
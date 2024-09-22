package models

import (
	"errors"

	"gorm.io/gorm"
)

// User model
type User struct {
    ID        uint   `gorm:"primaryKey"`
    Username  string `gorm:"unique;not null"`
    Password  string `gorm:"not null"`
    Email     string `gorm:"unique;not null"`
    Role      string `gorm:"not null"` // Role field
}

// GetUserByUsername fetches a user by username from the database.
func GetUserByUsername(db *gorm.DB, username string, user *User) error {
	result := db.Where("username = ?", username).First(user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("user not found")
	}
	return result.Error
}

// CreateUser creates a new user in the database.
func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}
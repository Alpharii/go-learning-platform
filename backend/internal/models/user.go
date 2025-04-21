package models

import "gorm.io/gorm"

// User represents the user table
type User struct {
    gorm.Model
    GoogleID string `gorm:"unique"` // Unique identifier from Google OAuth
    Email    string `gorm:"unique"` // Email address of the user
    Profile  Profile // One-to-one relationship with Profile
}

// Profile represents the profile table
type Profile struct {
    gorm.Model
    UserID uint   `gorm:"unique"` // Foreign key to User
    Name   string // Full name of the user
    Image  string // URL of the profile image
}

// Migrate runs database migrations for User and Profile
func Migrate(db *gorm.DB) error {
    return db.AutoMigrate(&User{}, &Profile{})
}
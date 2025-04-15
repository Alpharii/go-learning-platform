package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    GoogleID string `gorm:"unique"`
    Name     string
    Email    string `gorm:"unique"`
}

func Migrate(db *gorm.DB) error {
    return db.AutoMigrate(&User{})
}
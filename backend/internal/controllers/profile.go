package controllers

import (
    "net/http"

    "go-learn-platform/internal/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// GetProfile retrieves the profile of a user by their ID
func GetProfile(c *gin.Context, db *gorm.DB) {
    userID := c.Param("id")

    var user models.User
    result := db.Preload("Profile").First(&user, userID) // Load user with associated profile
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "id":    user.ID,
        "email": user.Email,
        "profile": gin.H{
            "name":  user.Profile.Name,
            "image": user.Profile.Image,
        },
    })
}

// UpdateProfile updates the profile of a user
func UpdateProfile(c *gin.Context, db *gorm.DB) {
    userID := c.Param("id")
    var input struct {
        Name  string `json:"name"`
        Image string `json:"image"` // URL of the new profile image
    }

    // Validate input
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // Find the user and their profile
    var user models.User
    result := db.Preload("Profile").First(&user, userID)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    // Update or create the profile
    if user.Profile.UserID == 0 { // If profile doesn't exist, create it
        user.Profile = models.Profile{
            UserID: user.ID,
            Name:   input.Name,
            Image:  input.Image,
        }
    } else { // If profile exists, update it
        user.Profile.Name = input.Name
        user.Profile.Image = input.Image
    }

    db.Save(&user.Profile) // Save the updated profile

    c.JSON(http.StatusOK, gin.H{
        "message": "Profile updated successfully",
        "user": gin.H{
            "id":    user.ID,
            "email": user.Email,
            "profile": gin.H{
                "name":  user.Profile.Name,
                "image": user.Profile.Image,
            },
        },
    })
}
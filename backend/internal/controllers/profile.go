package controllers

import (
	"fmt"
	"net/http"

	"go-learn-platform/internal/middleware"
	"go-learn-platform/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetMyProfile retrieves the profile of the currently authenticated user
func GetMyProfile(c *gin.Context, db *gorm.DB) {
    userID := c.MustGet("userID").(uint)

    var user models.User
    result := db.Preload("Profile").First(&user, userID)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    // Construct full image URL
    imageURL := ""
    if user.Profile.Image != "" {
        imageURL = fmt.Sprintf("http://localhost:8080%s", user.Profile.Image)
    }

    c.JSON(http.StatusOK, gin.H{
        "id":    user.ID,
        "email": user.Email,
        "profile": gin.H{
            "name":  user.Profile.Name,
            "image": imageURL,
        },
    })
}

// GetProfile retrieves the profile of a user by their ID
func GetProfile(c *gin.Context, db *gorm.DB) {
    userID := c.Param("id")

    var user models.User
    result := db.Preload("Profile").First(&user, userID) // Load user with associated profile
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    // Construct full image URL
    imageURL := ""
    if user.Profile.Image != "" {
        imageURL = fmt.Sprintf("http://localhost:8080%s", user.Profile.Image)
    }

    c.JSON(http.StatusOK, gin.H{
        "id":    user.ID,
        "email": user.Email,
        "profile": gin.H{
            "name":  user.Profile.Name,
            "image": imageURL,
        },
    })
}

func UpdateProfile(c *gin.Context, db *gorm.DB) {
    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
        return
    }


    // Ambil field name dari form
    name := c.PostForm("name")

    // Upload file image
    imageURL, err := middleware.UploadFile(c, "image")
    if err != nil && err.Error() != "http: no such file" {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Temukan user dan preload Profile
    var user models.User
    result := db.Preload("Profile").First(&user, userID)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    // Jika profil belum ada, buat baru
    if user.Profile.UserID == 0 {
        newProfile := models.Profile{
            UserID: user.ID,
            Name:   name,
            Image:  imageURL,
        }
        if err := db.Create(&newProfile).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create profile"})
            return
        }
        user.Profile = newProfile
    } else {
        // Update profil yang ada
        user.Profile.Name = name
        if imageURL != "" {
            user.Profile.Image = imageURL
        }
        if err := db.Save(&user.Profile).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
            return
        }
    }

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
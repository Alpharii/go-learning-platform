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
    result := db.
        Preload("Profile").
        Preload("Courses").
        Preload("Enrollments.Course").
        First(&user, userID)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    imageURL := ""
    if user.Profile.Image != "" {
        imageURL = fmt.Sprintf("http://localhost:8080%s", user.Profile.Image)
    }

    createdCourses := make([]gin.H, 0)
    for _, course := range user.Courses {
        courseImage := ""
        if course.Image != "" {
            courseImage = fmt.Sprintf("http://localhost:8080%s", course.Image)
        }

        createdCourses = append(createdCourses, gin.H{
            "id":          course.ID,
            "title":       course.Title,
            "description": course.Description,
            "image":       courseImage,
        })
    }

    enrolledCourses := make([]gin.H, 0)
    for _, enrollment := range user.Enrollments {
        courseImage := ""
        if enrollment.Course.Image != "" {
            courseImage = fmt.Sprintf("http://localhost:8080%s", enrollment.Course.Image)
        }

        enrolledCourses = append(enrolledCourses, gin.H{
            "id":          enrollment.Course.ID,
            "title":       enrollment.Course.Title,
            "description": enrollment.Course.Description,
            "image":       courseImage,
            "progress":    enrollment.Progress,
        })
    }

    c.JSON(http.StatusOK, gin.H{
        "id":    user.ID,
        "email": user.Email,
        "profile": gin.H{
            "name":  user.Profile.Name,
            "image": imageURL,
        },
        "created_courses":  createdCourses,
        "enrolled_courses": enrolledCourses,
    })
}


// GetProfile retrieves the profile of a user by their ID
func GetProfile(c *gin.Context, db *gorm.DB) {
    userID := c.Param("id")

    var user models.User
    result := db.
        Preload("Profile").
        Preload("Courses").
        Preload("Enrollments.Course").
        First(&user, userID)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    imageURL := ""
    if user.Profile.Image != "" {
        imageURL = fmt.Sprintf("http://localhost:8080%s", user.Profile.Image)
    }

    createdCourses := make([]gin.H, 0)
    for _, course := range user.Courses {
        createdCourses = append(createdCourses, gin.H{
            "id":          course.ID,
            "title":       course.Title,
            "description": course.Description,
            "image":       course.Image,
        })
    }

    enrolledCourses := make([]gin.H, 0)
    for _, enrollment := range user.Enrollments {
        enrolledCourses = append(enrolledCourses, gin.H{
            "id":          enrollment.Course.ID,
            "title":       enrollment.Course.Title,
            "description": enrollment.Course.Description,
            "image":       enrollment.Course.Image,
            "progress":    enrollment.Progress,
        })
    }

    c.JSON(http.StatusOK, gin.H{
        "id":    user.ID,
        "email": user.Email,
        "profile": gin.H{
            "name":  user.Profile.Name,
            "image": imageURL,
        },
        "created_courses":  createdCourses,
        "enrolled_courses": enrolledCourses,
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
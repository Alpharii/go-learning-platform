package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"go-learn-platform/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// EnrollUser enrolls a user to a course
func EnrollUser(c *gin.Context, db *gorm.DB) {
    var input struct {
        CourseID uint `json:"course_id" binding:"required"`
    }

    // Validasi input
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // Ambil user_id dari context (diatur oleh middleware)
    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
        return
    }

    // Pastikan userID adalah uint
    userIDUint, ok := userID.(uint)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user ID"})
        return
    }

    // Cek apakah kursus ada
    var course models.Course
    if err := db.First(&course, input.CourseID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
        return
    }

    // Cek apakah pengguna sudah terdaftar di kursus
    var existingEnrollment models.Enrollment
    if err := db.Where("user_id = ? AND course_id = ?", userIDUint, input.CourseID).First(&existingEnrollment).Error; err == nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User is already enrolled in this course"})
        return
    }

    // Buat pendaftaran baru
    enrollment := models.Enrollment{
        UserID:   userIDUint,
        CourseID: input.CourseID,
    }
    if err := db.Create(&enrollment).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enroll user"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "User enrolled successfully",
        "enrollment": enrollment,
    })
}

// GetEnrollments retrieves all courses a user is enrolled in
func GetEnrollments(c *gin.Context, db *gorm.DB) {
    userID, err := strconv.Atoi(c.Param("user_id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    var enrollments []models.Enrollment
    if err := db.Preload("Course").Where("user_id = ?", userID).Find(&enrollments).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch enrollments"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"enrollments": enrollments})
}

// CancelEnrollment cancels a user's enrollment in a course
func CancelEnrollment(c *gin.Context, db *gorm.DB) {
    enrollmentID, err := strconv.Atoi(c.Param("id"))
    fmt.Println(enrollmentID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid enrollment ID"})
        return
    }

    var enrollment models.Enrollment
    if err := db.First(&enrollment, enrollmentID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Enrollment not found"})
        return
    }

    // Ambil user_id dari context (diatur oleh middleware)
    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
        return
    }

    // Pastikan userID adalah uint
    userIDUint, ok := userID.(uint)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user ID"})
        return
    }

    // Verifikasi bahwa pengguna adalah pemilik pendaftaran
    if enrollment.UserID != userIDUint {
        c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to cancel this enrollment"})
        return
    }

    // Hapus pendaftaran
    if err := db.Delete(&enrollment).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel enrollment"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Enrollment canceled successfully"})
}
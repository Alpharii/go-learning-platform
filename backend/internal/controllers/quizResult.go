package controllers

import (
    "net/http"
    "strconv"

    "go-learn-platform/internal/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// GetQuizResults retrieves all quiz results from the database
func GetQuizResults(c *gin.Context, db *gorm.DB) {
    var results []models.QuizResult
    if err := db.Preload("Quiz").Find(&results).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch quiz results"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"quiz_results": results})
}

// CreateQuizResult creates a new quiz result in the database
func CreateQuizResult(c *gin.Context, db *gorm.DB) {
    var input struct {
        QuizID uint `json:"quiz_id" binding:"required"`
        Score  int  `json:"score" binding:"required"`
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

    // Cek apakah quiz ada
    var quiz models.Quiz
    if err := db.First(&quiz, input.QuizID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Quiz not found"})
        return
    }

    // Buat hasil quiz baru
    result := models.QuizResult{
        UserID: userIDUint,
        QuizID: input.QuizID,
        Score:  input.Score,
    }
    if err := db.Create(&result).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create quiz result"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Quiz result created successfully",
        "result":  result,
    })
}

// DeleteQuizResult deletes a quiz result by ID
func DeleteQuizResult(c *gin.Context, db *gorm.DB) {
    resultID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quiz result ID"})
        return
    }

    var result models.QuizResult
    if err := db.First(&result, resultID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Quiz result not found"})
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

    // Verifikasi bahwa pengguna adalah pemilik hasil quiz
    if result.UserID != userIDUint {
        c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to delete this quiz result"})
        return
    }

    // Hapus hasil quiz
    if err := db.Delete(&result).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete quiz result"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Quiz result deleted successfully"})
}
package controllers

import (
    "net/http"
    "strconv"

    "go-learn-platform/internal/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)


// GetAllQuizzes retrieves all quizzes from the database
func GetAllQuizzes(c *gin.Context, db *gorm.DB) {
    var quizzes []models.Quiz
    if err := db.Preload("Lesson").Find(&quizzes).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch quizzes"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"quizzes": quizzes})
}

// CreateQuiz creates a new quiz in the database
func CreateQuiz(c *gin.Context, db *gorm.DB) {
    var input struct {
        LessonID uint   `json:"lesson_id" binding:"required"`
        Question string `json:"question" binding:"required"`
        Options  string `json:"options" binding:"required"`
        Answer   string `json:"answer" binding:"required"`
    }

    // Validasi input
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // Cek apakah lesson ada
    var lesson models.Lesson
    if err := db.First(&lesson, input.LessonID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Lesson not found"})
        return
    }

    // Buat quiz baru
    quiz := models.Quiz{
        LessonID: input.LessonID,
        Question: input.Question,
        Options:  input.Options,
        Answer:   input.Answer,
    }
    if err := db.Create(&quiz).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create quiz"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Quiz created successfully",
        "quiz":    quiz,
    })
}

// DeleteQuiz deletes a quiz by ID
func DeleteQuiz(c *gin.Context, db *gorm.DB) {
    quizID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quiz ID"})
        return
    }

    var quiz models.Quiz
    if err := db.First(&quiz, quizID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Quiz not found"})
        return
    }

    // Hapus quiz
    if err := db.Delete(&quiz).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete quiz"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Quiz deleted successfully"})
}

// CompleteQuiz marks a quiz as completed and updates the course progress
func CompleteQuiz(c *gin.Context, db *gorm.DB) {
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

    // Ambil quiz_id dari parameter URL
    quizID, err := strconv.Atoi(c.Param("quiz_id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quiz ID"})
        return
    }

    // Cek apakah quiz ada dan muat data Lesson terkait
    var quiz models.Quiz
    if err := db.Preload("Lesson").First(&quiz, quizID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Quiz not found"})
        return
    }

    // Cek apakah pengguna sudah menyelesaikan quiz ini
    var existingResult models.QuizResult
    if err := db.Where("user_id = ? AND quiz_id = ?", userIDUint, quizID).First(&existingResult).Error; err == nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "You have already completed this quiz"})
        return
    }

    // Simpan hasil quiz baru
    var input struct {
        Score int `json:"score" binding:"required"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    result := models.QuizResult{
        UserID: userIDUint,
        QuizID: uint(quizID),
        Score:  input.Score,
    }
    if err := db.Create(&result).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save quiz result"})
        return
    }

    // Perbarui progress kursus
    if err := UpdateCourseProgress(db, userIDUint, quiz.Lesson.CourseID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update course progress"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Quiz completed successfully", "result": result})
}
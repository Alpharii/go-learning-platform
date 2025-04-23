package controllers

import (
    "net/http"
    "strconv"

    "go-learn-platform/internal/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// GetCourseProgress calculates and returns the progress of a user in a course
func GetCourseProgress(c *gin.Context, db *gorm.DB) {
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

    // Ambil course_id dari parameter URL
    courseID, err := strconv.Atoi(c.Param("course_id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
        return
    }

    // Cek apakah pengguna terdaftar di kursus
    var enrollment models.Enrollment
    if err := db.Where("user_id = ? AND course_id = ?", userIDUint, courseID).First(&enrollment).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User is not enrolled in this course"})
        return
    }

    // Hitung total lesson dalam kursus
    var totalLessons int64
    if err := db.Model(&models.Lesson{}).Where("course_id = ?", courseID).Count(&totalLessons).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count lessons"})
        return
    }

    // Hitung jumlah lesson yang telah diselesaikan oleh pengguna
    var completedLessons int64
    if err := db.Table("lesson_progress").
        Where("user_id = ? AND lesson_id IN (?)", userIDUint, db.Table("lessons").Select("id").Where("course_id = ?", courseID)).
        Count(&completedLessons).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count completed lessons"})
        return
    }

    // Hitung progress dalam persentase
    progress := 0.0
    if totalLessons > 0 {
        progress = (float64(completedLessons) / float64(totalLessons)) * 100
    }

    c.JSON(http.StatusOK, gin.H{
        "course_id": courseID,
        "progress":  progress,
    })
}

// CreateCourse creates a new course
func CreateCourse(c *gin.Context, db *gorm.DB) {
    var input struct {
        Title       string `json:"title" binding:"required"`
        Description string `json:"description" binding:"required"`
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

    // Buat kursus baru
    course := models.Course{
        Title:       input.Title,
        Description: input.Description,
        UserID:      userIDUint,
    }
    if err := db.Create(&course).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create course"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Course created successfully",
        "course":  course,
    })
}

// GetCourses retrieves all courses
func GetCourses(c *gin.Context, db *gorm.DB) {
    var courses []models.Course
    if err := db.Preload("Lessons").Find(&courses).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch courses"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"courses": courses})
}

// GetCourse retrieves a specific course by ID
func GetCourse(c *gin.Context, db *gorm.DB) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
        return
    }

    var course models.Course
    if err := db.Preload("Lessons").First(&course, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"course": course})
}

// UpdateCourse updates a specific course by ID
func UpdateCourse(c *gin.Context, db *gorm.DB) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
        return
    }

    var input struct {
        Title       string `json:"title"`
        Description string `json:"description"`
    }

    // Validasi input
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    var course models.Course
    if err := db.First(&course, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
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

    // Verifikasi bahwa pengguna memiliki kursus ini
    if course.UserID != userIDUint {
        c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to update this course"})
        return
    }

    // Perbarui data kursus
    course.Title = input.Title
    course.Description = input.Description
    if err := db.Save(&course).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update course"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Course updated successfully",
        "course":  course,
    })
}

// DeleteCourse deletes a specific course by ID
func DeleteCourse(c *gin.Context, db *gorm.DB) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
        return
    }

    var course models.Course
    if err := db.First(&course, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
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

    // Verifikasi bahwa pengguna memiliki kursus ini
    if course.UserID != userIDUint {
        c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to delete this course"})
        return
    }

    if err := db.Delete(&course).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete course"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}

// UpdateCourseProgress calculates and updates the progress of a user in a course
func UpdateCourseProgress(db *gorm.DB, userID uint, courseID uint) error {
    // Hitung total lesson dalam kursus
    var totalLessons int64
    if err := db.Model(&models.Lesson{}).Where("course_id = ?", courseID).Count(&totalLessons).Error; err != nil {
        return err
    }

    // Hitung jumlah lesson yang telah diselesaikan oleh pengguna
    var completedLessons int64
    if err := db.Table("quiz_results").
        Where("user_id = ? AND quiz_id IN (?)", userID, db.Table("quizzes").Select("id").Where("lesson_id IN (?)", db.Table("lessons").Select("id").Where("course_id = ?", courseID))).
        Distinct("quiz_id").
        Count(&completedLessons).Error; err != nil {
        return err
    }

    // Hitung progress dalam persentase
    progress := 0.0
    if totalLessons > 0 {
        progress = (float64(completedLessons) / float64(totalLessons)) * 100
    }

    // Perbarui progress di tabel Enrollment
    if err := db.Model(&models.Enrollment{}).Where("user_id = ? AND course_id = ?", userID, courseID).Update("progress", progress).Error; err != nil {
        return err
    }

    return nil
}
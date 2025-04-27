package controllers

import (
	"go-learn-platform/internal/middleware"
	"go-learn-platform/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateLesson handles creating a new lesson
func CreateLesson(c *gin.Context, db *gorm.DB) {
    // Ambil data dari form
    title := c.PostForm("title")
    content := c.PostForm("content")
    orderStr := c.PostForm("order")
    courseIDStr := c.PostForm("course_id")

    order, err := strconv.Atoi(orderStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order value"})
        return
    }

    courseID, err := strconv.ParseUint(courseIDStr, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
        return
    }

    // Upload image
    imageURL, err := middleware.UploadFile(c, "image")
    if err != nil && err.Error() != "failed to retrieve file: http: no such file" {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    lesson := models.Lesson{
        Title:    title,
        Content:  content,
        Order:    order,
        CourseID: uint(courseID),
        Image:    imageURL,
    }

    if err := db.Create(&lesson).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create lesson"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Lesson created successfully", "data": lesson})
}

// UpdateLesson updates an existing lesson
func UpdateLesson(c *gin.Context, db *gorm.DB) {
    lessonID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lesson ID"})
        return
    }

    var lesson models.Lesson
    if err := db.First(&lesson, lessonID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Lesson not found"})
        return
    }

    // Ambil data dari form
    title := c.PostForm("title")
    content := c.PostForm("content")
    orderStr := c.PostForm("order")
    courseIDStr := c.PostForm("course_id")

    order, err := strconv.Atoi(orderStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order value"})
        return
    }

    courseID, err := strconv.ParseUint(courseIDStr, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
        return
    }

    // Upload file baru jika ada
    imageURL, err := middleware.UploadFile(c, "image")
    if err != nil && err.Error() != "failed to retrieve file: http: no such file" {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Update field-fieldnya
    lesson.Title = title
    lesson.Content = content
    lesson.Order = order
    lesson.CourseID = uint(courseID)
    if imageURL != "" {
        lesson.Image = imageURL // kalau ada file baru, update image
    }

    if err := db.Save(&lesson).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update lesson"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Lesson updated successfully", "data": lesson})
}


func GetLesson(c *gin.Context, db *gorm.DB){
	lessonID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var lesson models.Lesson
    if err := db.Preload("Quizzes").First(&lesson, lessonID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Lesson not found"})
        return
    }

	c.JSON(http.StatusOK, gin.H{"data": lesson})
}

// DeleteLesson deletes a lesson by ID
func DeleteLesson(c *gin.Context, db *gorm.DB) {
    lessonID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lesson ID"})
        return
    }

    var lesson models.Lesson
    if err := db.First(&lesson, lessonID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Lesson not found"})
        return
    }

    if err := db.Delete(&lesson).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete lesson"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Lesson deleted successfully"})
}
package controllers

import (
	"go-learn-platform/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateLesson(c *gin.Context, db *gorm.DB){
	var input models.Lesson
	if err:= c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
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

    var input models.Lesson
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Update fields
    lesson.Title = input.Title
    lesson.Content = input.Content
    lesson.Order = input.Order

    if err := db.Save(&lesson).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update lesson"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Lesson updated successfully", "data": lesson})
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
package controllers

import (
	"go-learn-platform/internal/models"
	"net/http"

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

func GetLesson(c *gin.Context)
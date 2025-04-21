package routes

import (
	"go-learn-platform/internal/auth"
	"go-learn-platform/internal/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(r *gin.Engine, DB *gorm.DB) {
    // Root route
    r.GET("/", func(ctx *gin.Context) {
        ctx.JSON(http.StatusOK, gin.H{
            "message": "Welcome to Go Learn Platform!",
        })
    })

    // Google OAuth routes
    r.GET("/auth/google/login", auth.HandleGoogleLogin)
    r.GET("/auth/google/callback", func(c *gin.Context) {
        auth.HandleGoogleCallback(c, DB)
    })

    // Profile routes
    r.GET("/profile/:id", func(c *gin.Context) {
        controllers.GetProfile(c, DB)
    })
    r.PUT("/profile/:id", func(c *gin.Context) {
        controllers.UpdateProfile(c, DB)
    })

}
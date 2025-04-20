package routes

import (
    "go-learn-platform/internal/auth"
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

	r.PUT("/users/:id/name", func(c *gin.Context) {
		auth.UpdateUserName(c, DB)
	})
}
package routes

import (
	"go-learn-platform/internal/auth"
	"go-learn-platform/internal/controllers"
	"go-learn-platform/internal/middleware"
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


    protected := r.Group("/")

    protected.Use(middleware.AuthMiddleware())
    {
        protected.POST("/courses", func(c *gin.Context) {
            controllers.CreateCourse(c, DB)
        })
        protected.GET("/courses", func(c *gin.Context) {
            controllers.GetCourses(c, DB)
        })
        protected.GET("/courses/:id", func(c *gin.Context) {
            controllers.GetCourse(c, DB)
        })
        protected.PUT("/courses/:id", func(c *gin.Context) {
            controllers.UpdateCourse(c, DB)
        })
        protected.DELETE("/courses/:id", func(c *gin.Context) {
            controllers.DeleteCourse(c, DB)
        })
    }

}
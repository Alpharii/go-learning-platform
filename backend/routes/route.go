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
        // Course routes
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

        // Enrollment routes
        protected.POST("/enroll", func(c *gin.Context) {
            controllers.EnrollUser(c, DB)
        })
        protected.GET("/enrollments/:user_id", func(c *gin.Context) {
            controllers.GetEnrollments(c, DB)
        })
        protected.DELETE("/enroll/:id", func(c *gin.Context) {
            controllers.CancelEnrollment(c, DB)
        })

        // Lesson routes
        protected.POST("/lessons", func(c *gin.Context) {
            controllers.CreateLesson(c, DB)
        })
        protected.GET("/lessons/:course_id", func(c *gin.Context) {
            controllers.GetLesson(c, DB)
        })
        protected.GET("/lesson/:id", func(c *gin.Context) {
            controllers.GetLesson(c, DB)
        })
        protected.PUT("/lesson/:id", func(c *gin.Context) {
            controllers.UpdateLesson(c, DB)
        })
        protected.DELETE("/lesson/:id", func(c *gin.Context) {
            controllers.DeleteLesson(c, DB)
        })
    }
}
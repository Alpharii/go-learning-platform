package models

import "gorm.io/gorm"

// User represents the user table
type User struct {
    gorm.Model
    GoogleID string `gorm:"unique"` // Unique identifier from Google OAuth
    Email    string `gorm:"unique"` // Email address of the user
    Profile  Profile               // One-to-one relationship with Profile
    Courses  []Course              // Relasi one-to-many dengan Course (sebagai instruktur)
}

// Profile represents the profile table
type Profile struct {
    gorm.Model
    UserID uint   `gorm:"unique"` // Foreign key to User
    Name   string // Full name of the user
    Image  string // URL of the profile image
}

// Course represents the course table
type Course struct {
    gorm.Model
    Title       string   `gorm:"not null"` // Judul kursus
    Description string   `gorm:"not null"` // Deskripsi kursus
    UserID      uint     `gorm:"not null"` // ID pengguna (pembuat kursus)
    Lessons     []Lesson `gorm:"foreignKey:CourseID"` // Relasi one-to-many dengan Lesson
    Enrollments []Enrollment `gorm:"foreignKey:CourseID"` // Relasi one-to-many dengan Enrollment
}

// Lesson represents the lesson table
type Lesson struct {
    gorm.Model
    CourseID uint   `gorm:"not null"`
    Title    string `gorm:"not null"`
    Content  string `gorm:"not null"`
    Order    int    `gorm:"not null"`
    Quizzes  []Quiz `gorm:"foreignKey:LessonID"`
}

// Enrollment represents the enrollment table
type Enrollment struct {
    gorm.Model
    UserID   uint    `gorm:"not null"`
    CourseID uint    `gorm:"not null"`
    Progress float64 `gorm:"default:0"` // Progress in percentage (0-100)
}

type Quiz struct {
    gorm.Model
    LessonID uint   `gorm:"not null"`
    Lesson   Lesson `gorm:"foreignKey:LessonID"` // Relasi ke Lesson
    Question string `gorm:"not null"`
    Options  string `gorm:"not null"`
    Answer   string `gorm:"not null"`
    Results  []QuizResult `gorm:"foreignKey:QuizID"`
}

// QuizResult represents the quiz result table
type QuizResult struct {
    gorm.Model
    UserID uint `gorm:"not null"`
    QuizID uint `gorm:"not null"`
    Score  int  `gorm:"not null"`
}

// Migrate runs database migrations for all models
func Migrate(db *gorm.DB) error {
    return db.AutoMigrate(
        &User{},
        &Profile{},
        &Course{},
        &Lesson{},
        &Enrollment{},
        &Quiz{},
        &QuizResult{},
    )
}
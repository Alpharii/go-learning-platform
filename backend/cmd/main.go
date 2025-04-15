package main

import (
	"fmt"
	"go-learn-platform/internal/models"
	"go-learn-platform/pkg/config"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// initDB: Inisialisasi koneksi database
func initDB() (*gorm.DB, error) {
    dsn := config.LoadConfig().DBDsn
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    return db, nil
}

// migrateDB: Migrasi model ke database
func migrateDB(db *gorm.DB) error {
    err := models.Migrate(db)
    if err != nil {
        return err
    }
    return nil
}

func main() {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    DB, err = initDB()
    if err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }
    log.Println("Successfully connected to the database!")

    // Migrate database
    // err = migrateDB(DB)
    // if err != nil {
    //     log.Fatalf("Failed to migrate database: %v", err)
    // }
    // log.Println("Database migration completed successfully!")

    r := gin.Default()

    // Root route
    r.GET("/", func(ctx *gin.Context) {
        ctx.JSON(200, gin.H{
            "message": "Welcome to Go Learn Platform!",
        })
    })

    r.Run(":8080")
	fmt.Println("server running in http://localhost:8080")
}
package main

import (
	"fmt"
	"go-learn-platform/internal/auth"
	"go-learn-platform/internal/models"
	"go-learn-platform/internal/pkg/config"
	"go-learn-platform/internal/routes"

	"log"
    "time"

    "github.com/gin-contrib/cors"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg := config.LoadConfig()
	auth.InitGoogleConfig(cfg)

    DB, err = initDB()
    if err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }
    log.Println("Successfully connected to the database!")

    err = migrateDB(DB)
    if err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }
    log.Println("Database migration completed successfully!")

    r := gin.Default()

    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"}, // frontend origin kamu
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))
    
    routes.Routes(r, DB)

	fmt.Println("server running in http://localhost:8080")
    r.Run(":8080")
}
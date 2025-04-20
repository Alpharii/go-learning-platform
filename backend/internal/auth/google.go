package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"go-learn-platform/internal/models"
	"go-learn-platform/pkg/config"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
)

var googleOauthConfig *oauth2.Config

func InitGoogleConfig(cfg *config.Config) {
	googleOauthConfig = &oauth2.Config{
		ClientID:     cfg.GoogleClientID,
		ClientSecret: cfg.GoogleClientSecret,
		RedirectURL:  cfg.RedirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}

func HandleGoogleLogin(c *gin.Context) {
    url := googleOauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
    c.Redirect(http.StatusTemporaryRedirect, url)
}

func HandleGoogleCallback(c *gin.Context, db *gorm.DB) {
    code := c.Query("code")
    token, err := googleOauthConfig.Exchange(context.Background(), code)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to exchange token"})
        return
    }

    client := googleOauthConfig.Client(context.Background(), token)

    // Ambil user info dari Google API
    resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
        return
    }
    defer resp.Body.Close()

    // Decode JSON response
    var userInfo struct {
        ID    string `json:"id"`
        Email string `json:"email"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user info"})
        return
    }

    // Cari atau buat pengguna baru di database
    var user models.User
    result := db.Where("google_id = ?", userInfo.ID).First(&user)
    if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
        newUser := models.User{
            GoogleID: userInfo.ID,
            Email:    userInfo.Email,
        }
        db.Create(&newUser)
        user = newUser
    }

    // Kirim data pengguna ke frontend (tanpa nama jika belum diatur)
    c.JSON(http.StatusOK, gin.H{
        "id":    user.ID,
        "email": user.Email,
    })
}

func UpdateUserName(c *gin.Context, db *gorm.DB) {
    userID := c.Param("id") 
    var input struct {
        Name string `json:"name" binding:"required"`
    }

    // Validasi input
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // Cari pengguna berdasarkan ID
    var user models.User
    result := db.First(&user, userID)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    // Perbarui nama pengguna
    user.Name = input.Name
    db.Save(&user)

    // Kirim respons sukses
    c.JSON(http.StatusOK, gin.H{
        "message": "Name updated successfully",
        "user": gin.H{
            "id":    user.ID,
            "name":  user.Name,
            "email": user.Email,
        },
    })
}
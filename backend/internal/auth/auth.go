package auth

import (
    "time"

    "github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("your_secret_key") // Ganti dengan secret key yang aman

// GenerateJWT generates a JWT token for the user
func GenerateJWT(userID uint, email string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "email":   email,
        "exp":     time.Now().Add(time.Hour * 24).Unix(), // Token valid selama 24 jam
    })

    return token.SignedString(jwtKey)
}
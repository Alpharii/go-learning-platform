package middleware

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("your_secret_key") // Harus sama dengan secret key di auth.go

// AuthMiddleware verifies JWT and sets user_id in context
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
            c.Abort()
            return
        }

        // Hapus "Bearer " dari token jika ada
        tokenString = tokenString[len("Bearer "):]

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        // Ambil claims dari token
        claims := token.Claims.(jwt.MapClaims)
        userID := uint(claims["user_id"].(float64))
        c.Set("userID", userID) // Simpan user_id di context

        c.Next()
    }
}
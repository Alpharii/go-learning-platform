package middleware

import (
    "crypto/rand"
    "encoding/hex"
    "fmt"
    "os"
    "path/filepath"

    "github.com/gin-gonic/gin"
)

// UploadFile handles file uploads and stores them in the public folder
func UploadFile(c *gin.Context, field string) (string, error) {
    fmt.Println("In upload middleware")

    // Get the uploaded file
    file, err := c.FormFile(field)
    if err != nil {
        return "", fmt.Errorf("failed to retrieve file: %w", err)
    }

    // Create the public directory if it doesn't exist
    uploadDir := "./public"
    if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
        if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
            return "", fmt.Errorf("failed to create upload directory: %w", err)
        }
    }

    fmt.Println("Generate unique filename")
    // Generate a unique filename
    ext := filepath.Ext(file.Filename)
    uniqueName := fmt.Sprintf("%s%s", generateUniqueID(), ext)
    filePath := filepath.Join(uploadDir, uniqueName)

    fmt.Println("Save uploaded file")
    // Save the file
    if err := c.SaveUploadedFile(file, filePath); err != nil {
        return "", fmt.Errorf("failed to save file: %w", err)
    }

    fmt.Println("Done upload middleware")

    // Return the relative URL for the file
    return fmt.Sprintf("/public/%s", uniqueName), nil
}

// generateUniqueID generates a unique identifier using random bytes
func generateUniqueID() string {
    b := make([]byte, 16) // 16 bytes = 128 bits
    _, err := rand.Read(b)
    if err != nil {
        panic(err) // Handle error appropriately in production
    }
    return hex.EncodeToString(b) // Convert to hexadecimal string
}
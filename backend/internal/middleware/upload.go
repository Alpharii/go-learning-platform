package middleware

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"

    "github.com/gin-gonic/gin"
)

// UploadFile handles file uploads and stores them in the public folder
func UploadFile(c *gin.Context, field string) (string, error) {
    file, err := c.FormFile(field)
    if err != nil {
        return "", fmt.Errorf("failed to retrieve file: %w", err)
    }

    // Create the public directory if it doesn't exist
    uploadDir := "./public"
    if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
        os.MkdirAll(uploadDir, os.ModePerm)
    }

    // Generate a unique filename
    ext := filepath.Ext(file.Filename)
    uniqueName := fmt.Sprintf("%s%s", generateUniqueID(), ext)
    filePath := filepath.Join(uploadDir, uniqueName)

    // Save the file
    if err := c.SaveUploadedFile(file, filePath); err != nil {
        return "", fmt.Errorf("failed to save file: %w", err)
    }

    // Return the relative URL for the file
    return fmt.Sprintf("/public/%s", uniqueName), nil
}

// generateUniqueID generates a unique identifier for filenames
func generateUniqueID() string {
    b, _ := ioutil.ReadAll(os.Stdin)
    return fmt.Sprintf("%x", b)
}
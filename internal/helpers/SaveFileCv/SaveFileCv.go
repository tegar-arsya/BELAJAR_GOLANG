package SaveFileCv

import (
    "os"
    "path/filepath"
    "github.com/gin-gonic/gin"
    "log"
)

// SaveFile - Utility function to save uploaded files locally
func SaveFile(c *gin.Context, field string) (string, error) {
    file, err := c.FormFile(field)
    if err != nil {
        return "", err
    }

    // Define the upload path
    uploadPath := "./public"
    if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
        log.Printf("Failed to create directory: %v", err)
        return "", err
    }

    // Generate a unique filename
    fileName := filepath.Base(file.Filename)
    filePath := filepath.Join(uploadPath, fileName)

    // Save the uploaded file to the server
    if err := c.SaveUploadedFile(file, filePath); err != nil {
        return "", err
    }

    return filepath.ToSlash(filepath.Join("public", fileName)), nil
}

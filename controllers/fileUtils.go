// controllers/file_utils.go
package controllers

import (
    "os"
    "path/filepath"
    "github.com/gin-gonic/gin"
)

// SaveFile - Utility function to save uploaded files locally
func SaveFile(c *gin.Context, field string) (string, error) {
    file, err := c.FormFile(field)
    if err != nil {
        return "", err
    }

    // Define the upload path
    uploadPath := "uploads/"
    err = os.MkdirAll(uploadPath, os.ModePerm)
    if err != nil {
        return "", err
    }

    // Generate a unique filename
    fileName := filepath.Base(file.Filename)
    filePath := filepath.Join(uploadPath, fileName)

    // Save the uploaded file to the server
    if err := c.SaveUploadedFile(file, filePath); err != nil {
        return "", err
    }

    return filePath, nil
}

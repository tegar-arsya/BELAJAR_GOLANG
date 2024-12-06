// controllers/form_verification_controller.go
package controllers

import (
    "net/http"
    "portfolio-backend/config"
    "portfolio-backend/models"
    "time"

    "github.com/gin-gonic/gin"
)

// VerifyForm - verifikasi form dengan menyimpan verifikasi ke dalam database
func VerifyForm(c *gin.Context) {
    var form models.Form
    var formVerification models.FormVerification

    // Ambil ID form dari parameter URL
    formID := c.Param("id")

    // Cek apakah form ada
    if err := config.DB.Where("id = ?", formID).First(&form).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Form not found"})
        return
    }

    // Ambil nama user yang sedang login
    username := c.MustGet("username").(string)

    // Buat data verifikasi
    formVerification = models.FormVerification{
        FormID:     form.ID,
        VerifiedBy: username,
        VerifiedAt: time.Now(),
    }

    // Simpan verifikasi ke database
    if err := config.DB.Create(&formVerification).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating form verification"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Form verified successfully", "verification": formVerification})
}

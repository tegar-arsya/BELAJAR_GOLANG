package controllers

import (
    "net/http"
    "portfolio-backend/config"
    "portfolio-backend/models"
    "github.com/gin-gonic/gin"
)

// CreateNews
func CreateNews(c *gin.Context) {
    var news models.News
    username := c.MustGet("username").(string) // Ambil username dari context

    // Mengikat JSON ke struct
    if err := c.ShouldBindJSON(&news); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Set author_name dari username
    news.AuthorName = username

    // Simpan berita ke database
    if err := config.DB.Create(&news).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating news"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "News created successfully"})
}

// GetAllNews
func GetAllNews(c *gin.Context) {
    var news []models.News

    // Ambil semua berita dari database
    if err := config.DB.Find(&news).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching news"})
        return
    }

    c.JSON(http.StatusOK, gin.H{ "data": news})
}

// UpdateNews
func UpdateNews(c *gin.Context) {
    id := c.Param("id")
    var news models.News

    // Cari berita berdasarkan ID
    if err := config.DB.First(&news, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
        return
    }

    // Mengikat JSON ke struct
    if err := c.ShouldBindJSON(&news); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Update berita di database
    if err := config.DB.Save(&news).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating news"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "News updated successfully", "news": news})
}

// DeleteNews
func DeleteNews(c *gin.Context) {
    id := c.Param("id")

    // Hapus berita berdasarkan ID
    if err := config.DB.Delete(&models.News{}, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "News deleted successfully"})
}

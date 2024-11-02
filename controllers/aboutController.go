package controllers

import (
    "net/http"
    "portfolio-backend/config"
    "portfolio-backend/models"
    "github.com/gin-gonic/gin"
)

// Create About
func CreateAbout(c *gin.Context) {
    var about models.About
    about.Title = c.PostForm("title")
    about.Description = c.PostForm("description")

    imagePath, err := SaveFile(c, "image")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
        return
    }

    about.ImageUrl = imagePath
    config.DB.Create(&about)
    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "About created successfully"})
}

// Get All About
func GetAbout(c *gin.Context) {
    var about []models.About
    config.DB.Find(&about)
    c.JSON(http.StatusOK, gin.H{ "data": about})
}

// Get Single About
func GetAboutByID(c *gin.Context) {
    var about models.About
    if err := config.DB.Where("id = ?", c.Param("id")).First(&about).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "error": "About not found!"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": about})
}

// Update About
func UpdateAbout(c *gin.Context) {
    var about models.About
    if err := config.DB.Where("id = ?", c.Param("id")).First(&about).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "error": "About not found!"})
        return
    }

    about.Title = c.PostForm("title")
    about.Description = c.PostForm("description")

    imagePath, err := SaveFile(c, "image")
    if err == nil {
        about.ImageUrl = imagePath
    }

    config.DB.Save(&about)
    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "About updated successfully"})
}

// Delete About
func DeleteAbout(c *gin.Context) {
    var about models.About
    if err := config.DB.Where("id = ?", c.Param("id")).First(&about).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "error": "About not found!"})
        return
    }
    config.DB.Delete(&about)
    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "About deleted"})
}

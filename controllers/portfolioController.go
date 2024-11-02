package controllers

import (
    "net/http"
    "portfolio-backend/config"
    "portfolio-backend/models"
    "github.com/gin-gonic/gin"
)

// Create Portfolio
func CreatePortfolio(c *gin.Context) {
    var portfolio models.Portfolio
    portfolio.Title = c.PostForm("title")
    portfolio.Description = c.PostForm("description")

    // Handle file upload
    imagePath, err := SaveFile(c, "image")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
        return
    }

    portfolio.ImageUrl = imagePath
    config.DB.Create(&portfolio)
    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Portfolio created successfully"})
}

// Get All Portfolios
func GetPortfolios(c *gin.Context) {
    var portfolios []models.Portfolio
    config.DB.Find(&portfolios)
    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": portfolios})
}

// Get Single Portfolio
func GetPortfolio(c *gin.Context) {
    var portfolio models.Portfolio
    if err := config.DB.Where("id = ?", c.Param("id")).First(&portfolio).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "error": "Portfolio not found!"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": portfolio})
}

// Update Portfolio
func UpdatePortfolio(c *gin.Context) {
    var portfolio models.Portfolio
    if err := config.DB.Where("id = ?", c.Param("id")).First(&portfolio).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "error": "Portfolio not found!"})
        return
    }

    portfolio.Title = c.PostForm("title")
    portfolio.Description = c.PostForm("description")

    imagePath, err := SaveFile(c, "image")
    if err == nil {
        portfolio.ImageUrl = imagePath
    }

    config.DB.Save(&portfolio)
    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Portfolio updated successfully"})
}

// Delete Portfolio
func DeletePortfolio(c *gin.Context) {
    var portfolio models.Portfolio
    if err := config.DB.Where("id = ?", c.Param("id")).First(&portfolio).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "error": "Portfolio not found!"})
        return
    }
    config.DB.Delete(&portfolio)
    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Portfolio deleted"})
}

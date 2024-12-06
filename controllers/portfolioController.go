package controllers

import (
    "net/http"
    "portfolio-backend/config"
    "portfolio-backend/models"
    "github.com/gin-gonic/gin"
    "portfolio-backend/helpers"  // Import helper

)


// Create Portfolio
func CreatePortfolio(c *gin.Context) {
    var portfolio models.Portfolio
    portfolio.Title = c.PostForm("title")
    portfolio.Description = c.PostForm("description")

    // Handle file upload
    imagePath, err := SaveFile(c, "image")
    if err != nil {
        helpers.Respond(c, http.StatusBadRequest, "Failed to upload image", gin.H{"error": err.Error()})
		return
    }

    portfolio.ImageUrl = imagePath
    if err := config.DB.Create(&portfolio).Error; err != nil {
		helpers.Respond(c, http.StatusInternalServerError, "Failed to create portfolio", nil)
		return
	}
	helpers.Respond(c, http.StatusCreated, "Portfolio created successfully", portfolio)
}

// Get All Portfolios
func GetPortfolios(c *gin.Context) {
    var portfolios []models.Portfolio
    if err := config.DB.Find(&portfolios).Error; err != nil {
		helpers.Respond(c, http.StatusInternalServerError, "Failed to fetch portfolios", nil)
		return
	}
    helpers.Respond(c, http.StatusOK, "Portfolios retrieved successfully", portfolios)
}

// Get Single Portfolio
func GetPortfolio(c *gin.Context) {
    var portfolio models.Portfolio
    if err := config.DB.Where("id = ?", c.Param("id")).First(&portfolio).Error; err != nil {
		helpers.Respond(c, http.StatusNotFound, "Portfolio not found", nil)
		return
	}
    helpers.Respond(c, http.StatusOK, "Portfolio retrieved successfully", portfolio)
}

// Update Portfolio
func UpdatePortfolio(c *gin.Context) {
    var portfolio models.Portfolio
    if err := config.DB.Where("id = ?", c.Param("id")).First(&portfolio).Error; err != nil {
		helpers.Respond(c, http.StatusNotFound, "Portfolio not found", nil)
		return
	}

    portfolio.Title = c.PostForm("title")
    portfolio.Description = c.PostForm("description")

    imagePath, err := SaveFile(c, "image")
    if err == nil {
        portfolio.ImageUrl = imagePath
    }

    if err := config.DB.Save(&portfolio).Error; err != nil {
		helpers.Respond(c, http.StatusInternalServerError, "Failed to update portfolio", nil)
		return
	}
	helpers.Respond(c, http.StatusOK, "Portfolio updated successfully", portfolio)
}

// Delete Portfolio
func DeletePortfolio(c *gin.Context) {
    var portfolio models.Portfolio
    if err := config.DB.Where("id = ?", c.Param("id")).First(&portfolio).Error; err != nil {
		helpers.Respond(c, http.StatusNotFound, "Portfolio not found", nil)
		return
	}
    if err := config.DB.Delete(&portfolio).Error; err != nil {
		helpers.Respond(c, http.StatusInternalServerError, "Failed to delete portfolio", nil)
		return
	}
    helpers.Respond(c, http.StatusOK, "Portfolio deleted successfully", nil)
}

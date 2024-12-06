package controllers

import (
    "net/http"
    "portfolio-backend/config"
    "portfolio-backend/models"
    "github.com/gin-gonic/gin"
    "portfolio-backend/helpers"
)

// Create About
func CreateAbout(c *gin.Context) {
    var about models.About
    about.Title = c.PostForm("title")
    about.Description = c.PostForm("description")

    imagePath, err := SaveFile(c, "image")
    if err != nil {
        helpers.Respond(c, http.StatusBadRequest, "Failed to upload image", gin.H{"error": err.Error()})
        return
    }

    about.ImageUrl = imagePath
    
    if err := config.DB.Create(&about).Error; err != nil {
        helpers.Respond(c, http.StatusInternalServerError, "Failed to create about", nil)
        return
    }
    helpers.Respond(c, http.StatusCreated, "About created successfully", about)
}

// Get All About
func GetAbout(c *gin.Context) {
    var about []models.About
    
    if err := config.DB.Find(&about).Error; err != nil {
        helpers.Respond(c, http.StatusInternalServerError, "Error fetching about", nil)
        return
    }
    helpers.Respond(c, http.StatusOK, "About fetched successfully", about)
}

// Get Single About
func GetAboutByID(c *gin.Context) {
    var about models.About
    if err := config.DB.Where("id = ?", c.Param("id")).First(&about).Error; err != nil {
        helpers.Respond(c, http.StatusNotFound,"About not found!", nil)
        return
    }
    helpers.Respond(c, http.StatusOK, "About retrieved successfully", about)
}

// Update About
func UpdateAbout(c *gin.Context) {
    var about models.About
    if err := config.DB.Where("id = ?", c.Param("id")).First(&about).Error; err != nil {
        helpers.Respond(c, http.StatusNotFound, "About not found!", nil)
        return
    }

    about.Title = c.PostForm("title")
    about.Description = c.PostForm("description")

    imagePath, err := SaveFile(c, "image")
    if err == nil {
        about.ImageUrl = imagePath
    }

    if err:= config.DB.Save(&about).Error; err != nil {
        helpers.Respond(c, http.StatusInternalServerError, "Failed to update about", nil)
        return
    }
    helpers.Respond(c, http.StatusOK, "About updated successfully", about)
}

// Delete About
func DeleteAbout(c *gin.Context) {
    var about models.About
    if err := config.DB.Where("id = ?", c.Param("id")).First(&about).Error; err != nil {
       helpers.Respond(c, http.StatusNotFound, "About not found!", nil)
        return
    }
    if err := config.DB.Delete(&about).Error; err != nil {
        helpers.Respond(c, http.StatusInternalServerError, "Failed to delete about", nil)
        return
    }
    helpers.Respond(c, http.StatusOK, "About deleted successfully", nil)
}

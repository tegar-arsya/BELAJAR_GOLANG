package controllers

import (
    "net/http"
    "portfolio-backend/config"
    "portfolio-backend/models"
    "github.com/gin-gonic/gin"
)

// CreateForm
func CreateForm(c *gin.Context) {
    var form models.Form

    if err := c.ShouldBindJSON(&form); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Create(&form).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating form"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Form created successfully"})
}

// GetForms
func GetForms(c *gin.Context) {
    var forms []models.Form

    if err := config.DB.Find(&forms).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching forms"})
        return
    }

    c.JSON(http.StatusOK, forms)
}

// UpdateForm
func UpdateForm(c *gin.Context) {
    var form models.Form
    id := c.Param("id")

    if err := config.DB.Where("id = ?", id).First(&form).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Form not found"})
        return
    }

    if err := c.ShouldBindJSON(&form); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Save(&form).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating form"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Form updated successfully", "form": form})
}

// DeleteForm
func DeleteForm(c *gin.Context) {
    var form models.Form
    id := c.Param("id")

    if err := config.DB.Where("id = ?", id).First(&form).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Form not found"})
        return
    }

    if err := config.DB.Delete(&form).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting form"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Form deleted successfully"})
}
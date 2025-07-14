package helpers

import (
	"github.com/gin-gonic/gin"
	"time"
)

// Fungsi standar untuk format respons
func Respond(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, gin.H{
		"status":    statusCode,
		"message":   message,
		"timestamp": time.Now().Format(time.RFC3339),
		"data":      data,
	})
}

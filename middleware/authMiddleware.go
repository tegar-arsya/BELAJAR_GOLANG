// middleware/auth_middleware.go
package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"portfolio-backend/config"
	"portfolio-backend/controllers"
	"strings"
)

// Fungsi untuk memeriksa apakah token berada di daftar blacklist

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		// Pastikan token memiliki format "Bearer <token>"
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Ambil token dari header, tanpa "Bearer "
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Logging token yang diterima untuk debugging
		log.Printf("Received token: %s", tokenString)
		claims := &controllers.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return config.JwtKey, nil
		})

		if err != nil || !token.Valid {
			log.Printf("Invalid token: %v", err) // Log jika ada error parsing token
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Set username dari claims ke context
		c.Set("username", claims.Username)
		c.Next()
	}
}

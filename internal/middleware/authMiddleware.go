// middleware/auth_middleware.go
package middleware

import (
	"log"
	"net/http"
	"portfolio-backend/config"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Claims struct untuk JWT
type Claims struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	jwt.StandardClaims
}

// Fungsi validasi token JWT
func validateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return config.JwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	// Validasi apakah token sudah kadaluarsa
	if time.Unix(claims.ExpiresAt, 0).Before(time.Now()) {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}

// AuthMiddleware untuk validasi JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		// Pastikan token memiliki format "Bearer <token>"
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "Authorization header is required",
			})
			c.Abort()
			return
		}

		// Ambil token dari header, tanpa "Bearer "
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Validasi token
		claims, err := validateToken(tokenString)
		if err != nil {
			log.Printf("Invalid token: %v", err) // Logging error
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "Invalid or expired token",
			})
			c.Abort()
			return
		}

		// Set data dari token ke context
		c.Set("id", claims.Id)
		c.Set("username", claims.Username)
		c.Set("email", claims.Email)
		c.Set("claims", claims)
		c.Next()
	}
}

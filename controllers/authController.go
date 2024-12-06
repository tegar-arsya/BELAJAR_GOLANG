package controllers

import (
    "net/http"
    "portfolio-backend/config"
    "portfolio-backend/models"
    "time"

    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
)

type Claims struct {
    Username string `json:"username"`
    Role     string `json:"role"` // Tambahkan role ke dalam token
    jwt.StandardClaims
}

// Register
func Register(c *gin.Context) {
    var input struct {
        Username string `json:"username" binding:"required"`
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required,min=6"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  "error",
            "message": "Invalid input",
            "details": err.Error(),
        })
        return
    }

    // Cek apakah email sudah terdaftar
    var existingUser models.User
    if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
        c.JSON(http.StatusConflict, gin.H{
            "status":  "error",
            "message": "Email already registered",
        })
        return
    }

    // Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  "error",
            "message": "Error hashing password",
        })
        return
    }

    // Buat user baru
    newUser := models.User{
        Username: input.Username,
        Email:    input.Email,
        Password: string(hashedPassword),
        Role:     "user", // Role default
    }

    if err := config.DB.Create(&newUser).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  "error",
            "message": "Error creating user",
        })
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "status":  "success",
        "message": "User registered successfully",
        "data": gin.H{
            "id":       newUser.ID,
            "username": newUser.Username,
            "email":    newUser.Email,
            "role":     newUser.Role,
            "created":  newUser.CreatedAt,
        },
    })
}

// Login
func Login(c *gin.Context) {
    var input struct {
        Username string `json:"username" binding:"required"`
        Password string `json:"password" binding:"required"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  "error",
            "message": "Invalid input",
            "details": err.Error(),
        })
        return
    }

    var user models.User

    // Cari user berdasarkan username
    if err := config.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{
            "status":  "error",
            "message": "Invalid username or password",
        })
        return
    }

    // Cek password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{
            "status":  "error",
            "message": "Invalid username or password",
        })
        return
    }

    // Buat JWT token
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        Username: user.Username,
        Role:     user.Role,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(config.JwtKey)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  "error",
            "message": "Error generating token",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": "Login successful",
        "data": gin.H{
            "username": user.Username,
            "role":     user.Role,
            "token":    token,
        },
    })
}



func Logout(c *gin.Context) {
    // Instruksikan client untuk menghapus token
    c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
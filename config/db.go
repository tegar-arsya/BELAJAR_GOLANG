package config

import (
    "fmt"
    "os"

    "github.com/joho/godotenv"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB
var JwtKey []byte

func ConnectDatabase() {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        panic("Error loading .env file")
    }

    // Ambil variabel dari .env
    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASS")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")
    JwtKey = []byte(os.Getenv("JWT_SECRET"))

    // Format DSN
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
        dbUser, dbPass, dbHost, dbPort, dbName)

    // Koneksi ke database
    database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database!")
    }

    DB = database
}

package config

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

// Secret Key untuk JWT, dibuat publik agar bisa diakses dari package lain
var JwtKey = []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9")

func ConnectDatabase() {
    dsn := "root:@tcp(127.0.0.1:3306)/portfolio_db?charset=utf8mb4&parseTime=True&loc=Local"
    database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database!")
    }

    DB = database
}

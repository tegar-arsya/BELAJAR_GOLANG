package main

import (
    "github.com/gin-gonic/gin"
    "portfolio-backend/config"
    "portfolio-backend/middleware"
    "portfolio-backend/controllers"
)

func main() {
    r := gin.Default()

    // Connect to Database
    config.ConnectDatabase()

    // Auth Routes
    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)
    r.POST("/logout", middleware.AuthMiddleware(), controllers.Logout)  // Route untuk logout
    // Protected Routes
    protectedRoutes := r.Group("/")
    protectedRoutes.Use(middleware.AuthMiddleware())

    // Portfolio Routes
    portfolio := protectedRoutes.Group("portfolios")
    {
        portfolio.POST("/", controllers.CreatePortfolio)
        portfolio.GET("/", controllers.GetPortfolios)
        portfolio.GET("/:id", controllers.GetPortfolio)
        portfolio.PUT("/:id", controllers.UpdatePortfolio)
        portfolio.DELETE("/:id", controllers.DeletePortfolio)
    }

    // About Routes
    about := protectedRoutes.Group("about")
    {
        about.POST("/", controllers.CreateAbout)
        about.GET("/", controllers.GetAbout)
        about.PUT("/:id", controllers.UpdateAbout)
        about.DELETE("/:id", controllers.DeleteAbout)
    }

    // News Routes
    news := protectedRoutes.Group("news")
    {
        news.POST("/", controllers.CreateNews) // Rute untuk membuat berita
        news.GET("/", controllers.GetAllNews)   // Rute untuk mendapatkan semua berita
        news.PUT("/:id", controllers.UpdateNews) // Rute untuk update berita
        news.DELETE("/:id", controllers.DeleteNews) // Rute untuk delete berita
    }

    r.Run() // Default port 8080
}
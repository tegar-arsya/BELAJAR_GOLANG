package main

import (
    "github.com/gin-gonic/gin"
    "portfolio-backend/config"
    "portfolio-backend/migrations"
    "portfolio-backend/middleware"
    "portfolio-backend/controllers"
    "portfolio-backend/routes"
)

func main() {
    r := gin.Default()

    // Connect to Database
    config.ConnectDatabase()
    // Lakukan migrasi
    migrations.Migrate()
    // Auth Routes
    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)
    r.POST("/logout", middleware.AuthMiddleware(), controllers.Logout)  // Route untuk logout
    // Protected Routes
    protectedRoutes := r.Group("/")
    protectedRoutes.Use(middleware.AuthMiddleware())

    // Portfolio Routes
    // portfolio := protectedRoutes.Group("portfolios")
    // {
    //     portfolio.POST("/", controllers.CreatePortfolio)
    //     portfolio.GET("/", controllers.GetPortfolios)
    //     portfolio.GET("/:id", controllers.GetPortfolio)
    //     portfolio.PUT("/:id", controllers.UpdatePortfolio)
    //     portfolio.DELETE("/:id", controllers.DeletePortfolio)
    // }
    routes.PortfolioRoutes(protectedRoutes)

    
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

    forms := protectedRoutes.Group("forms")
    {
        forms.POST("/", controllers.CreateForm)
        forms.GET("/", controllers.GetForms)
        forms.PUT("/:id", controllers.UpdateForm)
        forms.DELETE("/:id", controllers.DeleteForm)
    }
    formsVeification := protectedRoutes.Group("forms-verification")
    {
        formsVeification.POST("/:id/verify", controllers.VerifyForm)
    }

    r.Run() // Default port 8080
}

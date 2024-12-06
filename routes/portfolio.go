package routes

import (
	"portfolio-backend/controllers"
	"github.com/gin-gonic/gin"
)

func PortfolioRoutes(router *gin.RouterGroup) {
	portfolio := router.Group("portfolios")
	{
		portfolio.POST("/", controllers.CreatePortfolio)
		portfolio.GET("/", controllers.GetPortfolios)
		portfolio.GET("/:id", controllers.GetPortfolio)
		portfolio.PUT("/:id", controllers.UpdatePortfolio)
		portfolio.DELETE("/:id", controllers.DeletePortfolio)
	}
}
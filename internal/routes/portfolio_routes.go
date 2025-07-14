package routes

import (
	"portfolio-backend/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterPortfolioRoutes(public, admin *gin.RouterGroup, h *handler.PortofolioHandler) {
	publicGroup := public.Group("portfolios")
	{
		publicGroup.GET("/", h.GetAll)
		publicGroup.GET("/:id", h.GetByID)
	}

	adminGroup := admin.Group("portfolio")
	{
		adminGroup.POST("/", h.Create)
		adminGroup.GET("/", h.GetAll)
		adminGroup.GET("/:id", h.GetByID)
		adminGroup.PUT("/:id", h.Update)
		adminGroup.DELETE("/:id", h.Delete)
	}
}

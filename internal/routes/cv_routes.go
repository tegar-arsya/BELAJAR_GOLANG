package routes

import (
	"portfolio-backend/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterCvRoutes(public, admin *gin.RouterGroup, h *handler.CvHandler) {
	publicGroup := public.Group("cvs")
	{
		publicGroup.GET("/", h.GetAll)
	}

	adminGroup := admin.Group("cv")
	{
		adminGroup.POST("/", h.Create)
		adminGroup.GET("/", h.GetAll)
		adminGroup.DELETE("/:id", h.Delete)
	}
}

package routes

import (
	"portfolio-backend/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterPengalamanRoutes(public, admin *gin.RouterGroup, h *handler.PengalamanHandler) {
	publicGroup := public.Group("pengalamans")
	{
		publicGroup.GET("/", h.GetAll)
	}

	adminGroup := admin.Group("pengalaman")
	{
		adminGroup.POST("/", h.Create)
		adminGroup.GET("/", h.GetAll)
		adminGroup.DELETE("/:id", h.Delete)
	}
}

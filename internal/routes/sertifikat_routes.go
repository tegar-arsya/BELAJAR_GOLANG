package routes

import (
	"portfolio-backend/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterSertifikatRoutes(public, admin *gin.RouterGroup, h *handler.SertifikatHandler) {
	publicGroup := public.Group("sertifikats")
	{
		publicGroup.GET("/", h.GetAll)
        publicGroup.GET("/:id", h.GetByID)
	}

	adminGroup := admin.Group("sertifikat")
	{
		adminGroup.POST("/", h.Create)
		adminGroup.GET("/", h.GetAll)
        adminGroup.GET("/:id", h.GetByID)
		adminGroup.DELETE("/:id", h.Delete)
	}
}

package routes

import (
	"portfolio-backend/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterAboutRoutes(public, admin *gin.RouterGroup, h *handler.AboutHandler) {
	publicAbout := public.Group("abouts")
	{
		publicAbout.GET("/", h.GetAll)
	}

	adminAbout := admin.Group("about")
	{
		adminAbout.POST("/", h.Create)
		adminAbout.GET("/", h.GetAll)
		adminAbout.PUT("/:id", h.Update)
		adminAbout.DELETE("/:id", h.Delete)
	}
}

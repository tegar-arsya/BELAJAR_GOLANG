package routes

import (
	"portfolio-backend/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterArticleRoutes(public *gin.RouterGroup, admin *gin.RouterGroup, h *handler.ArticleHandler) {
	public.GET("/articles", h.GetAll)
	public.GET("/articles/:id", h.GetByID)

	admin.POST("/article", h.Create)
	admin.PUT("/article/:id", h.Update)
	admin.DELETE("/article/:id", h.Delete)
    admin.POST("/upload/article-image", h.UploadImage)

}

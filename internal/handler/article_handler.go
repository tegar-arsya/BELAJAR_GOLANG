package handler

import (
	"net/http"
	"portfolio-backend/internal/domain/entity"
	"portfolio-backend/internal/helpers"
	"portfolio-backend/internal/service"
	"portfolio-backend/internal/storage"
	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	Service *service.ArticleService
}

func (h *ArticleHandler) Create(c *gin.Context) {
	var req entity.Article
	req.Title = c.PostForm("title")
	req.Content = c.PostForm("content")

	files, err := storage.SaveMultipleFiles(c, "thumbnail", "articles")
	if err == nil && len(files) > 0 {
		req.Thumbnail = files[0] // only first
	}

	if err := h.Service.Create(&req); err != nil {
		helpers.Respond(c, http.StatusInternalServerError, "Create failed", nil)
		return
	}

	helpers.Respond(c, http.StatusCreated, "Created", req)
}

func (h *ArticleHandler) GetAll(c *gin.Context) {
	data, err := h.Service.GetAll()
	if err != nil {
		helpers.Respond(c, http.StatusInternalServerError, "Fetch failed", nil)
		return
	}
	helpers.Respond(c, http.StatusOK, "Success", data)
}

func (h *ArticleHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	data, err := h.Service.GetByID(id)
	if err != nil {
		helpers.Respond(c, http.StatusNotFound, "Not found", nil)
		return
	}
	helpers.Respond(c, http.StatusOK, "Success", data)
}

func (h *ArticleHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req entity.Article

	req.Title = c.PostForm("title")
	req.Content = c.PostForm("content")

	uintID, err := helpers.StringToUint(id)
	if err != nil {
		helpers.Respond(c, http.StatusBadRequest, "Invalid ID", nil)
		return
	}
	req.ID = uintID

	files, _ := storage.SaveMultipleFiles(c, "thumbnail", "articles")
	if len(files) > 0 {
		req.Thumbnail = files[0]
	}

	if err := h.Service.Update(&req); err != nil {
		helpers.Respond(c, http.StatusInternalServerError, "Update failed", nil)
		return
	}
	helpers.Respond(c, http.StatusOK, "Updated", req)
}

func (h *ArticleHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.Service.Delete(id); err != nil {
		helpers.Respond(c, http.StatusInternalServerError, "Delete failed", nil)
		return
	}
	helpers.Respond(c, http.StatusOK, "Deleted", nil)
}

func (h *ArticleHandler) UploadImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		helpers.Respond(c, http.StatusBadRequest, "No image uploaded", nil)
		return
	}

	// Simpan file ke direktori public/article-content/
	filename := helpers.GenerateFileName(file.Filename) // kamu bisa buat helper atau langsung kasih nama random
	path := "public/article-content/" + filename

	if err := c.SaveUploadedFile(file, path); err != nil {
		helpers.Respond(c, http.StatusInternalServerError, "Failed to save image", nil)
		return
	}

	// Kirim URL lengkap ke FE
	fullURL := "http://localhost:9191/article-content/" + filename
	c.JSON(http.StatusOK, gin.H{"url": fullURL})
}

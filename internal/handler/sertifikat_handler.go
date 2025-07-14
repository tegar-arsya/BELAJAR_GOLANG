package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"portfolio-backend/internal/domain/entity"
	"portfolio-backend/internal/helpers"
	"portfolio-backend/internal/service"
	"portfolio-backend/internal/storage"
)

type SertifikatHandler struct {
	Service *service.SertifikatService
}

// ✅ CREATE
func (h *SertifikatHandler) Create(c *gin.Context) {
	var req entity.Sertifikat
	req.Title = c.PostForm("title")
	req.Description = c.PostForm("description")
	req.Site = c.PostForm("site")

	files, err := storage.SaveMultipleFiles(c, "image_urls", "serti")
	if err != nil {
		helpers.Respond(c, http.StatusBadRequest, "File upload failed", nil)
		return
	}

	req.ImageUrl = strings.Join(files, ",")

	if err := h.Service.Create(&req); err != nil {
		helpers.Respond(c, http.StatusInternalServerError, "Create failed", nil)
		return
	}

	helpers.Respond(c, http.StatusCreated, "Created", req)
}

// ✅ GET ALL
func (h *SertifikatHandler) GetAll(c *gin.Context) {
	result, err := h.Service.GetAll()
	if err != nil {
		helpers.Respond(c, http.StatusInternalServerError, "Fetch failed", nil)
		return
	}

	helpers.Respond(c, http.StatusOK, "Success", result)
}

// ✅ GET BY ID
func (h *SertifikatHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	result, err := h.Service.GetByID(id)
	if err != nil {
		helpers.Respond(c, http.StatusNotFound, "Sertifikat not found", nil)
		return
	}
	helpers.Respond(c, http.StatusOK, "Success", result)
}

// ✅ DELETE
func (h *SertifikatHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.Service.Delete(id); err != nil {
		helpers.Respond(c, http.StatusInternalServerError, "Delete failed", nil)
		return
	}
	helpers.Respond(c, http.StatusOK, "Deleted", nil)
}

// ✅ UPDATE
func (h *SertifikatHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req entity.Sertifikat

	req.Title = c.PostForm("title")
	req.Description = c.PostForm("description")
	req.Site = c.PostForm("site")

	// ID wajib diisi untuk GORM.Save()
	uintID, err := helpers.StringToUint(id)
	if err != nil {
		helpers.Respond(c, http.StatusBadRequest, "Invalid ID", nil)
		return
	}
	req.ID = uintID

	// Jika ada file baru
	files, err := storage.SaveMultipleFiles(c, "image_urls", "serti")
	if err == nil && len(files) > 0 {
		req.ImageUrl = strings.Join(files, ",")
	} else {
		// Kalau tidak upload ulang, ambil image lama
		existing, err := h.Service.GetByID(id)
		if err != nil {
			helpers.Respond(c, http.StatusNotFound, "Sertifikat not found", nil)
			return
		}
		req.ImageUrl = existing.ImageUrl
	}

	if err := h.Service.Update(&req); err != nil {
		helpers.Respond(c, http.StatusInternalServerError, "Update failed", nil)
		return
	}

	helpers.Respond(c, http.StatusOK, "Updated", req)
}

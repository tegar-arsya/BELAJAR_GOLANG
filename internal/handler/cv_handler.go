package handler

import (
	"net/http"
	"portfolio-backend/internal/domain/entity"
	"portfolio-backend/internal/helpers"
	"portfolio-backend/internal/service"
	"portfolio-backend/internal/storage"
	"strings"
	"github.com/gin-gonic/gin"
)


type CvHandler struct {
    Service *service.CvService
}

func (h *CvHandler) Create(c *gin.Context) {
    var req entity.Cv
    // Mengambil file dari form data
    file, err := storage.SaveFile(c, "file", "file")
    if err != nil {
        helpers.Respond(c, http.StatusBadRequest, "File upload failed", nil)
        return
    }

    req.File = strings.Join(file, ",") // Jika ada lebih dari satu file, bisa disesuaikan
    if err := h.Service.Create(&req); err != nil {
        helpers.Respond(c, http.StatusInternalServerError, "Create failed", nil)
        return
    }
    helpers.Respond(c, http.StatusCreated, "Created", req)
}

func (h *CvHandler) GetAll(c *gin.Context) {
    result, err := h.Service.GetAll()
    if err != nil {
        helpers.Respond(c, http.StatusInternalServerError, "Fetch failed", nil)
        return
    }
    helpers.Respond(c, http.StatusOK, "Success", result)
}

func (h *CvHandler) Delete(c *gin.Context) {
    id := c.Param("id")
    if err := h.Service.Delete(id); err != nil {
        helpers.Respond(c, http.StatusInternalServerError, "Delete failed", nil)
        return
    }
    helpers.Respond(c, http.StatusOK, "Deleted", nil)
}

// func (h *CvHandler) GetByID(c *gin.Context) {
//     id := c.Param("id")
//     result, err := h.Service.GetByID(id)
//     if err != nil {
//         helpers.Respond(c, http.StatusNotFound, "CV not found", nil)
//         return
//     }
//     helpers.Respond(c, http.StatusOK, "Success", result)
// }

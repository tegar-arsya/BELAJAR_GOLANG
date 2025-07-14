package handler


import (
    "net/http"
    "portfolio-backend/internal/domain/entity"
    "portfolio-backend/internal/helpers"
    "portfolio-backend/internal/service"
    "portfolio-backend/internal/storage"
    "github.com/gin-gonic/gin"
    "strings"
)


type PengalamanHandler struct {
    Service *service.PengalamanService
}

func (h *PengalamanHandler) Create(c *gin.Context) {
    var req entity.Pengalaman
    req.Title = c.PostForm("title")
    req.Description = c.PostForm("description")

    files, err := storage.SaveMultipleFiles(c, "image_urls", "joruney")
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

func (h *PengalamanHandler) GetAll(c *gin.Context) {
    result, err := h.Service.GetAll()
    if err != nil {
        helpers.Respond(c, http.StatusInternalServerError, "Fetch failed", nil)
        return
    }
    helpers.Respond(c, http.StatusOK, "Success", result)
}

func (h *PengalamanHandler) Delete(c *gin.Context) {
    id := c.Param("id")
    if err := h.Service.Delete(id); err != nil {
        helpers.Respond(c, http.StatusInternalServerError, "Delete failed", nil)
        return
    }
    helpers.Respond(c, http.StatusOK, "Deleted", nil)
}

func (h *PengalamanHandler) GetByID(c *gin.Context) {
    id := c.Param("id")
    result, err := h.Service.GetByID(id)
    if err != nil {
        helpers.Respond(c, http.StatusNotFound, "Pengalaman not found", nil)
        return
    }
    helpers.Respond(c, http.StatusOK, "Success", result)
}

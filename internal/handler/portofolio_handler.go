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


type PortofolioHandler struct {
    Service *service.PortofolioService
}

func (h *PortofolioHandler) Create(c *gin.Context) {
    var req entity.Portfolio
    req.Title = c.PostForm("title")
    req.Description = c.PostForm("description")
    req.Site = c.PostForm("site")
    req.GithubUrl = c.PostForm("github_url")

    files, err := storage.SaveMultipleFiles(c, "image_urls", "porto")
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

func (h *PortofolioHandler) GetAll(c *gin.Context) {
    result, err := h.Service.GetAll()
    if err != nil {
        helpers.Respond(c, http.StatusInternalServerError, "Fetch failed", nil)
        return
    }

    helpers.Respond(c, http.StatusOK, "Success", result)
}

func (h *PortofolioHandler) GetByID(c *gin.Context) {
    id := c.Param("id")
    result, err := h.Service.GetByID(id)
    if err != nil {
        helpers.Respond(c, http.StatusNotFound, "Portfolio not found", nil)
        return
    }

    helpers.Respond(c, http.StatusOK, "Success", result)
}
func (h *PortofolioHandler) Update(c *gin.Context) {
    id := c.Param("id")
    var req entity.Portfolio

    req.Title = c.PostForm("title")
    req.Description = c.PostForm("description")
    req.Site = c.PostForm("site")
    req.GithubUrl = c.PostForm("github_url")

    // Convert id dari string ke uint
    uintID, err := helpers.StringToUint(id)
    if err != nil {
        helpers.Respond(c, http.StatusBadRequest, "Invalid ID", nil)
        return
    }
    req.ID = uintID

    // Jika ada file baru
    files, err := storage.SaveMultipleFiles(c, "image_urls", "portofolio")
    if err == nil && len(files) > 0 {
        req.ImageUrl = strings.Join(files, ",")
    } else {
        // Ambil data lama supaya tidak kosongkan image
        existing, err := h.Service.GetByID(id)
        if err != nil {
            helpers.Respond(c, http.StatusNotFound, "Portfolio not found", nil)
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

// Note: Update method should handle the case where no file is uploaded, so it doesn't overwrite
// the existing image URL with an empty string. You can add a check before saving the files.
func (h *PortofolioHandler) Delete(c *gin.Context) {
    id := c.Param("id")
    if err := h.Service.Delete(id); err != nil {
        helpers.Respond(c, http.StatusInternalServerError, "Delete failed", nil)
        return
    }

    helpers.Respond(c, http.StatusOK, "Deleted", nil)
}

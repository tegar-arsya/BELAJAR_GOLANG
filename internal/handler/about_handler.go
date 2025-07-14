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


type AboutHandler struct {
    Service *service.AboutService
}

func (h *AboutHandler) Create(c *gin.Context) {
    var req entity.About
    req.Title = c.PostForm("title")
    req.Description = c.PostForm("description")

    files, err := storage.SaveMultipleFiles(c, "image_url", "me")
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


func (h *AboutHandler) GetAll(c *gin.Context) {
    result, err := h.Service.GetAll()
    if err != nil {
        helpers.Respond(c, http.StatusInternalServerError, "Fetch failed", nil)
        return
    }

    helpers.Respond(c, http.StatusOK, "Success", result)
}

func (h *AboutHandler) Update(c *gin.Context) {
    id := c.Param("id")
    var req entity.About

    // Convert id from string to uint
    uintID, err := helpers.StringToUint(id)
    if err != nil {
        helpers.Respond(c, http.StatusBadRequest, "Invalid ID", nil)
        return
    }
    req.ID = uintID
    req.Title = c.PostForm("title")
    req.Description = c.PostForm("description")

    files, err := storage.SaveMultipleFiles(c, "image_url", "me")
    if err != nil {
        helpers.Respond(c, http.StatusBadRequest, "File upload failed", nil)
        return
    }

    req.ImageUrl = strings.Join(files, ",")

    if err := h.Service.Update(&req); err != nil {
        helpers.Respond(c, http.StatusInternalServerError, "Update failed", nil)
        return
    }

    helpers.Respond(c, http.StatusOK, "Updated", req)
}

func (h *AboutHandler) Delete(c *gin.Context) {
    id := c.Param("id")
    if err := h.Service.Delete(id); err != nil {
        helpers.Respond(c, http.StatusInternalServerError, "Delete failed", nil)
        return
    }

    helpers.Respond(c, http.StatusOK, "Deleted", nil)
}

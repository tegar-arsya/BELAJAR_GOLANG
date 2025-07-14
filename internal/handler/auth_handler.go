package handler

import (
	"net/http"
	"portfolio-backend/internal/service"
	"portfolio-backend/internal/domain/entity"
	"portfolio-backend/internal/domain/dto"
	"portfolio-backend/internal/helpers"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service *service.AuthService
}

func (h *AuthHandler) Register(c *gin.Context) {
	var input dto.RegisterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.Respond(c, http.StatusBadRequest, "Input salah", nil)
		return
	}

	user := entity.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}

	err := h.Service.Register(user)
	if err != nil {
		helpers.Respond(c, http.StatusConflict, "Email sudah digunakan", nil)
		return
	}

	helpers.Respond(c, http.StatusCreated, "Register berhasil", nil)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input dto.LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.Respond(c, http.StatusBadRequest, "Input salah", nil)
		return
	}

	user, token, err := h.Service.Login(input.Username, input.Password)
	if err != nil {
		helpers.Respond(c, http.StatusUnauthorized, "Username/password salah", nil)
		return
	}

	helpers.Respond(c, http.StatusOK, "Login sukses", gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"token":    token,
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	// client tinggal hapus token
	helpers.Respond(c, http.StatusOK, "Logout sukses", nil)
}

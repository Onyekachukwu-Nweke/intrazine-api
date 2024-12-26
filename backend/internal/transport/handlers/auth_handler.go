package handlers

import (
	"fmt"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/interfaces"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"regexp"
)

type AuthHandler struct {
	Service interfaces.AuthService
}

func NewAuthHandler(service interfaces.AuthService) *AuthHandler {
	return &AuthHandler{Service: service}
}

type UserSignupRequest struct {
	Username        string `json:"username" valid:"required"`
	Email           string `json:"email" valid:"email,required"`
	Password        string `json:"password" valid:"required"`
	PasswordConfirm string `json:"password_confirm" valid:"required"`
}

// regex for validating an email
var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}

func (h *AuthHandler) Signup(c *gin.Context) {
	var body UserSignupRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if body.Password != body.PasswordConfirm {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		return
	}

	if !isEmailValid(body.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
	}

	exists, field, err := h.Service.CheckUserExists(c.Request.Context(), body.Username, body.Email)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}
	if exists {
		msg := fmt.Sprintf("%s already exists", field)
		c.JSON(http.StatusConflict, gin.H{"error": msg})
	}

}

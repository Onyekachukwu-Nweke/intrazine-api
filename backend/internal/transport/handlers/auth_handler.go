package handlers

import (
	"fmt"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/interfaces"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/models"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"regexp"
	"time"
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

type LoginRequest struct {
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required"`
}

type ForgotPasswordRequest struct {
	Username string `json:"username" valid:"required"`
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
		return
	}

	exists, field, err := h.Service.CheckUserExists(c.Request.Context(), body.Username, body.Email)
	if err != nil {
		log.Print(err)
	}
	if exists {
		msg := fmt.Sprintf("%s already exists", field)
		c.JSON(http.StatusConflict, gin.H{"error": msg})
		return
	}

	passwordHash, err := utils.HashPassword(body.Password)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	user := models.User{
		Username:     body.Username,
		Email:        body.Email,
		PasswordHash: passwordHash,
	}

	createdUser, err := h.Service.Signup(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		log.Print(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": createdUser})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var body LoginRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := h.Service.Login(c.Request.Context(), body.Username, body.Password)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "username or password not correct"})
		return
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("missionimpossible"))
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})

}

func (h *AuthHandler) ForgotPassword(c *gin.Context) {
	var body ForgotPasswordRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username"})
	}

	token, err := h.Service.ForgotPassword(c.Request.Context(), body.Username)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset token sent", "token": token})
}

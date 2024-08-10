package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/user"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/utils"
	"github.com/golang-jwt/jwt"
)

type UserService interface {
	CreateUser(ctx context.Context, user user.User) (user.User, error)
	CheckUserExists(ctx context.Context, username, email string) (exists bool, field string, err error)
	Login(ctx context.Context, username, password string) (user.User, error)
}

type UserRequest struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func convertUserRequestToUser(u UserRequest, PasswordHash string) user.User {
	return user.User{
		Username: u.Username,
		Email: u.Email,
		PasswordHash: PasswordHash,
	}
}

// regex for validating an email
var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
			return false
	}
	return emailRegex.MatchString(e)
}

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var usrReq UserRequest
	if err := json.NewDecoder(r.Body).Decode(&usrReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "Invalid Request Body"})
		return 
	}

	if usrReq.Password != usrReq.PasswordConfirm {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "Passwords do not match"})
		return 
	}

	// Validate email format
	if !isEmailValid(usrReq.Email) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "Invalid email format"})
		return
	}

	exists, field, err := h.UserService.CheckUserExists(r.Context(), usrReq.Username, usrReq.Email)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{Message: "Internal server error"})
	}
	if exists {
		msg := fmt.Sprintf("%s already exists", field)
		json.NewEncoder(w).Encode(Response{Message: msg})
		return
	}

	// Proceed to create user if validations pass
	passwordHash, err := utils.HashPassword(usrReq.Password)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{Message: "Internal server error"})
		return
	}

	convertedUser := convertUserRequestToUser(usrReq, passwordHash)
	createdUser, err := h.UserService.CreateUser(r.Context(), convertedUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{Message: "User Sign Up Failed"})
		// TODO: Add Logger
		log.Print(err)
		return
	}

	// TODO: Add Logger
	log.Print("Signup Successful")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": createdUser.ID})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var logReq LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&logReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "Invalid Request Body"})
		return
	}

	user, err := h.UserService.Login(r.Context(), logReq.Username, logReq.Password)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(Response{Message: "username or password not correct"})
		return
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("missionimpossible"))
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{Message: "Internal server error"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
package handlers

import (
	"encoding/json"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/interfaces"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/models"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

type PostHandler struct {
	Service interfaces.PostService
}

func NewPostHandler(service interfaces.PostService) *PostHandler {
	return &PostHandler{Service: service}
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(string) // Retrieve user ID from context
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	post.UserId = userID
	validate := validator.New()
	err := validate.Struct(post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		// TODO: Add Logger
		log.Print(err)
		return
	}

	createdPost, err := h.Service.CreatePost(r.Context(), post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//json.NewEncoder(w).Encode(Response{Message: "Post could not created"})
		// TODO: Add Logger
		log.Print(err)
		return
	}

	// TODO: Add Logger
	log.Print("Post Successfully Created")

	if err := json.NewEncoder(w).Encode(createdPost); err != nil {
		panic(err)
	}
}

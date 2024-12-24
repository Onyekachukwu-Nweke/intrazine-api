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
	Service interfaces.PostStore
}

func NewPostHandler(service interfaces.PostStore) *PostHandler {
	return &PostHandler{Service: service}
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	validate := validator.New()
	err := validate.Struct(post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//json.NewEncoder(w).Encode(Response{Message: "Post not valid"})
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

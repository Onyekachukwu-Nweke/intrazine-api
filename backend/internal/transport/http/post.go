package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/post"
	"github.com/go-playground/validator/v10"
)


type PostService interface {
	CreatePost(context.Context, post.Post) (post.Post, error)
	GetPost(ctx context.Context, ID string) (post.Post, error)
}

type Response struct {
	Message string
}

type PostRequest struct {
	User_id string `json:"user_id" validate:"required"`
	Title string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

func convertPostRequestToPost(p PostRequest) post.Post {
	return post.Post{
		User_id: p.User_id,
		Title: p.Title,
		Content: p.Content,
	}
}

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var pst PostRequest
	if err := json.NewDecoder(r.Body).Decode(&pst); err != nil {
		return
	}

	validate := validator.New()
	err := validate.Struct(pst)
	if err != nil {
		http.Error(w, "not a valid post", http.StatusBadRequest)
		return
	}

	convertedPost := convertPostRequestToPost(pst)
	createdPost, err := h.PostService.CreatePost(r.Context(), convertedPost)
	if err != nil {
		log.Print(err)
		return
	}

	if err := json.NewEncoder(w).Encode(createdPost); err != nil {
		panic(err)
	}
}
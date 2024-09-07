package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/post"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)


type PostService interface {
	CreatePost(context.Context, post.Post) (post.Post, error)
	GetPostByID(ctx context.Context, ID string) (post.Post, error)
	GetAllPosts(context.Context) ([]post.Post, error)
	UpdatePost(ctx context.Context, ID string, updatedPost post.Post) (post.Post, error)
	DeletePost(ctx context.Context, ID string) error
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

// Create Post Function on the transport layer
func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID, err  := utils.GetUserIDFromContext(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(Response{Message: "Not Authorized"})
		return
	}

	var pst PostRequest
	if err := json.NewDecoder(r.Body).Decode(&pst); err != nil {
		return
	}

	pst.User_id = userID

	validate := validator.New()
	err = validate.Struct(pst)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "Post not valid"})
		// TODO: Add Logger
		log.Print(err)
		return
	}

	convertedPost := convertPostRequestToPost(pst)
	createdPost, err := h.PostService.CreatePost(r.Context(), convertedPost)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "Post could not created"})
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


/************** GetPostByID (Transport Layer) ************/
func (h *Handler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "Missing Post ID"})
		return
	}

	pst, err := h.PostService.GetPostByID(r.Context(), id)
	if err != nil {
		if strings.Contains(err.Error(), "no post found") {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(Response{Message: "Post not found"})
				return
		}
		log.Print(err) // TODO: Replace with structured logging
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{Message: "Internal server error"})
		return
	}

	if err := json.NewEncoder(w).Encode(pst); err != nil {
		panic(err)
	}
}

/************** GetAllPosts (Transport Layer) ************/
func (h *Handler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	psts, err := h.PostService.GetAllPosts(r.Context())
	if err != nil {
		log.Print(err) // TODO: Replace with structured logging
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{Message: "Internal server error"})
		return
	}

	if err := json.NewEncoder(w).Encode(psts); err != nil {
		panic(err)
	}
}

/* This endpoint needs a rework to avoid malicious
 * person's from modifing ppl's posts on the platform
 */
func (h *Handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	userID, err  := utils.GetUserIDFromContext(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(Response{Message: "Not Authorized"})
		return
	}

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "Missing Post ID"})
		return
	}

	var pst post.Post
	if err = json.NewDecoder(r.Body).Decode(&pst); err != nil {
		return
	}

	pst.User_id = userID

	// Use the centralized authorization service
	if !h.AuthorizationService.IsUserAuthorized(r.Context(), userID, id, "post") {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(Response{Message: "Forbidden"})
		return
	}

	updatedPost, err := h.PostService.UpdatePost(r.Context(), id, pst)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{Message: "Internal server error"})
		return
	}

	if err := json.NewEncoder(w).Encode(updatedPost); err != nil {
		panic(err)
	}
}


func (h *Handler) DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	userID, err  := utils.GetUserIDFromContext(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(Response{Message: "Not Authorized"})
		return
	}
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "Missing Post ID"})
		return
	}

	// Use the centralized authorization service
	if !h.AuthorizationService.IsUserAuthorized(r.Context(), userID, id, "post") {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(Response{Message: "Forbidden"})
		return
	}

	err = h.PostService.DeletePost(r.Context(), id)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{Message: "Internal server error"})
		return
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Successfully deleted"}); err != nil {
		panic(err)
	}
}
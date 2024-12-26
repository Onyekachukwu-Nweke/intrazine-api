package handlers

import (
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/interfaces"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/models"
	"github.com/gin-gonic/gin"
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

func (h *PostHandler) CreatePost(c *gin.Context) {
	// Extract user ID from the context (set by JWT middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Parse request body into the `models.Post` struct
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	// Set the user ID on the post
	post.UserId = userID.(string)

	// Validate the input
	validate := validator.New()
	if err := validate.Struct(post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}

	// Call the service layer to create the post
	createdPost, err := h.Service.CreatePost(c.Request.Context(), post)
	if err != nil {
		log.Printf("Failed to create post: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	// Log success and respond
	log.Printf("Post successfully created: %+v", createdPost)
	c.JSON(http.StatusCreated, gin.H{"message": "Post created successfully", "data": createdPost})
}

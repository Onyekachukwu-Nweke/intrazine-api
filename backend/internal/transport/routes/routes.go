package routes

import (
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/transport/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterRoutes sets up all routes for the application.
func RegisterRoutes(router *gin.Engine, postHandler *handlers.PostHandler) {
	// Health check
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Pong, API is up"})
	})

	api := router.Group("/api/v1")
	{
		RegisterPostRoutes(api, postHandler)
	}
	// Apply JSON middleware globally
	//router.Use(middleware.JSONMiddleware)
	// Post routes
	//router.HandleFunc("/posts", postHandler.CreatePost).Methods("POST") // Create a new post
	//router.HandleFunc("/posts", postHandler.GetAllPosts).Methods("GET") // Get all posts
}

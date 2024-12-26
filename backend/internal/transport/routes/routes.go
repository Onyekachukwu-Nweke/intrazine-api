package routes

import (
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/transport/handlers"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/transport/middleware"
	"github.com/gorilla/mux"
)

// RegisterRoutes sets up all routes for the application.
func RegisterRoutes(router *mux.Router, postHandler *handlers.PostHandler) {
	// Apply JSON middleware globally
	router.Use(middleware.JSONMiddleware)
	// Post routes
	router.HandleFunc("/posts", postHandler.CreatePost).Methods("POST") // Create a new post
	//router.HandleFunc("/posts", postHandler.GetAllPosts).Methods("GET") // Get all posts
}

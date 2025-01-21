package routes

import (
	"net/http"

	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/transport/handlers"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up all routes for the application.
func RegisterRoutes(router *gin.Engine, postHandler *handlers.PostHandler, authHandler *handlers.AuthHandler, commentHandler *handlers.CommentHandler) {
	// Health check
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Pong, API is up"})
	})

	api := router.Group("/api/v1")
	{
		RegisterPostRoutes(api, postHandler)
		RegisterCommentRoutes(api, commentHandler)
		RegisterAuthRoutes(api, authHandler)
	}
}

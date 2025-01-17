package routes

import (
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/transport/handlers"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/transport/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(router *gin.RouterGroup, postHandler *handlers.PostHandler) {
	// Group routes under `/posts`
	posts := router.Group("/posts")
	{
		posts.POST("/", middleware.JWTAuth(), postHandler.CreatePost)      // Create a new post
		posts.GET("/:id", postHandler.GetPostById)                         // Get a post by ID
		posts.GET("/", postHandler.GetAllPosts)                            // Get all posts
		posts.PATCH("/:id", middleware.JWTAuth(), postHandler.UpdatePost)  // Update a post
		posts.DELETE("/:id", middleware.JWTAuth(), postHandler.DeletePost) // Delete a post
	}
}

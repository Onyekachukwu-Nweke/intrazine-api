package routes

import (
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/transport/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(router *gin.RouterGroup, postHandler *handlers.PostHandler) {
	// Group routes under `/posts`
	posts := router.Group("/posts")
	{
		posts.POST("/", postHandler.CreatePost) // Create a new post
		//posts.GET("/:id", postHandler.GetPostByID)       // Get a post by ID
		//posts.GET("/", postHandler.GetAllPosts) // Get all posts
		//posts.PATCH("/:id", postHandler.UpdatePost)      // Update a post
		//posts.DELETE("/:id", postHandler.DeletePost)     // Delete a post
	}
}

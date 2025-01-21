package routes

import (
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/transport/handlers"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/transport/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(router *gin.RouterGroup, postHandler *handlers.PostHandler) {
	posts := router.Group("/posts")
	{
		posts.POST("/", middleware.JWTAuth(), postHandler.CreatePost)
		posts.GET("/", postHandler.GetAllPosts)
		posts.GET("/:id", postHandler.GetPostById)
		posts.PATCH("/:id", middleware.JWTAuth(), postHandler.UpdatePost)
		posts.DELETE("/:id", middleware.JWTAuth(), postHandler.DeletePost)
	}
}

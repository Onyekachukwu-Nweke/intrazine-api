package routes

import (
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/transport/handlers"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/transport/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterCommentRoutes(router *gin.RouterGroup, commentHandler *handlers.CommentHandler) {
	comments := router.Group("/posts/:id/comments")
	{
		comments.POST("/", middleware.JWTAuth(), commentHandler.CreateComment)
		comments.GET("/", commentHandler.GetCommentsByPost)
		comments.GET("/:commentId", commentHandler.GetComment)
		comments.PUT("/:commentId", middleware.JWTAuth(), commentHandler.UpdateComment)
		comments.DELETE("/:commentId", middleware.JWTAuth(), commentHandler.DeleteComment)
	}
}

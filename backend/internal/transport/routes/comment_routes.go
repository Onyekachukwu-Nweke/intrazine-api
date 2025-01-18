package routes

import (
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/transport/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterCommentRoutes(router *gin.RouterGroup, commentHandler *handlers.CommentHandler) {
	comments := router.Group("/posts/:id/comments")
	{
		comments.POST("/", commentHandler.CreateComment)
		comments.GET("/", commentHandler.GetCommentsByPost)
		comments.GET("/:commentId", commentHandler.GetComment)
		comments.PUT("/:comentId", commentHandler.UpdateComment)
		comments.DELETE("/:commentId", commentHandler.DeleteComment)
	}
}

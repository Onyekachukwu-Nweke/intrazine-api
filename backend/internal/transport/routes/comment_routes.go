package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/transport/handlers"
)

func RegisterCommentRoutes(r *gin.RouterGroup, h *handlers.CommentHandler) {
	comments := r.Group("/posts/:postId/comments")
	{
		comments.POST("", h.CreateComment)
		comments.GET("", h.GetCommentsByPost)
	}

	// Direct comment routes
	r.GET("/comments/:commentId", h.GetComment)
	r.PUT("/comments/:commentId", h.UpdateComment)
	r.DELETE("/comments/:commentId", h.DeleteComment)
}

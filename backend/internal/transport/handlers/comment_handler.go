package handlers

import (
	"net/http"

	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/interfaces"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/models"
	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	commentService interfaces.CommentService
}

func NewCommentHandler(commentService interfaces.CommentService) *CommentHandler {
	return &CommentHandler{
		commentService: commentService,
	}
}

func (h *CommentHandler) CreateComment(c *gin.Context) {
	postID := c.Param("id")
	userID := c.GetString("user_id") // From auth middleware

	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	comment.PostId = postID
	comment.UserId = userID

	if err := h.commentService.CreateComment(c.Request.Context(), &comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": comment})
}

func (h *CommentHandler) GetComment(c *gin.Context) {
	commentID := c.Param("commentId")

	comment, err := h.commentService.GetCommentByID(c.Request.Context(), commentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comment)
}

func (h *CommentHandler) GetCommentsByPost(c *gin.Context) {
	postID := c.Param("id")

	comments, err := h.commentService.GetCommentsByPostID(c.Request.Context(), postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (h *CommentHandler) UpdateComment(c *gin.Context) {
	commentID := c.Param("commentId")
	userID := c.GetString("user_id") // From auth middleware

	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	comment.Id = commentID

	if err := h.commentService.UpdateComment(c.Request.Context(), userID, &comment); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "unauthorized to perform this action" {
			status = http.StatusForbidden
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comment})
}

func (h *CommentHandler) DeleteComment(c *gin.Context) {
	commentID := c.Param("commentId")
	userID := c.GetString("user_id") // From auth middleware

	if err := h.commentService.DeleteComment(c.Request.Context(), userID, commentID); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "unauthorized to perform this action" {
			status = http.StatusForbidden
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
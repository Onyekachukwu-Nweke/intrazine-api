package interfaces

import (
	"context"

	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/models"
)

type CommentRepo interface {
	Create(ctx context.Context, comment *models.Comment) error
	GetByID(ctx context.Context, id string) (*models.Comment, error)
	GetByPostID(ctx context.Context, postID string) ([]models.Comment, error)
	Update(ctx context.Context, comment *models.Comment) error
	Delete(ctx context.Context, id string) error
}

type CommentService interface {
	CreateComment(ctx context.Context, comment *models.Comment) error
	GetCommentByID(ctx context.Context, id string) (*models.Comment, error)
	GetCommentsByPostID(ctx context.Context, postID string) ([]models.Comment, error)
	UpdateComment(ctx context.Context, userID string, comment *models.Comment) error
	DeleteComment(ctx context.Context, userID string, commentID string) error
}

package interfaces

import (
	"context"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/models"
)

type PostStore interface {
	//GetPostByID(context.Context, string) (models.Post, error)
	CreatePost(context.Context, models.Post) (models.Post, error)
	GetAllPosts(context.Context) ([]models.Post, error)
	//UpdatePost(context.Context, string, *models.Post) (models.Post, error)
	//DeletePost(context.Context, string) error
	//GetOwnerIDByPostID(ctx context.Context, postID string) (string, error)
}

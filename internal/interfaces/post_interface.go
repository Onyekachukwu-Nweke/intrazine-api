package interfaces

import (
	"context"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/models"
)

type PostRepo interface {
	CreatePost(context.Context, models.Post) (models.Post, error)
	GetAllPosts(context.Context) ([]models.Post, error)
	GetPostById(context.Context, string) (models.Post, error)
	UpdatePost(context.Context, string, models.Post) (models.Post, error)
	DeletePost(context.Context, string) error
}

type PostService interface {
	CreatePost(context.Context, models.Post) (models.Post, error)
	GetAllPosts(context.Context) ([]models.Post, error)
	GetPostById(context.Context, string) (models.Post, error)
	UpdatePost(context.Context, string, models.Post) (models.Post, error)
	DeletePost(context.Context, string, string) error
}

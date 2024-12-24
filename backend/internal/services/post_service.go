package services

import (
	"context"
	"fmt"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/interfaces"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/models"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/repositories"
)

type PostService struct {
	Repo interfaces.PostStore
}

func NewPostService(repo *repositories.PostRepository) interfaces.PostStore {
	return &PostService{Repo: repo}
}

func (s *PostService) CreatePost(ctx context.Context, post models.Post) (models.Post, error) {
	insertedPost, err := s.Repo.CreatePost(ctx, post)
	if err != nil {
		return models.Post{}, err
	}
	return insertedPost, nil
}

func (s *PostService) GetAllPosts(ctx context.Context) ([]models.Post, error) {
	fmt.Println("Retrieving All Posts")
	post, err := s.Repo.GetAllPosts(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting all posts: %w", err)
	}

	return post, nil
}

//func (s *PostService) ToggleLike(postID int) error {
//	return s.Repo.ToggleLike(postID)
//}

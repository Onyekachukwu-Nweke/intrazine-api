package services

import (
	"context"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/interfaces"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/models"
	//"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/repositories"
)

type PostService struct {
	Repo interfaces.PostRepo
}

func NewPostService(repo interfaces.PostRepo) *PostService {
	return &PostService{Repo: repo}
}

func (s *PostService) CreatePost(ctx context.Context, post models.Post) (models.Post, error) {
	insertedPost, err := s.Repo.CreatePost(ctx, post)
	if err != nil {
		return models.Post{}, err
	}
	return insertedPost, nil
}

//func (s *PostService) GetAllPosts(ctx context.Context) ([]models.Post, error) {
//	fmt.Println("Retrieving All Posts")
//	post, err := s.Repo.GetAllPosts(ctx)
//	if err != nil {
//		return nil, fmt.Errorf("error getting all posts: %w", err)
//	}
//
//	return post, nil
//}

//func (s *PostService) GetAllPosts(ctx context.Context) ([]Post, error) {
//	fmt.Println("Retrieving All Posts")
//	post, err := s.PostStore.GetAllPosts(ctx)
//	if err != nil {
//		return nil, fmt.Errorf("error getting all posts: %w", err)
//	}
//
//	return post, nil
//}
//
///**
// *
// *
// */
//func (s *PostService) UpdatePost(ctx context.Context, id string, updatedPost Post) (Post, error) {
//	pst, err := s.PostStore.UpdatePost(ctx, id, updatedPost)
//	if err != nil {
//		fmt.Println("error updating post")
//		return Post{}, err
//	}
//
//	return pst, nil
//}
//
///**
// *
// *
// */
//func (s *PostService) DeletePost(ctx context.Context, id string) error {
//	return s.PostStore.DeletePost(ctx, id)
//}

//func (s *PostService) ToggleLike(postID int) error {
//	return s.Repo.ToggleLike(postID)
//}

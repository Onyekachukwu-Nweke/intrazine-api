package post

import (
	"context"
	"fmt"

	// "fmt"
	"time"
)

type Post struct {
	ID string
	User_id string
	Title string
	Content string
	Created_at time.Time
	Updated_at time.Time
}

type PostStore interface {
	GetPost(context.Context, string) (Post, error)
	CreatePost(context.Context, Post) (Post, error)
}

type PostService struct {
	PostStore PostStore
}

func NewPostService(store PostStore) *PostService {
	return &PostService{
			PostStore: store,
	}
}

func (s *PostService) CreatePost(ctx context.Context, post Post) (Post, error) {
	insertedPost, err := s.PostStore.CreatePost(ctx, post)
	if err != nil {
		return Post{}, err
	}
	return insertedPost, nil
}

func (s *PostService) GetPost(ctx context.Context, id string) (Post, error) {
	fmt.Println("Retrieving a post")
	post, err := s.PostStore.GetPost(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Post{}, err
	}

	return post, nil
}
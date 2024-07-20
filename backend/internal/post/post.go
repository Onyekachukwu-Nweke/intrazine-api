package post

import (
	"context"
	"fmt"

	// "fmt"
	"time"
)

type Post struct {
	ID string `json:"id"`
	User_id string `json:"user_id" validate:"required"`
	Title string  `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type PostStore interface {
	GetPostByID(context.Context, string) (Post, error)
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

/******************** CreatePost  ***********************/
func (s *PostService) CreatePost(ctx context.Context, post Post) (Post, error) {
	insertedPost, err := s.PostStore.CreatePost(ctx, post)
	if err != nil {
		return Post{}, err
	}
	return insertedPost, nil
}

/*
 * func: GetPostByID
 * Interacts with servcie layer to get a post by id from repository layer
 */
func (s *PostService) GetPostByID(ctx context.Context, id string) (Post, error) {
	fmt.Println("Retrieving a post")
	post, err := s.PostStore.GetPostByID(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Post{}, err
	}

	return post, nil
}
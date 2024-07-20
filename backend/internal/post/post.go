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
	GetAllPosts(context.Context) ([]Post, error)
	UpdatePost(context.Context, string, Post) (Post, error)
	DeletePost(context.Context, string) error
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

/**
 * GetPostByID
 * * Interacts with servcie layer to get a post by id from repository layer
 * @params: ctx, id
 */
func (s *PostService) GetPostByID(ctx context.Context, id string) (Post, error) {
	fmt.Printf("Retrieving a post with ID: %s", id)
	post, err := s.PostStore.GetPostByID(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Post{}, err
	}

	return post, nil
}

/**
 * GetAllPosts
 * * Gets all posts from the repository layer and returns it
 * @param: ctx
 */
func (s *PostService) GetAllPosts(ctx context.Context) ([]Post, error) {
	fmt.Println("Retrieving All Posts")
	post, err := s.PostStore.GetAllPosts(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting all posts: %w", err)
	}

	return post, nil
}

/**
 *
 *
 */
func (s *PostService) UpdatePost(ctx context.Context, id string, updatedPost Post) (Post, error) {
	pst, err := s.PostStore.UpdatePost(ctx, id, updatedPost)
	if err != nil {
		fmt.Println("error updating post")
		return Post{}, err
	}

	return pst, nil
}

/**
 *
 *
 */
func (s *PostService) DeletePost(ctx context.Context, id string) error {
	return s.PostStore.DeletePost(ctx, id)
}
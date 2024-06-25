package post

import (
	"context"
	"fmt"
	// "fmt"
	"time"

	// "github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/common"
)

type Post struct {
	ID string
	Title string
	Content string
	Created_at time.Time
	Updated_at time.Time
}

type PostStore interface {
	GetPost(context.Context, string) (Post, error)
}

type PostService struct {
	PostStore PostStore
}

func NewPostService(store PostStore) *PostService {
	return &PostService{
			PostStore: store,
	}
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
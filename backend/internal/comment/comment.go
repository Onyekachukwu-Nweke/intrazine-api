package comment

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var (
	ErrFetchingComment = errors.New("[Error] Failed to fetch comment by id")
	ErrNotImplemented = errors.New("[Warning] Not implemented")
)

// Comment - a representation of the comment
// structure for out service
type Comment struct {
	ID string						`json:"id"`
	PostID string				`json:"post_id" validate:"required"`
	UserID string				`json:"user_id" validate:"required"`
	Content string			`json:"content" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Store - this interface defines all the methods
// that our service needs in order to operate 
type CommentStore interface {
	GetComment(context.Context, string) (Comment, error)
	PostComment(context.Context, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
	UpdateComment(context.Context, string, Comment) (Comment, error)
}

// Service - is the struct on which all our
// logic will be built on top of
type CommentService struct {
	CommentStore CommentStore
}

// NewService - returns a pointer to a new
//  service
func NewCommentService(store CommentStore) *CommentService {
	return &CommentService{
		CommentStore: store,
	}
}

func (s *CommentService) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Retrieving a comment")
	cmt, err := s.CommentStore.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}

	return cmt, nil
}

func (s *CommentService) UpdateComment(
	ctx context.Context, 
	ID string, 
	updatedCmt Comment,
) (Comment, error) {
	cmt, err := s.CommentStore.UpdateComment(ctx, ID, updatedCmt)
	if err != nil {
		fmt.Println("error updating commenting")
		return Comment{}, err
	}
	return cmt, nil
}

func (s *CommentService) DeleteComment(ctx context.Context, id string) error {
	return s.CommentStore.DeleteComment(ctx, id)
}

func (s *CommentService) PostComment(ctx context.Context, cmt Comment) (Comment, error) {
	insertedCmt, err := s.CommentStore.PostComment(ctx, cmt)
	if err != nil {
		return Comment{}, err
	}
	return insertedCmt, nil
}
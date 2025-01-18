package services

import (
	"context"
	"errors"
	"strings"

	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/interfaces"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/models"
	uuid "github.com/satori/go.uuid"
)

var (
	ErrCommentNotFound     = errors.New("comment not found")
	ErrUnauthorized        = errors.New("unauthorized to perform this action")
	ErrInvalidCommentInput = errors.New("invalid comment input")
)

type CommentService struct {
	commentRepo interfaces.CommentRepo
	postService interfaces.PostService
}

func NewCommentService(
	commentRepo interfaces.CommentRepo,
	postService interfaces.PostService,
) *CommentService {
	return &CommentService{
		commentRepo: commentRepo,
		postService: postService,
	}
}

func (s *CommentService) validateComment(ctx context.Context, comment *models.Comment) error {
	// Trim spaces from content
	comment.Content = strings.TrimSpace(comment.Content)

	if len(comment.Content) < 1 {
		return ErrInvalidCommentInput
	}

	if len(comment.Content) > 1000 { // Maximum comment length
		return ErrInvalidCommentInput
	}

	// Verify post exists
	_, err := s.postService.GetPostById(ctx, comment.PostId)
	if err != nil {
		return errors.New("post not found")
	}

	return nil
}

func (s *CommentService) CreateComment(ctx context.Context, comment *models.Comment) error {
	if err := s.validateComment(ctx, comment); err != nil {
		return err
	}

	// Generate new UUID for the comment
	comment.Id = uuid.NewV4().String()

	return s.commentRepo.Create(ctx, comment)
}

func (s *CommentService) GetCommentByID(ctx context.Context, id string) (*models.Comment, error) {
	comment, err := s.commentRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if comment == nil {
		return nil, ErrCommentNotFound
	}
	return comment, nil
}

func (s *CommentService) GetCommentsByPostID(ctx context.Context, postID string) ([]models.Comment, error) {
	return s.commentRepo.GetByPostID(ctx, postID)
}

func (s *CommentService) UpdateComment(ctx context.Context, userID string, comment *models.Comment) error {
	existing, err := s.GetCommentByID(ctx, comment.Id)
	if err != nil {
		return err
	}

	// Authorization check
	if existing.UserId != userID {
		return ErrUnauthorized
	}

	if err := s.validateComment(ctx, comment); err != nil {
		return err
	}

	// Preserve original metadata
	comment.UserId = existing.UserId
	comment.PostId = existing.PostId
	comment.CreatedAt = existing.CreatedAt

	return s.commentRepo.Update(ctx, comment)
}

func (s *CommentService) DeleteComment(ctx context.Context, userID string, commentID string) error {
	existing, err := s.GetCommentByID(ctx, commentID)
	if err != nil {
		return err
	}

	// Authorization check
	if existing.UserId != userID {
		return ErrUnauthorized
	}

	return s.commentRepo.Delete(ctx, commentID)
}

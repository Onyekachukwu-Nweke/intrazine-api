package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/interfaces"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type CommentRepository struct {
	DB *sqlx.DB
}

func NewCommentRepository(db *sqlx.DB) interfaces.CommentRepo {
	return &CommentRepository{DB: db}
}

func (r *CommentRepository) Create(ctx context.Context, comment *models.Comment) error {
	query := `
        INSERT INTO comments (id, post_id, user_id, content, created_at, updated_at)
        VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
    `
	_, err := r.DB.ExecContext(ctx, query, comment.Id, comment.PostId, comment.UserId, comment.Content)
	return err
}

func (r *CommentRepository) GetByID(ctx context.Context, id string) (*models.Comment, error) {
	comment := &models.Comment{}
	query := `SELECT * FROM comments WHERE id = $1`
	fmt.Println(id)
	err := r.DB.QueryRowxContext(ctx, query, id).StructScan(comment)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return comment, err
}

func (r *CommentRepository) GetByPostID(ctx context.Context, postID string) ([]models.Comment, error) {
	query := `
        SELECT * FROM comments 
        WHERE post_id = $1 
        ORDER BY created_at DESC
    `
	rows, err := r.DB.QueryxContext(ctx, query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		err := rows.StructScan(&comment)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func (r *CommentRepository) Update(ctx context.Context, comment *models.Comment) error {
	query := `
        UPDATE comments 
        SET content = $1, updated_at = CURRENT_TIMESTAMP
        WHERE id = $2
    `
	_, err := r.DB.ExecContext(ctx, query, comment.Content, comment.Id)
	return err
}

func (r *CommentRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM comments WHERE id = $1`
	_, err := r.DB.ExecContext(ctx, query, id)
	return err
}

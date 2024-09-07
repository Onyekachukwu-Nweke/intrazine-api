package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/comment"
	uuid "github.com/satori/go.uuid"
)

type CommentRow struct {
	ID        string           `db:"id"`
	PostID    sql.NullString   `db:"post_id"`
	UserID    sql.NullString   `db:"user_id"`
	Content   sql.NullString   `db:"content"`
	CreatedAt sql.NullTime     `db:"created_at"`
	UpdatedAt sql.NullTime     `db:"updated_at"`
}

func convertCommentRowToComment(c CommentRow) comment.Comment {
	return comment.Comment{
		ID: c.ID,
		PostID: c.PostID.String,
		UserID: c.UserID.String,
		Content: c.Content.String,
		CreatedAt: c.CreatedAt.Time,
		UpdatedAt: c.UpdatedAt.Time,
	}
}

func (d *Database) GetComment(
	ctx context.Context, 
	uuid string,
	) (comment.Comment, error) {
		var cmtRow CommentRow
		// _, err := d.Client.ExecContext(ctx, "SELECT pg_sleep(16)")
		// if err != nil {
		// 	return comment.Comment{}, err
		// }

		row := d.Client.QueryRowContext(
			ctx,
			`SELECT id, post_id, user_id, content, created_at, updated_at 
			FROM comments
			WHERE id = $1`,
			uuid,
		)
		err := row.Scan(&cmtRow.ID, &cmtRow.PostID, &cmtRow.UserID, &cmtRow.Content, &cmtRow.CreatedAt, &cmtRow.UpdatedAt)
		if err != nil {
			return comment.Comment{}, fmt.Errorf("error fetching the comment by uuid: %w", err)
		}

		return convertCommentRowToComment(cmtRow), nil
}

func (d *Database) GetOwnerIDByCommentID(ctx context.Context, commentID string) (string, error) {
	var userID string
	err := d.Client.QueryRowContext(ctx, "SELECT user_id FROM comments WHERE id = $1", commentID).Scan(&userID)
	return userID, err
}

func (d *Database) PostComment(ctx context.Context, cmt comment.Comment) (comment.Comment, error) {
	cmt.ID = uuid.NewV4().String()
	cmt.CreatedAt = time.Now()
	cmt.UpdatedAt = cmt.CreatedAt
	postRow := CommentRow{
		ID: cmt.ID,
		UserID: sql.NullString{String: cmt.UserID, Valid: true},
		PostID: sql.NullString{String: cmt.PostID, Valid: true},
		Content: sql.NullString{String: cmt.Content, Valid: true},
		CreatedAt: sql.NullTime{Time: cmt.CreatedAt, Valid: true},
		UpdatedAt: sql.NullTime{Time: cmt.UpdatedAt, Valid: true},
	}
	rows, err := d.Client.NamedQueryContext(
		ctx,
		`INSERT INTO comments
		(id, post_id, user_id, content, created_at, updated_at) 
		VALUES
		(:id, :post_id, :user_id, :content, :created_at, :updated_at)`,
		postRow,
	)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("failed to insert comment: %w", err)
	}
	if err := rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return cmt, nil
}

func (d *Database) DeleteComment(ctx context.Context, id string) error {
	_, err := d.Client.ExecContext(
		ctx,
		`DELETE FROM comments where id = $1`,
		id,
	)
	if err != nil {
		return fmt.Errorf("failed to delete comment from database: %w", err)
	}
	return nil
}

func (d *Database) UpdateComment(
	ctx context.Context, 
	id string, 
	cmt comment.Comment,
) (comment.Comment, error) {
	cmt.UpdatedAt = time.Now()
	cmtRow := CommentRow{
		ID: id,
		PostID: sql.NullString{String: cmt.PostID, Valid: true},
		UserID: sql.NullString{String: cmt.UserID, Valid: true},
		Content: sql.NullString{String: cmt.Content, Valid: true},
		CreatedAt: sql.NullTime{Time: cmt.CreatedAt, Valid: true},
		UpdatedAt: sql.NullTime{Time: cmt.UpdatedAt, Valid: true},
	}

	rows, err := d.Client.NamedQueryContext(
		ctx,
		`UPDATE comments SET
		content = :content,
		updated_at = :updated_at
		WHERE id = :id`,
		cmtRow,
	)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("failed to update comment: %w", err)
	}
	if err := rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return convertCommentRowToComment(cmtRow), nil
}
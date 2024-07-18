package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/post"
	// "github.com/google/uuid"
	uuid "github.com/satori/go.uuid"
)



type PostRow struct {
	ID string
	User_id sql.NullString
	Title sql.NullString
	Content sql.NullString
	Created_at sql.NullTime
	Updated_at sql.NullTime
}

func convertPostRowToPost(p PostRow) post.Post {
	return post.Post{
		ID: p.ID,
		Title: p.Title.String,
		User_id: p.User_id.String, 
		Content: p.Content.String,
		Created_at: p.Created_at.Time,
		Updated_at: p.Updated_at.Time,
	}
}

func (d *Database) CreatePost(ctx context.Context, pst post.Post) (post.Post, error) {
	pst.ID = uuid.NewV4().String()
	pst.Created_at = time.Now()
	pst.Updated_at = pst.Created_at
	pstRow := PostRow {
		ID: pst.ID,
		User_id: sql.NullString{String: pst.User_id, Valid: true},
		Title: sql.NullString{String: pst.Title, Valid: true},
		Content: sql.NullString{String: pst.Content, Valid: true},
		Created_at: sql.NullTime{Time: pst.Created_at, Valid: true},
		Updated_at: sql.NullTime{Time: pst.Updated_at, Valid: true},
	}
	rows, err := d.Client.NamedQueryContext(
		ctx,
		`INSERT INTO Posts (id, user_id, title, content, created_at, updated_at) VALUES (:id, :user_id, :title, :content, :created_at, :updated_at)`, pstRow,
	)
	if err != nil {
		return post.Post{}, fmt.Errorf("failed to insert comment: %w", err)
	}
	if err := rows.Close(); err != nil {
		return post.Post{}, fmt.Errorf("failed to close row:  %w", err)
	}

	return pst, nil
}

func (d *Database) GetPost(
	ctx context.Context,
	uuid string,
) (post.Post, error) {
	var postRow PostRow

	row := d.Client.QueryRowContext(
		ctx, 
		`SELECT id, title
		FROM post
		WHERE id = $1`,
		uuid)
	
		err := row.Scan(&postRow.ID, &postRow.Title, &postRow.Content, &postRow.Created_at, &postRow.Updated_at)

		if err != nil {
			return post.Post{}, fmt.Errorf("error fetching post from id: %w", err)
		}

		return convertPostRowToPost(postRow), nil
}
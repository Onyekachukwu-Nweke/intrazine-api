package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/post"
	// uuid "github.com/satori/go.uuid"
)



type PostRow struct {
	ID string
	Title sql.NullString
	Content sql.NullString
	Created_at sql.NullTime
	Updated_at sql.NullTime
}

func convertPostRowToPost(p PostRow) post.Post {
	return post.Post{
		ID: p.ID,
		Title: p.Title.String,
		Content: p.Content.String,
		Created_at: p.Created_at.Time,
		Updated_at: p.Updated_at.Time,
	}
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
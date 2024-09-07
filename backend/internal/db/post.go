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
		`INSERT INTO Posts (id, user_id, title, content, created_at, updated_at) VALUES (:id, :user_id, :title, :content, :created_at, :updated_at)`, pstRow)
	if err != nil {
		return post.Post{}, fmt.Errorf("failed to insert post: %w", err)
	}
	if err := rows.Close(); err != nil {
		return post.Post{}, fmt.Errorf("failed to close row:  %w", err)
	}

	return pst, nil
}

func (d *Database) GetOwnerIDByPostID(ctx context.Context, postID string) (string, error) {
	var userID string
	err := d.Client.QueryRowContext(ctx, "SELECT user_id FROM posts WHERE id = $1", postID).Scan(&userID)
	return userID, err
}

func (d *Database) GetPostByID(
	ctx context.Context,
	uuid string,
) (post.Post, error) {
	var postRow PostRow

	row := d.Client.QueryRowContext(
		ctx, 
		`SELECT id, user_id, title, content, created_at, updated_at
		FROM Posts
		WHERE id = $1`,
		uuid)
	
		err := row.Scan(&postRow.ID, &postRow.User_id,&postRow.Title, &postRow.Content, &postRow.Created_at, &postRow.Updated_at)

		if err != nil {
			if err == sql.ErrNoRows {
				return post.Post{}, fmt.Errorf("no post found with ID: %w", err) // Case II and III distinction
		}
			return post.Post{}, fmt.Errorf("error fetching post from id: %w", err)
		}

		return convertPostRowToPost(postRow), nil
}

func (d *Database) GetAllPosts(ctx context.Context) ([]post.Post, error) {
	rows, err := d.Client.QueryContext(
		ctx,
		`SELECT * from posts`,
	)
	if err != nil {
		return nil, fmt.Errorf("error querying posts: %w", err)
	}

	defer rows.Close()

	var posts []post.Post
	for rows.Next() {
		var p PostRow
		err := rows.Scan(&p.ID, &p.User_id, &p.Title, &p.Content, &p.Created_at, &p.Updated_at)
		if err != nil {
				return nil, fmt.Errorf("error scanning post: %w", err)
		}
		posts = append(posts, convertPostRowToPost(p))
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading post rows: %w", err)
	}

	return posts, nil
}

func (d *Database) UpdatePost(ctx context.Context, id string, pst post.Post) (post.Post, error) {
	pst.Updated_at = time.Now()
	postRow := PostRow{
		ID: id,
		User_id: sql.NullString{String: pst.User_id, Valid: true},
		Title: sql.NullString{String: pst.Title, Valid: true},
		Content: sql.NullString{String: pst.Content, Valid: true},
		Created_at: sql.NullTime{Time: pst.Created_at, Valid: true},
		Updated_at: sql.NullTime{Time: pst.Updated_at, Valid: true},
	}

	rows, err := d.Client.NamedQueryContext(
		ctx,
		`UPDATE posts SET 
		title = :title,
		content = :content,
		updated_at = :updated_at
		WHERE id = :id`,
		postRow,
	)
	if err != nil {
		return post.Post{}, fmt.Errorf("failed to update post: %w", err)
	}

	if err := rows.Close(); err != nil {
		return post.Post{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return convertPostRowToPost(postRow), nil
}

func (d *Database) DeletePost(ctx context.Context, id string) error {
	_, err := d.Client.ExecContext(
		ctx,
		`DELETE FROM posts where id = $1`,
		id,
	)
	if err != nil {
		return fmt.Errorf("failed to delete post from database: %w", err)
	}
	return nil
}
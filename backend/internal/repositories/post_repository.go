package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/interfaces"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/models"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	"time"
)

type PostRepository struct {
	DB *sqlx.DB
}

func NewPostRepository(db *sqlx.DB) interfaces.PostRepo {
	return &PostRepository{DB: db}
}

func (r *PostRepository) CreatePost(ctx context.Context, post models.Post) (models.Post, error) {
	post.ID = uuid.NewV4().String()
	post.CreatedAt = time.Now().Format(time.DateTime)
	post.UpdatedAt = post.CreatedAt
	rows, err := r.DB.NamedQueryContext(
		ctx,
		`INSERT INTO posts (id, user_id, title, content, created_at, updated_at) VALUES (:id, :user_id, :title, :content, :created_at, :updated_at)`, post)
	if err != nil {
		return models.Post{}, fmt.Errorf("failed to insert post: %w", err)
	}
	if err := rows.Close(); err != nil {
		return models.Post{}, fmt.Errorf("failed to close row:  %w", err)
	}

	return post, nil
}

func (r *PostRepository) GetPostById(ctx context.Context, uuid string) (models.Post, error) {
	row := r.DB.QueryRowContext(
		ctx,
		`SELECT id, user_id, title, content, created_at, updated_at
			FROM Posts
			WHERE id = $1`,
		uuid)

	post := models.Post{}
	err := row.Scan(&post.ID, &post.UserId, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return post, fmt.Errorf("no post found with ID: %w", err) // Case II and III distinction
		}
		return post, fmt.Errorf("error fetching post from ID: %w", err)
	}
	return post, nil
}

func (r *PostRepository) GetAllPosts(ctx context.Context) ([]models.Post, error) {
	rows, err := r.DB.QueryContext(
		ctx,
		`SELECT * from posts`,
	)
	if err != nil {
		return nil, fmt.Errorf("error querying posts: %w", err)
	}

	defer rows.Close()
	var posts []models.Post
	for rows.Next() {
		post := models.Post{}
		var coverPhoto sql.NullString
		err := rows.Scan(&post.ID, &post.UserId, &post.Title, &post.Content, &coverPhoto, &post.Likes, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning post: %w", err)
		}
		post.CoverPhoto = coverPhoto.String // Will be empty string if NULL
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading post rows: %w", err)
	}

	return posts, nil
}

func (r *PostRepository) UpdatePost(ctx context.Context, post models.Post) (models.Post, error) {
	post.UpdatedAt = time.Now().Format(time.DateTime)
	rows, err := r.DB.NamedQueryContext(
		ctx,
		`UPDATE posts SET
				title = :title,
				content = :content,
				updated_at = :updated_at
				WHERE id = :id`,
		post,
	)
	if err != nil {
		return models.Post{}, fmt.Errorf("error updating post: %w", err)
	}

	if err := rows.Close(); err != nil {
		return models.Post{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return post, nil
}

func (r *PostRepository) DeletePost(ctx context.Context, id string) error {
	_, err := r.DB.ExecContext(
		ctx,
		`DELETE FROM posts where id = $1`,
		id,
	)
	if err != nil {
		return fmt.Errorf("failed to delete post from database: %w", err)
	}
	return nil
}

//func (r *PostRepository) ToggleLike(postID int) error {
//	query := "UPDATE posts SET likes = likes + 1 WHERE id = $1"
//	_, err := r.DB.Exec(query, postID)
//	return err
//}

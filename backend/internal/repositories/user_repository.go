package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/models"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	"time"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	query := `INSERT INTO users (id, username, email, password_hash, created_at, updated_at) VALUES (:id, :username, :email, :password_hash, :created_at, :updated_at)`
	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now().Format(time.DateTime)
	user.UpdatedAt = user.CreatedAt
	rows, err := r.DB.NamedQueryContext(ctx, query, user)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to create user: %w", err)
	}
	if err := rows.Close(); err != nil {
		return models.User{}, fmt.Errorf("failed to close row:  %w", err)
	}

	return user, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	query := `SELECT id, username, email, password_hash, created_at, updated_at FROM users WHERE email = $1`
	row := r.DB.QueryRowContext(ctx, query, email)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, userID string) (models.User, error) {
	query := `SELECT id, username, email, created_at, updated_at FROM users WHERE id = $1`
	row := r.DB.QueryRowContext(ctx, query, userID)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user models.User) (models.User, error) {
	query := `UPDATE users SET username = :username, email = :email, updated_at = NOW() WHERE id = :id`
	rows, err := r.DB.NamedQueryContext(ctx, query, user)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to update user: %w", err)
	}

	if err := rows.Close(); err != nil {
		return models.User{}, fmt.Errorf("failed to close rows: %w", err)
	}
	return user, nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, userID string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.DB.ExecContext(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("failed to delete user from database: %w", err)
	}
	return err
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	var user models.User

	row := r.DB.QueryRowContext(
		ctx,
		`SELECT id, username, email, password_hash FROM users WHERE username = $1`,
		username)

	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, fmt.Errorf("no user found with username: %w", err)
		}
		return models.User{}, fmt.Errorf("error fetching user from username: %w", err)
	}

	return user, nil
}

func (r *UserRepository) CheckUserExists(ctx context.Context, username, email string) (exists bool, field string, err error) {
	var userCount int
	err = r.DB.QueryRowContext(ctx, `SELECT COUNT(*) FROM users WHERE username = $1 OR email = $2`, username, email).Scan(&userCount)
	if err != nil {
		return false, "", err
	}
	if userCount > 0 {
		// Check which field is duplicated
		if err := r.DB.QueryRowContext(ctx, `SELECT COUNT(*) FROM users WHERE username = $1`, username).Scan(&userCount); err == nil && userCount > 0 {
			return true, "username", nil
		}
		return true, "email", nil
	}
	return false, "", nil
}

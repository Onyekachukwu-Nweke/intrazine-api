package repositories

import (
	"context"
	"fmt"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/interfaces"
	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/models"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	"time"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) interfaces.UserRepo {
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

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT id, username, email, password_hash, created_at, updated_at FROM users WHERE email = $1`
	row := r.DB.QueryRowContext(ctx, query, email)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, userID string) (*models.User, error) {
	query := `SELECT id, username, email, created_at, updated_at FROM users WHERE id = $1`
	row := r.DB.QueryRowContext(ctx, query, userID)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) UpdateUser(user models.User) error {
	query := `UPDATE users SET username = $1, email = $2, updated_at = NOW() WHERE id = $3`
	_, err := r.DB.Exec(query, user.Username, user.Email, user.ID)
	return err
}

func (r *UserRepository) DeleteUser(userID string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.DB.Exec(query, userID)
	return err
}

//func (d *Database) GetUserByUsername(ctx context.Context, username string) (user.User, error) {
//	var usrRow UserRow
//
//	row := d.Client.QueryRowContext(
//		ctx,
//		`SELECT id, username, email, password_hash FROM users WHERE username = $1`,
//		username)
//
//	err := row.Scan(&usrRow.ID, &usrRow.Username, &usrRow.Email, &usrRow.PasswordHash)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			return user.User{}, fmt.Errorf("no user found with username: %w", err)
//		}
//		return user.User{}, fmt.Errorf("error fetching user from username: %w", err)
//	}
//
//	return convertUserRowToUser(usrRow), nil
//}

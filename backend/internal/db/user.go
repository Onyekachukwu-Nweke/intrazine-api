package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/user"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRow struct {
	ID string
	Username sql.NullString
	Email sql.NullString
	PasswordHash sql.NullString
	Created_at sql.NullTime
	Updated_at sql.NullTime
}

func hashPassword(password string) (string, error) {
	hash_pass, err := bcrypt.GenerateFromPassword([]byte(password), 112)
	return string(hash_pass), err
}

func (d *Database) CreateUser(ctx context.Context, usr user.User) (user.User, error) {
	usr.ID = uuid.NewV4().String()
	usr.Created_at = time.Now()
	usr.Updated_at = usr.Created_at
	usrRow := UserRow {
		ID: usr.ID,
		Username: sql.NullString{String: usr.Username, Valid: true},
		Email: sql.NullString{String: usr.Email, Valid: true},
		PasswordHash: sql.NullString{String: usr.PasswordHash, Valid: true},
		Created_at: sql.NullTime{Time: usr.Created_at, Valid: true},
		Updated_at: sql.NullTime{Time: usr.Updated_at, Valid: true},
	}
	rows, err := d.Client.NamedQueryContext(
		ctx,
	  `INSERT INTO users (id, username, email, password_hash, created_at, updated_at) VALUES`, usrRow)
	if err != nil {
		return user.User{}, fmt.Errorf("failed to create user: %w", err)
	}
}
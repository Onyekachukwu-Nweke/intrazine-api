package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Onyekachukwu-Nweke/piko-blog/backend/internal/user"
	uuid "github.com/satori/go.uuid"
)

type UserRow struct {
	ID string											`db:"id"`
	Username sql.NullString				`db:"username"`
	Email sql.NullString					`db:"email"`
	PasswordHash sql.NullString		`db:"password_hash"`
	Created_at sql.NullTime				`db:"created_at"`
	Updated_at sql.NullTime				`db:"updated_at"`
}

func convertUserRowToUser(u UserRow) user.User {
	return user.User{
		ID: u.ID,
		Username: u.Username.String,
		Email: u.Email.String,
		PasswordHash: u.PasswordHash.String,
	}
}

func (d *Database) CheckUserExists(ctx context.Context, username, email string) (exists bool, field string, err error) {
	var userCount int
	err = d.Client.QueryRowContext(ctx, `SELECT COUNT(*) FROM users WHERE username = $1 OR email = $2`, username, email).Scan(&userCount)
	if err != nil {
		return false, "", err
	}
	if userCount > 0 {
		// Check which field is duplicated
		if err := d.Client.QueryRowContext(ctx, `SELECT COUNT(*) FROM users WHERE username = $1`, username).Scan(&userCount); err == nil && userCount > 0 {
				return true, "username", nil
		}
		return true, "email", nil
  }
  return false, "", nil
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
	  `INSERT INTO users (id, username, email, password_hash, created_at, updated_at) VALUES (:id, :username, :email, :password_hash, :created_at, :updated_at)`, usrRow)
	if err != nil {
		return user.User{}, fmt.Errorf("failed to create user: %w", err)
	}
	if err := rows.Close(); err != nil {
		return user.User{}, fmt.Errorf("failed to close row:  %w", err)
	}

	return usr, nil
}

func (d *Database) GetUserByUsername(ctx context.Context, username string) (user.User, error) {
	var usrRow  UserRow

	row := d.Client.QueryRowContext(
		ctx,
		`SELECT id, username, email, password_hash FROM users WHERE username = $1`,
		username)
	
	err := row.Scan(&usrRow.ID, &usrRow.Username, &usrRow.Email, &usrRow.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
				return user.User{}, fmt.Errorf("no user found with username: %w", err)
		}
		return user.User{}, fmt.Errorf("error fetching user from username: %w", err)
	}

	return convertUserRowToUser(usrRow), nil
}
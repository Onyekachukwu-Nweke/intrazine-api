package models

type User struct {
	ID           string `json:"id" db:"id"`
	Username     string `json:"username" db:"username" validate:"required"`
	Email        string `json:"email" db:"email" validate:"required"`
	Photo        string `json:"photo" db:"photo"`
	PasswordHash string `json:"-" db:"password_hash"`
	CreatedAt    string `json:"created_at" db:"created_at"`
	UpdatedAt    string `json:"updated_at" db:"updated_at"`
}

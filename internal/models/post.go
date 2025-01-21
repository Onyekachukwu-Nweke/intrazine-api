package models

type Post struct {
	ID         string `json:"id" db:"id"`
	Title      string `json:"title" db:"title" validate:"required"`
	Content    string `json:"content" db:"content" validate:"required"`
	CoverPhoto string `json:"cover_photo" db:"cover_photo"`
	Likes      int    `json:"likes" db:"likes"`
	UserId     string `json:"user_id" db:"user_id" validate:"required"`
	CreatedAt  string `json:"created_at" db:"created_at"`
	UpdatedAt  string `json:"updated_at" db:"updated_at"`
}

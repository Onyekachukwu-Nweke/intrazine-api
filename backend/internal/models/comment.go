package models

type Comment struct {
	Id        string `json:"id" db:"id"`
	Content   string `json:"content" db:"content"`
	PostId    string `json:"post_id" db:"post_id"`
	UserId    string `json:"user_id" db:"user_id"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}


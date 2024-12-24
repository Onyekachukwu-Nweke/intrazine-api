package models

type Post struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CoverPhoto string `json:"cover_photo"`
	Likes      int    `json:"likes"`
	UserId     string `json:"user_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

package models

type Comment struct {
	Id        string `json:"id"`
	Content   string `json:"content"`
	PostId    string `json:"post_id"`
	UserId    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

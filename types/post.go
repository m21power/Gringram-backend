package types

type PostPayload struct {
	UserID  int    `json:"user_id" db:"user_id"`
	Content string `json:"content" db:"content"`
}
type UpdatePostPayload struct {
	Content string `json:"content" db:"content"`
}
type PostImagePayload struct {
	PostID int `json:"post_id" db:"post_id"`
	// ImageURL string `json:"image_url" db:"image_url"`
}

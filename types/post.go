package types

type PostPayload struct {
	UserID  int    `json:"user_id" db:"user_id"`
	Content string `json:"content" db:"content"`
}
type UpdatePostPayload struct {
	Content string `json:"content" db:"content"`
}

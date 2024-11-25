package domain

import "time"

type Interaction struct {
	ID     int       `json:"id" db:"id"`
	UserID int       `json:"user_id" db:"user_id"`
	PostID int       `json:"post_id" db:"post_id"`
	Seen   bool      `json:"seen" db:"seen"`
	SeenAt time.Time `json:"seen_at" db:"seen_at"`
}

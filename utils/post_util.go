package utils

import "time"

type PostResponse struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	Content   string    `json:"content" db:"content"`
	Images    string    `json:"images"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

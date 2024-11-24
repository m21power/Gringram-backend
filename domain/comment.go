package domain

import (
	"database/sql"
	"time"
)

type Comment struct {
	ID        int           `json:"id" db:"id"`
	Text      string        `json:"text" db:"text"`
	UserID    int           `json:"user_id" db:"user_id"`
	PostID    int           `json:"post_id" db:"post_id"`
	ParentID  sql.NullInt32 `json:"parent_id" db:"parent_id"`
	CreatedAt time.Time     `json:"created_at" db:"created_at"`
}

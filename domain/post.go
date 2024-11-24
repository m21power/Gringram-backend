package domain

import (
	"context"
	"database/sql"
	"time"
)

type Post struct {
	ID             int       `json:"id" db:"id"`
	Content        string    `json:"content" db:"content"`
	UserID         int       `json:"user_id" db:"user_id"`
	Status         string    `json:"status" db:"status"`
	Image_url      string    `json:"image_url" db:"image_url"`
	Likes_count    int       `json:"likes_count" db:"likes_count"`
	Comments_count int       `json:"comments_count" db:"comments_count"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}
type PostImage struct {
	ID       int    `json:"id" db:"id"`
	PostID   int    `json:"post_id" db:"post_id"`
	ImageURL string `json:"image_url" db:"image_url"`
}
type PostRepository interface {
	CreatePost(ctx context.Context, tx *sql.Tx, post *Post) (*Post, error)
	UpdatePost(ctx context.Context, post *Post) error
	DeletePost(ctx context.Context, tx *sql.Tx, id int) error
	GetPostByID(ctx context.Context, id int) (*Post, error)
	GetPostsByUserID(ctx context.Context, userID int) ([]*Post, error)
	BeginTransaction(ctx context.Context) (*sql.Tx, error)
}

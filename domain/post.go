package domain

import (
	"context"
	"database/sql"
	"time"
)

type Post struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	Content   string    `json:"content" db:"content"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
type PostImage struct {
	ID       int    `json:"id" db:"id"`
	PostID   int    `json:"post_id" db:"post_id"`
	ImageURL string `json:"image_url" db:"image_url"`
}
type PostRepository interface {
	CreatePost(ctx context.Context, tx *sql.Tx, post *Post) (*Post, error)
	UpdatePost(ctx context.Context, post *Post) error
	DeletePost(ctx context.Context, id int) error
	GetPostByID(ctx context.Context, id int) (*Post, error)
	GetPostsByUserID(ctx context.Context, userID int) ([]*Post, error)

	CreatePostImage(ctx context.Context, image *PostImage) (*PostImage, error)
	UpdatePostImage(ctx context.Context, image *PostImage) error
	DeletePostImage(ctx context.Context, id int) error
	GetPostImageByID(ctx context.Context, id int) (*PostImage, error)

	BeginTransaction(ctx context.Context) (*sql.Tx, error)
}

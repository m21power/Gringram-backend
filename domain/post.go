package domain

import "time"

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
	CreatePost(post *Post) (*Post, error)
	UpdatePost(post *Post) error
	DeletePost(id int) error
	GetPostByID(id int) (*Post, error)
	GetPostsByUserID(userID int) ([]*Post, error)

	CreatePostImage(image *PostImage) (*PostImage, error)
	UpdatePostImage(image *PostImage) error
	DeletePostImage(id int) error
	GetPostImage(id int) (*PostImage, error)
}

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
type FeedPayload struct {
	UnseenPost []*Post `json:"unseen_post"`
	SeenPost   []*Post `json:"seen_post"`
}
type PostRepository interface {
	CreatePost(ctx context.Context, tx *sql.Tx, post *Post) (*Post, error)
	UpdatePost(ctx context.Context, post *Post) error
	GetPosts(ctx context.Context) ([]*Post, error)
	DeletePost(ctx context.Context, tx *sql.Tx, id int) error
	GetPostByID(ctx context.Context, id int) (*Post, error)
	GetPostsByUserID(ctx context.Context, userID int) ([]*Post, error)
	IncrementCommentCount(ctx context.Context, id int) error
	IncrementLikeCount(ctx context.Context, id int) error
	DecrementCommentCount(ctx context.Context, postID int, commentID int) error
	DecrementLikeCount(ctx context.Context, id int) error
	BeginTransaction(ctx context.Context) (*sql.Tx, error)
	//comment
	CreateComment(ctx context.Context, tx *sql.Tx, comment *Comment) (*Comment, error)
	UpdateComment(ctx context.Context, comment *Comment) error
	DeleteComment(ctx context.Context, tx *sql.Tx, id int) error
	GetCommentByID(ctx context.Context, id int) (*Comment, error)

	// like
	MakeLike(ctx context.Context, like *Like) (*Like, error)
	GetLikers(ctx context.Context, postID int) ([]int, error)

	// interaction
	GetUnseenPostID(ctx context.Context, userID int) ([][]int, error)
	ViewPost(ctx context.Context, userID int, postID int) error

	//waitinglist
	UpdateWaitingList(ctx context.Context, tx *sql.Tx, postId int, status string) error
}

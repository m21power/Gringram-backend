package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/m21power/GrinGram/domain"
)

type PostStore struct {
	db *sql.DB
}

func NewPostStore(db *sql.DB) *PostStore {
	return &PostStore{db: db}
}
func (s *PostStore) CreatePost(ctx context.Context, tx *sql.Tx, post *domain.Post) (*domain.Post, error) {
	query := "INSERT INTO posts(user_id,content,image_url) VALUES(?,?,?)"
	res, err := tx.ExecContext(ctx, query, post.UserID, post.Content, post.Image_url)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	post.ID = int(id)
	post.CreatedAt = time.Now()
	return post, nil
}
func (s *PostStore) UpdatePost(ctx context.Context, post *domain.Post) error {
	query := "UPDATE posts SET content=?,image_url=? WHERE id=?"
	_, err := s.db.ExecContext(ctx, query, post.Content, post.Image_url, post.ID)
	if err != nil {
		return err
	}
	return nil

}

func (s *PostStore) DeletePost(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM posts WHERE id=?"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostStore) GetPostByID(ctx context.Context, id int) (*domain.Post, error) {
	query := "SELECT * FROM posts WHERE id=?"
	row := s.db.QueryRowContext(ctx, query, id)
	var post domain.Post
	err := row.Scan(&post.ID, &post.Content, &post.UserID, &post.Status, &post.Image_url, &post.Likes_count, &post.Comments_count, &post.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (s *PostStore) GetPostsByUserID(ctx context.Context, userID int) ([]*domain.Post, error) {
	query := "SELECT * FROM posts WHERE user_id=?"
	rows, err := s.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	posts, err := scanIntoList(rows)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
func (s *PostStore) BeginTransaction(ctx context.Context) (*sql.Tx, error) {
	return s.db.BeginTx(ctx, nil)
}
func scanIntoList(rows *sql.Rows) ([]*domain.Post, error) {
	var ans []*domain.Post
	defer rows.Close()
	for rows.Next() {
		var post domain.Post
		err := rows.Scan(&post.ID, &post.Content, &post.UserID, &post.Status, &post.Image_url, &post.Likes_count, &post.Comments_count, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		ans = append(ans, &post)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ans, nil
}

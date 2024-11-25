package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/m21power/GrinGram/domain"
)

func (s *PostStore) MakeLike(ctx context.Context, like *domain.Like) (*domain.Like, error) {
	// if like request send again it is going to be disliked
	row := s.db.QueryRowContext(ctx, "SELECT * FROM likes WHERE user_id=? && post_id=?", like.UserID, like.PostID)
	var likePayload domain.Like
	err := row.Scan(&likePayload.ID, &likePayload.UserID, &likePayload.PostID, &likePayload.CreatedAt)
	if err == sql.ErrNoRows {
		query := "INSERT INTO likes(user_id,post_id) VALUES(?,?)"
		res, err := s.db.ExecContext(ctx, query, like.UserID, like.PostID)
		if err != nil {
			return nil, err
		}
		id, err := res.LastInsertId()
		if err != nil {
			return nil, err
		}
		err = s.IncrementLikeCount(ctx, like.PostID)
		if err != nil {
			return nil, err
		}
		like.ID = int(id)
		like.CreatedAt = time.Now()

		return like, nil
	}
	if err != nil {
		return nil, err
	}
	query := "DELETE FROM likes WHERE user_id=? and post_id=?"
	_, err = s.db.ExecContext(ctx, query, like.UserID, like.PostID)
	if err != nil {
		return nil, err
	}
	err = s.DecrementLikeCount(ctx, like.PostID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
func (s *PostStore) GetLikers(ctx context.Context, postID int) ([]int, error) {
	query := "SELECT user_id FROM likes WHERE post_id=?"
	rows, err := s.db.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []int
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		result = append(result, id)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *PostStore) IncrementLikeCount(ctx context.Context, id int) error {
	query := "UPDATE posts SET likes_count=likes_count + 1 WHERE id=?"
	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
func (s *PostStore) DecrementLikeCount(ctx context.Context, id int) error {
	query := "UPDATE posts SET likes_count=likes_count - 1 WHERE id=?"
	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

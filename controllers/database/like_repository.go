package database

import (
	"context"
	"time"

	"github.com/m21power/GrinGram/domain"
)

func (s *PostStore) MakeLike(ctx context.Context, like *domain.Like) (*domain.Like, error) {
	query := "INSERT INTO likes(user_id,post_id) VALUES(?,?)"
	res, err := s.db.ExecContext(ctx, query, like.UserID, like.PostID)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	like.ID = int(id)
	like.CreatedAt = time.Now()

	return like, nil
}
func (s *PostStore) DisLike(ctx context.Context, like *domain.Like) error {
	query := "DELETE FROM likes WHERE user_id=? and post_id=?"
	_, err := s.db.ExecContext(ctx, query, like.UserID, like.PostID)
	if err != nil {
		return err
	}
	return nil

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

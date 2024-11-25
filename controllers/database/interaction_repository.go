package database

import (
	"context"
	"database/sql"
)

func (s *PostStore) CreateInteraction(ctx context.Context, tx *sql.Tx, postId int) error {
	var UserID []int
	rows, err := tx.QueryContext(ctx, "SELECT id FROM users")
	if err != nil {
		return err
	}
	for rows.Next() {
		var userId int
		err := rows.Scan(&userId)
		if err != nil {
			return err
		}
		UserID = append(UserID, userId)
	}
	for _, userID := range UserID {
		query := "INSERT INTO interactions(user_id,post_id) VALUES(?,?)"
		_, err := tx.ExecContext(ctx, query, userID, postId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (s *PostStore) GetUnseenPost(ctx context.Context, userID int) ([]int, error) {
	var PostID []int
	query := "SELECT post_id FROM interactions WHERE seen=? and user_id=?"
	rows, err := s.db.QueryContext(ctx, query, false, userID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var postId int
		err := rows.Scan(&postId)
		if err != nil {
			return nil, err
		}
		PostID = append(PostID, postId)
	}
	return PostID, nil
}

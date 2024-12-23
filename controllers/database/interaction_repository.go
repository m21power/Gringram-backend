package database

import (
	"context"
	"database/sql"
	"time"
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
func (s *PostStore) GetUnseenPostID(ctx context.Context, userID int) ([][]int, error) {
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
	var seenPostId []int
	query = "SELECT post_id FROM interactions WHERE seen=? and user_id=?"
	rows, err = s.db.QueryContext(ctx, query, true, userID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var postId int
		err := rows.Scan(&postId)
		if err != nil {
			return nil, err
		}
		seenPostId = append(seenPostId, postId)
	}
	var result [][]int
	result = append(result, PostID)
	result = append(result, seenPostId)

	return result, nil
}
func (s *PostStore) ViewPost(ctx context.Context, userId int, postID int) error {
	query := "UPDATE interactions SET seen=?,seen_at=? WHERE user_id=? and post_id=?"
	_, err := s.db.ExecContext(ctx, query, true, time.Now(), userId, postID)
	if err != nil {
		return err
	}
	return nil
}
func (s *PostStore) AddToWaitingList(ctx context.Context, tx *sql.Tx, postId int) error {
	query := "INSERT INTO waitinglist(post_id) VALUES(?)"
	_, err := tx.ExecContext(ctx, query, postId)
	if err != nil {
		return err
	}
	return nil
}
func (s *PostStore) UpdateWaitingList(ctx context.Context, tx *sql.Tx, postId int, status string) error {
	query := "UPDATE waitinglist SET status=? WHERE post_id=?"
	_, err := tx.ExecContext(ctx, query, status, postId)
	if err != nil {
		return err
	}
	if status == "approved" {
		err := s.CreateInteraction(ctx, tx, postId)
		if err != nil {
			return err
		}
		err = s.DeleteFromWaitingList(ctx, tx, postId)
		if err != nil {
			return err
		}
	}
	if status == "rejected" {
		err = s.DeleteFromWaitingList(ctx, tx, postId)
		if err != nil {
			return err
		}
		err = s.DeletePost(ctx, tx, postId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (s *PostStore) DeleteFromWaitingList(ctx context.Context, tx *sql.Tx, postId int) error {
	query := "DELETE FROM waitinglist WHERE post_id=?"
	_, err := tx.ExecContext(ctx, query, postId)
	if err != nil {
		return err
	}
	return nil
}

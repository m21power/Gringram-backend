package database

import (
	"context"
	"database/sql"

	"github.com/m21power/GrinGram/domain"
)

func (s *PostStore) CreateComment(ctx context.Context, tx *sql.Tx, comment *domain.Comment) (*domain.Comment, error) {
	if comment.ParentID.Valid {
		query := "INSERT INTO comments(text,user_id,post_id,parent_id) VALUES(?,?,?,?)"
		res, err := tx.ExecContext(ctx, query, comment.Text, comment.UserID, comment.PostID, comment.ParentID)
		if err != nil {
			return nil, err
		}
		id, err := res.LastInsertId()
		if err != nil {
			return nil, err
		}
		comment.ID = int(id)
		return comment, nil
	} else {
		query := "INSERT INTO comments(text,user_id,post_id) VALUES(?,?,?)"
		res, err := tx.ExecContext(ctx, query, comment.Text, comment.UserID, comment.PostID)
		if err != nil {
			return nil, err
		}
		id, err := res.LastInsertId()
		if err != nil {
			return nil, err
		}
		comment.ID = int(id)
		return comment, nil
	}

}
func (s *PostStore) UpdateComment(ctx context.Context, comment *domain.Comment) error {
	query := "UPDATE comments SET text=? WHERE id=?"
	_, err := s.db.Exec(query, comment.Text, comment.ID)
	if err != nil {
		return err
	}
	return nil
}
func (s *PostStore) DeleteComment(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM comments WHERE id=?"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
func (s *PostStore) GetCommentByID(ctx context.Context, id int) (*domain.Comment, error) {
	query := "SELECT * FROM comments WHERE id=?"
	var comment domain.Comment
	row := s.db.QueryRow(query, id)
	err := row.Scan(&comment.ID, &comment.Text, &comment.UserID, &comment.PostID, &comment.ParentID, &comment.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &comment, nil

}
func (s *PostStore) IncrementCommentCount(ctx context.Context, tx *sql.Tx, id int) error {
	query := "UPDATE posts SET comments_count=comments_count + 1 WHERE id=?"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
func (s *PostStore) IncrementLikeCount(ctx context.Context, tx *sql.Tx, id int) error {
	query := "UPDATE posts SET likes_count=likes_count + 1 WHERE id=?"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
func (s *PostStore) DecrementCommentCount(ctx context.Context, tx *sql.Tx, postID int, commentID int) error {
	row := s.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM comments WHERE parent_id=?", commentID)
	var count int
	if err := row.Scan(&count); err != nil {
		return err
	}
	query := "UPDATE posts SET comments_count=comments_count - ? WHERE id=?"
	_, err := tx.ExecContext(ctx, query, count+1, postID)
	if err != nil {
		return err
	}
	return nil
}
func (s *PostStore) DecrementLikeCount(ctx context.Context, tx *sql.Tx, id int) error {
	query := "UPDATE posts SET likes_count=likes_count - 1 WHERE id=?"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
func (s *PostStore) BeginTransaction(ctx context.Context) (*sql.Tx, error) {
	return s.db.BeginTx(ctx, nil)
}

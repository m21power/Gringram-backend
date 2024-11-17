package database

import (
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
func (s *PostStore) CreatePost(post *domain.Post) (*domain.Post, error) {
	query := "INSERT INTO post(user_id,content) VALUES(?,?)"
	res, err := s.db.Exec(query)
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
func (s *PostStore) UpdatePost(post *domain.Post) error {
	query := "UPDATE post SET content=? WHERE id=?"
	_, err := s.db.Exec(query, post.Content, post.ID)
	if err != nil {
		return err
	}
	return nil

}

func (s *PostStore) DeletePost(id int) error {
	query := "DELETE FROM post WHERE id=?"
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostStore) GetPostByID(id int) (*domain.Post, error) {
	query := "SELECT * FROM post WHERE id=?"
	row := s.db.QueryRow(query, id)
	var post domain.Post
	err := row.Scan(&post.ID, &post.UserID, &post.Content, &post.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (s *PostStore) GetPostsByUserID(userID int) ([]*domain.Post, error) {
	query := "SELECT * FROM post WHERE user_id=?"
	rows, err := s.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	posts, err := scanIntoList(rows)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostStore) CreatePostImage(image *domain.PostImage) (*domain.PostImage, error) {
	query := "INSERT INTO post_image(post_id,image_url) VALUES(?,?)"
	res, err := s.db.Exec(query, image.PostID, image.ImageURL)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	image.ID = int(id)
	return image, nil
}
func (s *PostStore) UpdatePostImage(image *domain.PostImage) error {
	query := "UPDATE post_image SET image_url=? WHERE id=?"
	_, err := s.db.Exec(query, image.ID)
	if err != nil {
		return err
	}
	return nil
}
func (s *PostStore) DeletePostImage(id int) error {
	query := "DELETE FROM post_image WHERE id=?"
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
func (s *PostStore) GetPostImage(id int) (*domain.PostImage, error) {
	query := "SELECT * FROM post_image WHERE id=?"
	row := s.db.QueryRow(query, id)
	var image domain.PostImage
	err := row.Scan(&image.ID, &image.PostID, &image.ImageURL)
	if err != nil {
		return nil, err
	}
	return &image, nil
}
func scanIntoList(rows *sql.Rows) ([]*domain.Post, error) {
	var ans []*domain.Post
	defer rows.Close()
	for rows.Next() {
		var post domain.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Content, &post.CreatedAt)
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

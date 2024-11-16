package database

import (
	"database/sql"
	"time"

	auth "github.com/m21power/GrinGram/Auth"
	"github.com/m21power/GrinGram/domain"
)

type UserStore struct {
	db *sql.DB
}

func UserNewStore(db *sql.DB) *UserStore {
	return &UserStore{db: db}
}

func (s *UserStore) CreateUser(user *domain.User) (*domain.User, error) {
	query := "INSERT INTO user(name,username,email,password,bio,profile_image_url) VALUES(?,?,?,?,?,?)"
	hashedPassword, err := auth.HashedPassword(user.Password)
	if err != nil {
		return nil, err
	}
	res, err := s.db.Exec(query, user.Name, user.Username, user.Email, hashedPassword, user.Bio, user.ProfileImageUrl)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.ID = int(id)
	user.CreatedAt = time.Now()
	return user, nil
}

func (s *UserStore) GetUserByID(id int) (*domain.User, error) {
	query := "SELECT id, name, username, COALESCE(bio, ''), password, email, profile_image_url,created_at FROM user WHERE id=?"
	row := s.db.QueryRow(query, id)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Username, &user.Bio, &user.Password, &user.Email, &user.ProfileImageUrl, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil

}
func (s *UserStore) GetUserByUsername(username string) (*domain.User, error) {
	query := "SELECT id, name, username, COALESCE(bio, ''), password, email,profile_image_url, created_at FROM user WHERE username=?"
	row := s.db.QueryRow(query, username)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Username, &user.Bio, &user.Password, &user.Email, &user.ProfileImageUrl, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (s *UserStore) GetUserByEmail(email string) (*domain.User, error) {
	query := "SELECT id, name, username, COALESCE(bio, ''), password, email, profile_image_url,created_at FROM user WHERE email=?"

	row := s.db.QueryRow(query, email)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Username, &user.Bio, &user.Password, &user.Email, &user.ProfileImageUrl, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil

}
func (s *UserStore) UpdateUser(user *domain.User) error {

	query := "UPDATE user SET name=?,username=?,bio=?,password=?,email=?,profile_image_url=? WHERE id=?"
	_, err := s.db.Exec(query, user.Name, user.Username, user.Bio, user.Password, user.Email, user.ProfileImageUrl, user.ID)
	if err != nil {
		return err
	}
	return nil

}
func (s *UserStore) DeleteUser(id int) error {
	query := "DELETE FROM user WHERE id=?"
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserStore) DeleteUserImage(id int) error {
	query := "UPDATE user SET profile_image_url=? WHERE id=?"
	_, err := s.db.Exec(query, "", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserStore) GetProfileURL(id int) (string, error) {
	query := "SELECT profile_image_url FROM user WHERE id=?"
	row := s.db.QueryRow(query, id)
	var profileImageUrl string
	err := row.Scan(&profileImageUrl)
	if err != nil {
		return "", err
	}
	return profileImageUrl, nil
}

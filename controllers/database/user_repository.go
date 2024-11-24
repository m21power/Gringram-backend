package database

import (
	"context"
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

func (s *UserStore) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	query := "INSERT INTO users(name,username,email,password,bio,profile_url) VALUES(?,?,?,?,?,?)"
	hashedPassword, err := auth.HashedPassword(user.Password)
	if err != nil {
		return nil, err
	}
	res, err := s.db.ExecContext(ctx, query, user.Name, user.Username, user.Email, hashedPassword, user.Bio, user.ProfileImageUrl)
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

func (s *UserStore) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	query := "SELECT id, name, username, COALESCE(bio, ''), password, role,email, profile_url,created_at FROM users WHERE id=?"
	row := s.db.QueryRowContext(ctx, query, id)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Username, &user.Bio, &user.Password, &user.Role, &user.Email, &user.ProfileImageUrl, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil

}
func (s *UserStore) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	query := "SELECT id, name, username, COALESCE(bio, ''), password, role,email, profile_url,created_at FROM users WHERE username=?"
	row := s.db.QueryRowContext(ctx, query, username)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Username, &user.Bio, &user.Password, &user.Role, &user.Email, &user.ProfileImageUrl, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (s *UserStore) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := "SELECT id, name, username, COALESCE(bio, ''), password,role, email, profile_url,created_at FROM users WHERE email=?"

	row := s.db.QueryRowContext(ctx, query, email)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Username, &user.Bio, &user.Password, &user.Role, &user.Email, &user.ProfileImageUrl, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil

}
func (s *UserStore) UpdateUser(ctx context.Context, user *domain.User) error {
	query := "UPDATE users SET name=?,username=?,bio=?,password=?,role=?,email=?,profile_url=? WHERE id=?"
	_, err := s.db.ExecContext(ctx, query, user.Name, user.Username, user.Bio, user.Password, user.Role, user.Email, user.ProfileImageUrl, user.ID)
	if err != nil {
		return err
	}
	return nil

}
func (s *UserStore) DeleteUser(ctx context.Context, id int) error {
	query := "DELETE FROM users WHERE id=?"
	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserStore) DeleteUserImage(ctx context.Context, tx *sql.Tx, id int) error {
	query := "UPDATE users SET profile_url=? WHERE id=?"
	_, err := tx.ExecContext(ctx, query, "", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserStore) GetProfileURL(ctx context.Context, tx *sql.Tx, id int) (string, error) {
	query := "SELECT profile_url FROM users WHERE id=?"
	var url string
	err := tx.QueryRowContext(ctx, query, id).Scan(&url)
	if err != nil {
		return "", err
	}
	return url, nil
}
func (s *UserStore) BeginTransaction(ctx context.Context) (*sql.Tx, error) {
	return s.db.BeginTx(ctx, nil)
}

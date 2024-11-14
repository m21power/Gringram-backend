package database

import (
	"database/sql"

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
	query := "INSERT INTO user(name,username,password,email) VALUES(?,?,?,?)"
	hashedPassword, err := auth.HashedPassword(user.Password)
	if err != nil {
		return nil, err
	}
	res, err := s.db.Exec(query, user.Name, user.Username, hashedPassword, user.Email)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.ID = int(id)
	return user, nil
}

func (s *UserStore) GetUserByID(id int) (*domain.User, error) {
	query := "SELECT * FROM user WHERE id=?"
	row := s.db.QueryRow(query, id)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Username, &user.Bio, &user.Password, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil

}
func (s *UserStore) GetUserByUsername(username string) (*domain.User, error) {
	query := "SELECT * FROM user WHERE username=?"
	row := s.db.QueryRow(query, username)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Username, &user.Bio, &user.Password, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (s *UserStore) GetUserByEmail(email string) (*domain.User, error) {
	query := "SELECT * FROM user WHERE email=?"
	row := s.db.QueryRow(query, email)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Username, &user.Bio, &user.Password, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil

}
func (s *UserStore) UpdateUser(user *domain.User) error {
	query := "UPDATE user SET name=?,username=?,bio=?,password=?,email=? WHERE id=?"
	_, err := s.db.Exec(query, user.Name, user.Username, user.Bio, user.Password, user.Email, user.ID)
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

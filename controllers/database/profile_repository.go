package database

import (
	"database/sql"

	"github.com/m21power/GrinGram/domain"
)

type ProfileStore struct {
	db *sql.DB
}

func NewProfileStore(db *sql.DB) *ProfileStore {
	return &ProfileStore{db: db}
}
func (s *ProfileStore) CreateProfile(profile *domain.Profile) (*domain.Profile, error) {
	query := "INSERT INTO profile_image (user_id,url) VALUES(?,?)"
	res, err := s.db.Exec(query, profile.User_ID, profile.ImageURL)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	profile.ID = int(id)
	return profile, nil
}

func (s *ProfileStore) GetProfileByID(id int) (*domain.Profile, error) {
	query := "SELECT * FROM profile_image WHERE id=?"
	row := s.db.QueryRow(query, id)
	profile := &domain.Profile{}
	err := row.Scan(&profile.ID, &profile.User_ID, &profile.ImageURL, &profile.CreatedAt)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (s *ProfileStore) UpdateProfile(profile *domain.Profile) error {
	query := "UPDATE profile_image SET url=? WHERE id=?"
	_, err := s.db.Exec(query, profile.ImageURL, profile.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProfileStore) DeleteProfile(id int) error {
	query := "DELETE FROM profile_image WHERE id=?"
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

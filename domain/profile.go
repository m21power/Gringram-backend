package domain

type Profile struct {
	ID        int    `json:"id" db:"id"`
	User_ID   int    `json:"user_id" db:"user_id"`
	ImageURL  string `json:"url" db:"url"`
	CreatedAt string `json:"created_at" db:"created_at"`
}
type ProfileRepository interface {
	CreateProfile(profile *Profile) (*Profile, error)
	GetProfileByID(id int) (*Profile, error)
	UpdateProfile(profile *Profile) error
	DeleteProfile(id int) error
}

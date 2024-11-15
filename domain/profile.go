package domain

type Profile struct {
	ID        int    `json:"id" db:"id"`
	Url       string `json:"url" db:"url"`
	CreatedAt string `json:"created_at" db:"created_at"`
}
type ProfileRepository interface {
	CreateProfile(profile *Profile) (*Profile, error)
	GetProfileByID(id int) (*Profile, error)
	UpdateProfile(profile *Profile) error
	DeleteProfile(id int) error
}

package types

type UserPayload struct {
	Name     string `json:"name" db:"name"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Email    string `json:"email" db:"email"`
	Bio      string `json:"bio" db:"bio"`
	// ProfileImageUrl string `json:"image_url" db:"profile_image_url"`
}

package types

type UserPayload struct {
	Name     string `json:"name" db:"name"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Email    string `json:"email" db:"email"`
}

type ProfilePayload struct {
	Url string `json:"url" db:"url"`
}

type UpdateProfilePayload struct {
	ID  int    `json:"id" db:"id"`
	Url string `json:"url" db:"url"`
}

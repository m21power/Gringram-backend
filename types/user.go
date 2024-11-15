package types

type UserPayload struct {
	Name     string `json:"name" db:"name"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Email    string `json:"email" db:"email"`
}

type ImagePayload struct {
	User_ID string `json:"user_id"`
	// ImageData []byte `json:"-"`
}
type UpdatePayload struct {
	ID int `json:"id"`
}

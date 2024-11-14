package domain

type User struct {
	ID        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Username  string `json:"username" db:"username"`
	Bio       string `json:"bio" db:"bio"`
	Password  string `json:"password" db:"password"`
	Email     string `json:"email" db:"email"`
	ProfileID int    `json:"profile_id" db:"profile_id"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

type UserRepository interface {
	CreateUser(user *User) (*User, error)
	GetUserByID(id int) (*User, error)
	GetUserByUsername(username string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id int) error
}

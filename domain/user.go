package domain

import "time"

type User struct {
	ID              int       `json:"id" db:"id"`
	Name            string    `json:"name" db:"name"`
	Username        string    `json:"username" db:"username"`
	Password        string    `json:"password" db:"password"`
	Email           string    `json:"email" db:"email"`
	Bio             string    `json:"bio" db:"bio"`                     // optional
	ProfileImageUrl string    `json:"image_url" db:"profile_image_url"` // optional
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
}

type UserRepository interface {
	CreateUser(user *User) (*User, error)
	GetUserByID(id int) (*User, error)
	GetUserByUsername(username string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id int) error
}

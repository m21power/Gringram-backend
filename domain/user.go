package domain

import (
	"context"
	"database/sql"
	"time"
)

type User struct {
	ID              int       `json:"id" db:"id"`
	Name            string    `json:"name" db:"name"`
	Username        string    `json:"username" db:"username"`
	Password        string    `json:"password" db:"password"`
	Email           string    `json:"email" db:"email"`
	Bio             string    `json:"bio" db:"bio"`
	Role            string    `json:"role" db:"role"`
	ProfileImageUrl string    `json:"image_url" db:"profile_image_url"` // optional
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByID(ctx context.Context, id int) (*User, error)
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	UpdateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, id int) error
	DeleteUserImage(ctx context.Context, tx *sql.Tx, id int) error
	GetProfileURL(ctx context.Context, tx *sql.Tx, id int) (string, error)
	BeginTransaction(ctx context.Context) (*sql.Tx, error)
}

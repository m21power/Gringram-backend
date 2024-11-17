package usecases

import (
	"context"
	"database/sql"

	"github.com/m21power/GrinGram/domain"
)

type UserUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(userRepository domain.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepository: userRepository,
	}
}

func (u *UserUsecase) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	return u.userRepository.CreateUser(ctx, user)
}
func (u *UserUsecase) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	return u.userRepository.GetUserByID(ctx, id)
}
func (u *UserUsecase) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	return u.userRepository.GetUserByUsername(ctx, username)
}

func (u *UserUsecase) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	return u.userRepository.GetUserByEmail(ctx, email)
}
func (u *UserUsecase) UpdateUser(ctx context.Context, user *domain.User) error {
	return u.userRepository.UpdateUser(ctx, user)
}
func (u *UserUsecase) DeleteUser(ctx context.Context, id int) error {
	return u.userRepository.DeleteUser(ctx, id)
}
func (u *UserUsecase) DeleteUserImage(ctx context.Context, tx *sql.Tx, id int) error {
	return u.userRepository.DeleteUserImage(ctx, tx, id)
}
func (u *UserUsecase) GetProfileURL(ctx context.Context, tx *sql.Tx, id int) (string, error) {
	return u.userRepository.GetProfileURL(ctx, tx, id)
}
func (u *UserUsecase) BeginTransaction(ctx context.Context) (*sql.Tx, error) {
	return u.userRepository.BeginTransaction(ctx)
}

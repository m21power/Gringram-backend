package usecases

import (
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

func (u *UserUsecase) CreateUser(user *domain.User) (*domain.User, error) {
	return u.userRepository.CreateUser(user)
}
func (u *UserUsecase) GetUserByID(id int) (*domain.User, error) {
	return u.userRepository.GetUserByID(id)
}
func (u *UserUsecase) GetUserByUsername(username string) (*domain.User, error) {
	return u.userRepository.GetUserByUsername(username)
}

func (u *UserUsecase) GetUserByEmail(email string) (*domain.User, error) {
	return u.userRepository.GetUserByEmail(email)
}
func (u *UserUsecase) UpdateUser(user *domain.User) error {
	return u.userRepository.UpdateUser(user)
}
func (u *UserUsecase) DeleteUser(id int) error {
	return u.userRepository.DeleteUser(id)
}

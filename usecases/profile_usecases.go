package usecases

import "github.com/m21power/GrinGram/domain"

type ProfileUsecase struct {
	profileRepository domain.ProfileRepository
}

func NewProfileUsecase(profileRepository domain.ProfileRepository) *ProfileUsecase {
	return &ProfileUsecase{profileRepository: profileRepository}
}

func (u *ProfileUsecase) CreateProfile(profile *domain.Profile) (*domain.Profile, error) {
	return u.profileRepository.CreateProfile(profile)
}
func (u *ProfileUsecase) GetProfileByID(id int) (*domain.Profile, error) {
	return u.profileRepository.GetProfileByID(id)
}
func (u *ProfileUsecase) UpdateProfile(profile *domain.Profile) error {
	return u.profileRepository.UpdateProfile(profile)
}
func (u *ProfileUsecase) DeleteProfile(id int) error {
	return u.profileRepository.DeleteProfile(id)
}

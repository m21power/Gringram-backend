package usecases

import "github.com/m21power/GrinGram/domain"

type PostUsecase struct {
	postRepository domain.PostRepository
}

func NewPostRepository(postRepository domain.PostRepository) *PostUsecase {
	return &PostUsecase{postRepository: postRepository}
}

func (u *PostUsecase) CreatePost(post *domain.Post) (*domain.Post, error) {
	return u.postRepository.CreatePost(post)
}
func (u *PostUsecase) UpdatePost(post *domain.Post) error {
	return u.postRepository.UpdatePost(post)
}
func (u *PostUsecase) DeletePost(id int) error {
	return u.postRepository.DeletePost(id)
}
func (u *PostUsecase) GetPostByID(id int) (*domain.Post, error) {
	return u.postRepository.GetPostByID(id)
}
func (u *PostUsecase) GetPostsByUserID(userID int) ([]*domain.Post, error) {
	return u.postRepository.GetPostsByUserID(userID)
}

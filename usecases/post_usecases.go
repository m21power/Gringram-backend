package usecases

import (
	"context"
	"database/sql"

	"github.com/m21power/GrinGram/domain"
)

type PostUsecase struct {
	postRepository domain.PostRepository
}

func NewPostRepository(postRepository domain.PostRepository) *PostUsecase {
	return &PostUsecase{postRepository: postRepository}
}

func (u *PostUsecase) CreatePost(ctx context.Context, tx *sql.Tx, post *domain.Post) (*domain.Post, error) {
	return u.postRepository.CreatePost(ctx, tx, post)
}
func (u *PostUsecase) UpdatePost(ctx context.Context, post *domain.Post) error {
	return u.postRepository.UpdatePost(ctx, post)
}
func (u *PostUsecase) DeletePost(ctx context.Context, tx *sql.Tx, id int) error {
	return u.postRepository.DeletePost(ctx, tx, id)
}
func (u *PostUsecase) GetPostByID(ctx context.Context, id int) (*domain.Post, error) {
	return u.postRepository.GetPostByID(ctx, id)
}
func (u *PostUsecase) GetPostsByUserID(ctx context.Context, userID int) ([]*domain.Post, error) {
	return u.postRepository.GetPostsByUserID(ctx, userID)
}
func (u *PostUsecase) IncrementCommentCount(ctx context.Context, tx *sql.Tx, id int) error {
	return u.postRepository.IncrementCommentCount(ctx, tx, id)
}
func (u *PostUsecase) IncrementLikeCount(ctx context.Context, tx *sql.Tx, id int) error {
	return u.postRepository.IncrementLikeCount(ctx, tx, id)
}
func (u *PostUsecase) DecrementCommentCount(ctx context.Context, tx *sql.Tx, postID int, commentID int) error {
	return u.postRepository.DecrementCommentCount(ctx, tx, postID, commentID)
}
func (u *PostUsecase) DecrementLikeCount(ctx context.Context, tx *sql.Tx, id int) error {
	return u.postRepository.DecrementLikeCount(ctx, tx, id)
}
func (u *PostUsecase) BeginTransaction(ctx context.Context) (*sql.Tx, error) {
	return u.postRepository.BeginTransaction(ctx)
}

// comment
func (u *PostUsecase) CreateComment(ctx context.Context, tx *sql.Tx, comment *domain.Comment) (*domain.Comment, error) {
	return u.postRepository.CreateComment(ctx, tx, comment)
}
func (u *PostUsecase) UpdateComment(ctx context.Context, comment *domain.Comment) error {
	return u.postRepository.UpdateComment(ctx, comment)
}
func (u *PostUsecase) DeleteComment(ctx context.Context, tx *sql.Tx, id int) error {
	return u.postRepository.DeleteComment(ctx, tx, id)
}
func (u *PostUsecase) GetCommentByID(ctx context.Context, id int) (*domain.Comment, error) {
	return u.postRepository.GetCommentByID(ctx, id)
}

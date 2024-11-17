package handlers

import (
	"net/http"

	"github.com/m21power/GrinGram/usecases"
)

type PostHandler struct {
	postUsecase usecases.PostUsecase
}

func NewPostHandler(postUsecase usecases.PostUsecase) *PostHandler {
	return &PostHandler{postUsecase: postUsecase}
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {

}
func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {

}
func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {

}
func (h *PostHandler) GetPostByID(w http.ResponseWriter, r *http.Request) {

}
func (h *PostHandler) GetPostsByUserID(w http.ResponseWriter, r *http.Request) {

}

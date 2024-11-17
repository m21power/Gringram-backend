package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/m21power/GrinGram/domain"
	"github.com/m21power/GrinGram/types"
	"github.com/m21power/GrinGram/usecases"
	"github.com/m21power/GrinGram/utils"
)

type PostHandler struct {
	postUsecase *usecases.PostUsecase
}

func NewPostHandler(postUsecase *usecases.PostUsecase) *PostHandler {
	return &PostHandler{postUsecase: postUsecase}
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tx, err := h.postUsecase.BeginTransaction(ctx)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	var post types.PostPayload
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	up := toDomainPost(post)
	p, err := h.postUsecase.CreatePost(ctx, tx, up)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, p)

}
func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var postPayload types.UpdatePostPayload
	err := json.NewDecoder(r.Body).Decode(&postPayload)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	id, err := utils.GetID(r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	oldPost, err := h.postUsecase.GetPostByID(ctx, id)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	post := fromUpdateToDomainPost(postPayload, oldPost)
	err = h.postUsecase.UpdatePost(ctx, post)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, post)

}
func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := utils.GetID(r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	err = h.postUsecase.DeletePost(ctx, id)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "deleted successfully!"})
}
func (h *PostHandler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := utils.GetID(r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	post, err := h.postUsecase.GetPostByID(ctx, id)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, post)
}
func (h *PostHandler) GetPostsByUserID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, err := utils.GetID(r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	posts, err := h.postUsecase.GetPostsByUserID(ctx, userID)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string][]*domain.Post{"posts": posts})

}
func (h *PostHandler) CreatePostImage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tx, err := h.postUsecase.BeginTransaction(ctx)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	url, err := utils.GetImageUrl(r)
	if url == "" {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"message": "image no uploaded"})
		return
	}
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	imagePayload, err := utils.GetImagePayload(w, r)
	if err != nil {
		err := utils.DeleteImageFromCloud(r, url)
		if err != nil {
			utils.WriteError(w, err)
			return
		}
		utils.WriteError(w, err)
		return
	}
	image := toPostImage(imagePayload, url)
	pimage, err := h.postUsecase.CreatePostImage(ctx, tx, image)
	if err != nil {
		if er := utils.DeleteImageFromCloud(r, url); er != nil {
			utils.WriteError(w, er)
			return
		}
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, pimage)

}
func (h *PostHandler) UpdatePostImage(w http.ResponseWriter, r *http.Request) {
	// this section is mess cuz there is no transaction in clouddinary
	ctx := r.Context()
	id, err := utils.GetID(r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	oldImage, err := h.postUsecase.GetPostImageByID(ctx, id)
	oldURL := oldImage.ImageURL
	if err != nil {
		utils.WriteError(w, err)
		return
	}

	url, err := utils.GetImageUrl(r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	im := updateOldImage(oldImage, url)
	err = h.postUsecase.UpdatePostImage(ctx, im)
	if err != nil {
		if er := utils.DeleteImageFromCloud(r, url); er != nil {
			utils.WriteError(w, er)
			return
		}
		utils.WriteError(w, err)
		return
	}
	err = utils.DeleteImageFromCloud(r, oldURL)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, im)
}
func (h *PostHandler) DeletePostImage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := utils.GetID(r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	image, err := h.postUsecase.GetPostImageByID(ctx, id)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	err = h.postUsecase.DeletePostImage(ctx, id)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	err = utils.DeleteImageFromCloud(r, image.ImageURL)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Deleted successfully"})

}
func (h *PostHandler) GetPostImageByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := utils.GetID(r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	image, err := h.postUsecase.GetPostImageByID(ctx, id)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, image)
}
func toDomainPost(post types.PostPayload) *domain.Post {
	return &domain.Post{UserID: post.UserID, Content: post.Content}
}
func fromUpdateToDomainPost(newPost types.UpdatePostPayload, oldPost *domain.Post) *domain.Post {
	if newPost.Content != "" {
		oldPost.Content = newPost.Content
	}
	return oldPost
}
func toPostImage(imagePayload *types.PostImagePayload, url string) *domain.PostImage {
	return &domain.PostImage{PostID: imagePayload.PostID, ImageURL: url}
}
func updateOldImage(old *domain.PostImage, url string) *domain.PostImage {
	if url != "" {
		old.ImageURL = url
	}
	return old
}

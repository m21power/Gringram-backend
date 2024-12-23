package handlers

import (
	"fmt"
	"net/http"

	auth "github.com/m21power/GrinGram/Auth"
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
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	userInfo, err := GetUserInfo(r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	var postContent types.PostContent
	_, err = utils.GetPayload(w, r, &postContent)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	url, err := utils.GetPostImageURL(r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	if url == "" {
		utils.WriteJSON(w, http.StatusAccepted, map[string]string{"message": "empty post!"})
		return
	}
	up := toDomainPost(userInfo.UserID, postContent.Content, url)
	p, err := h.postUsecase.CreatePost(ctx, tx, up)
	if err != nil {
		utils.DeleteImageFromCloud(r, url)
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]*domain.Post{"post": p})

}

func (h *PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	posts, err := h.postUsecase.GetPosts(ctx)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string][]*domain.Post{"posts": posts})
}
func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var postContent types.PostContent
	_, err := utils.GetPayload(w, r, &postContent)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	url, err := utils.GetPostImageURL(r)
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
	// is he allowed get it's value from the token
	Token, err := GetUserInfo(r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	if Token.Role == "user" && oldPost.UserID != Token.UserID {
		utils.WriteError(w, fmt.Errorf("you are not allowed to delete"))
		return
	}

	if oldPost.Image_url != "" {
		utils.DeleteImageFromCloud(r, oldPost.Image_url)
	}
	post := fromUpdateToDomainPost(postContent.Content, oldPost, url)
	err = h.postUsecase.UpdatePost(ctx, post)
	if err != nil {
		utils.DeleteImageFromCloud(r, url)
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, post)

}
func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	// before deleting first grap the image url of the post since
	// we have to delete from the cloud too
	ctx := r.Context()
	tx, err := h.postUsecase.BeginTransaction(ctx)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
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
	Token, err := GetUserInfo(r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	if Token.Role == "user" && post.UserID != Token.UserID {
		utils.WriteError(w, fmt.Errorf("you are not allowed to delete"))
		return
	}
	err = h.postUsecase.DeletePost(ctx, tx, id)
	if err != nil {
		utils.WriteError(w, err)
		return
	}

	utils.DeleteImageFromCloud(r, post.Image_url)
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
	utils.WriteJSON(w, http.StatusOK, map[string]*domain.Post{"post": post})
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
func (h *PostHandler) UpdateWaitingList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tx, err := h.postUsecase.BeginTransaction(ctx)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	postId, err := utils.GetID(r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	// err = json.NewDecoder(r.Body).Decode(&status)
	status := r.URL.Query().Get("status")
	if status == "" {
		utils.WriteError(w, fmt.Errorf("status is required"))
		return
	}
	err = h.postUsecase.UpdateWaitingList(ctx, tx, postId, status)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "updated successfully"})
}
func (h *PostHandler) GetFeed(w http.ResponseWriter, r *http.Request) {
	// later we take user id from our token
	// for now lets take from the request
	ctx := r.Context()
	userId, err := utils.GetID(r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	unseenPost, err := h.GetUnseenPost(ctx, userId)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]*domain.FeedPayload{"posts": unseenPost})
}
func toDomainPost(userId int, content string, url string) *domain.Post {
	return &domain.Post{UserID: userId, Content: content, Image_url: url}
}
func fromUpdateToDomainPost(content string, oldPost *domain.Post, url string) *domain.Post {
	if content != "" {
		oldPost.Content = content
	}
	if url != "" {
		oldPost.Image_url = url
	}
	return oldPost
}
func GetUserInfo(r *http.Request) (*types.Token, error) {
	token, err := auth.GetTokens(r)
	if err != nil {
		return nil, err
	}
	Token, err := auth.GetTokenValues(token)
	if err != nil {
		return nil, err
	}

	return Token, nil
}

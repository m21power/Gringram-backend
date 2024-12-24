package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/m21power/GrinGram/domain"
	"github.com/m21power/GrinGram/utils"
)

func (h *PostHandler) ViewPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	type ViewPayload struct {
		UserID int `json:"user_id"`
		PostID int `json:"post_id"`
	}
	var payload ViewPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	err = h.postUsecase.ViewPost(ctx, payload.UserID, payload.PostID)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	utils.WriteJSON(w, http.StatusOK, utils.ApiResponse{Message: "seen", Success: true})

}

func (h *PostHandler) GetUnseenPost(ctx context.Context, userID int) (*domain.FeedPayload, error) {
	result, err := h.postUsecase.GetUnseenPostID(ctx, userID)
	if err != nil {
		return nil, err
	}
	var Posts []*domain.Post
	for _, postID := range result[0] {
		post, err := h.postUsecase.GetPostByID(ctx, postID)
		if err != nil {
			return nil, err
		}
		Posts = append(Posts, post)
	}
	var seenPost []*domain.Post
	for _, postID := range result[1] {
		post, err := h.postUsecase.GetPostByID(ctx, postID)
		if err != nil {
			return nil, err
		}
		seenPost = append(seenPost, post)
	}
	var ans domain.FeedPayload
	ans.UnseenPost = Posts
	ans.SeenPost = seenPost
	return &ans, nil
}

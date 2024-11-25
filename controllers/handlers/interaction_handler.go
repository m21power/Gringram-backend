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
		utils.WriteError(w, err)
		return
	}
	err = h.postUsecase.ViewPost(ctx, payload.UserID, payload.PostID)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "seen"})

}

func (h *PostHandler) GetUnseenPost(ctx context.Context, userID int) ([]*domain.Post, error) {
	postsID, err := h.postUsecase.GetUnseenPostID(ctx, userID)
	if err != nil {
		return nil, err
	}
	var Posts []*domain.Post
	for _, postID := range postsID {
		post, err := h.postUsecase.GetPostByID(ctx, postID)
		if err != nil {
			return nil, err
		}
		Posts = append(Posts, post)
	}
	return Posts, nil
}

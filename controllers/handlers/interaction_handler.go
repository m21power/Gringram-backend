package handlers

import (
	"net/http"

	"github.com/m21power/GrinGram/domain"
	"github.com/m21power/GrinGram/utils"
)

func (h *PostHandler) GetUnseenPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId, err := utils.GetID(r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	postsID, err := h.postUsecase.GetUnseenPost(ctx, userId)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	var Posts []*domain.Post
	for _, postID := range postsID {
		post, err := h.postUsecase.GetPostByID(ctx, postID)
		if err != nil {
			utils.WriteError(w, err)
			return
		}
		Posts = append(Posts, post)
	}
	utils.WriteJSON(w, http.StatusOK, map[string][]*domain.Post{"Unseen Post": Posts})
}

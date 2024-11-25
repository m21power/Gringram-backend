package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/m21power/GrinGram/domain"
	"github.com/m21power/GrinGram/utils"
)

func (h *PostHandler) MakeLike(w http.ResponseWriter, r *http.Request) {
	var like domain.Like
	ctx := r.Context()
	err := json.NewDecoder(r.Body).Decode(&like)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	l, err := h.postUsecase.MakeLike(ctx, &like)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	if l == nil {
		utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "disliked"})
		return
	}
	utils.WriteJSON(w, http.StatusCreated, l)
}
func (h *PostHandler) GetLikers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := utils.GetID(r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	usersID, err := h.postUsecase.GetLikers(ctx, id)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string][]int{"usersID": usersID})
}

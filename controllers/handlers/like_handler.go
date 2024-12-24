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
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	l, err := h.postUsecase.MakeLike(ctx, &like)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	if l == nil {
		utils.WriteJSON(w, http.StatusOK, utils.ApiResponse{Message: "disliked", Success: true})
		return
	}
	utils.WriteJSON(w, http.StatusCreated, utils.ApiResponse{Message: "liked", Success: true})
}
func (h *PostHandler) GetLikers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := utils.GetID(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	usersID, err := h.postUsecase.GetLikers(ctx, id)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	utils.WriteJSON(w, http.StatusOK, utils.ApiResponse{Message: "likes found", Success: true, Data: usersID})
}

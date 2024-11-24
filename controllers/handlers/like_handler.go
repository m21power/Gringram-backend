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
	err = json.NewDecoder(r.Body).Decode(&like)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	l, err := h.postUsecase.MakeLike(ctx, &like)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	err = h.postUsecase.IncrementLikeCount(ctx, tx, like.PostID)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, l)
}
func (h *PostHandler) DisLike(w http.ResponseWriter, r *http.Request) {
	var like domain.Like
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
	err = json.NewDecoder(r.Body).Decode(&like)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	err = h.postUsecase.DisLike(ctx, &like)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	err = h.postUsecase.DecrementLikeCount(ctx, tx, like.PostID)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "disliked"})
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

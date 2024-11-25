package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/m21power/GrinGram/domain"
	"github.com/m21power/GrinGram/utils"
)

func (h *PostHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tx, err := h.postUsecase.BeginTransaction(ctx)
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
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	var commentPayload domain.Comment
	err = json.NewDecoder(r.Body).Decode(&commentPayload)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	comment, err := h.postUsecase.CreateComment(ctx, tx, &commentPayload)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	err = h.postUsecase.IncrementCommentCount(ctx, comment.PostID)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, map[string]domain.Comment{"comment": *comment})

}
func (h *PostHandler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	type Payload struct {
		Text string `json:"text"`
	}
	var payload Payload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	id, err := utils.GetID(r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	oldComment, err := h.postUsecase.GetCommentByID(ctx, id)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	if payload.Text != "" {
		oldComment.Text = payload.Text
	}
	err = h.postUsecase.UpdateComment(ctx, oldComment)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]domain.Comment{"comment": *oldComment})

}
func (h *PostHandler) GetCommentByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := utils.GetID(r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	comment, err := h.postUsecase.GetCommentByID(ctx, id)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]domain.Comment{"comment": *comment})
}
func (h *PostHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tx, err := h.postUsecase.BeginTransaction(ctx)
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
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	id, err := utils.GetID(r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	comment, err := h.postUsecase.GetCommentByID(ctx, id)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	err = h.postUsecase.DecrementCommentCount(ctx, comment.PostID, id)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	err = h.postUsecase.DeleteComment(ctx, tx, id)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "comment deleted successfully!"})
}

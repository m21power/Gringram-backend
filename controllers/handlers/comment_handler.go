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
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	var commentPayload domain.Comment
	err = json.NewDecoder(r.Body).Decode(&commentPayload)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	comment, err := h.postUsecase.CreateComment(ctx, tx, &commentPayload)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	err = h.postUsecase.IncrementCommentCount(ctx, comment.PostID)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	utils.WriteJSON(w, http.StatusCreated, utils.ApiResponse{Message: "comment found", Success: true, Data: comment})

}
func (h *PostHandler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	type Payload struct {
		Text string `json:"text"`
	}
	var payload Payload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	id, err := utils.GetID(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	oldComment, err := h.postUsecase.GetCommentByID(ctx, id)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	if payload.Text != "" {
		oldComment.Text = payload.Text
	}
	err = h.postUsecase.UpdateComment(ctx, oldComment)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	utils.WriteJSON(w, http.StatusOK, utils.ApiResponse{Message: "comment updated", Success: true, Data: oldComment})

}
func (h *PostHandler) GetCommentByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := utils.GetID(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	comment, err := h.postUsecase.GetCommentByID(ctx, id)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	utils.WriteJSON(w, http.StatusOK, utils.ApiResponse{Message: "comment found", Success: true, Data: comment})
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
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	id, err := utils.GetID(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	comment, err := h.postUsecase.GetCommentByID(ctx, id)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	err = h.postUsecase.DecrementCommentCount(ctx, comment.PostID, id)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	err = h.postUsecase.DeleteComment(ctx, tx, id)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	utils.WriteJSON(w, http.StatusOK, utils.ApiResponse{Message: "comment deleted successfully", Success: true})
}

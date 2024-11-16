package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/m21power/GrinGram/domain"
	"github.com/m21power/GrinGram/types"
	"github.com/m21power/GrinGram/usecases"
	"github.com/m21power/GrinGram/utils"
)

type ProfileHandler struct {
	usecases *usecases.ProfileUsecase
}

// cloud name : dl6vahv6t
// api key : 639632577282947
// api secret: _qyu3umAppkfaRNR84QUuWiIa7U
func NewProfileHandler(usecase *usecases.ProfileUsecase) *ProfileHandler {
	return &ProfileHandler{usecases: usecase}
}

func (h *ProfileHandler) CreateProfile(w http.ResponseWriter, r *http.Request) {
	url, err := utils.GetImageUrl(r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	user_id, err := utils.GetID(r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	profile := &domain.Profile{User_ID: user_id, ImageURL: url}
	pr, err := h.usecases.CreateProfile(profile)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, pr)

}

func (h *ProfileHandler) GetProfileByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetID(r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	pr, err := h.usecases.GetProfileByID(id)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, pr)
}

func (h *ProfileHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	url, err := utils.GetImageUrl(r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	var payload types.UpdatePayload
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	up := &domain.Profile{ID: payload.ID, ImageURL: url}
	err = h.usecases.UpdateProfile(up)
	if err != nil {
		utils.WriteError(w, err)
	}
	utils.WriteJSON(w, http.StatusOK, up)

}

func (h *ProfileHandler) DeleteProfile(w http.ResponseWriter, r *http.Request) {

	id, err := utils.GetID(r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	pr, err := h.usecases.GetProfileByID(id)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	err = utils.DeleteProfileFromCloud(r, pr.ImageURL)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	err = h.usecases.DeleteProfile(id)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, "Delete successfully")
}

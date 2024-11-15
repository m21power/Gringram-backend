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

func NewProfileHandler(usecase *usecases.ProfileUsecase) *ProfileHandler {
	return &ProfileHandler{usecases: usecase}
}

func (h *ProfileHandler) CreateProfile(w http.ResponseWriter, r *http.Request) {
	var payload types.ProfilePayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	profile := payloadToDomainProfile(&payload)
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
	var payload types.UpdateProfilePayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	up := updateProfile(&payload)
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
	err = h.usecases.DeleteProfile(id)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, "Delete successfully")
}

func payloadToDomainProfile(payload *types.ProfilePayload) *domain.Profile {
	return &domain.Profile{
		Url: payload.Url,
	}
}

func updateProfile(payload *types.UpdateProfilePayload) *domain.Profile {
	return &domain.Profile{
		ID:  payload.ID,
		Url: payload.Url,
	}
}

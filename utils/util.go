package utils

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/m21power/GrinGram/domain"
	"github.com/m21power/GrinGram/types"
)

func WriteJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func WriteError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	response := map[string]string{
		"error": err.Error(),
	}
	json.NewEncoder(w).Encode(response)
}

func PayloadToDomainUser(payload types.UserPayload) *domain.User {
	return &domain.User{
		Name:     payload.Name,
		Username: payload.Username,
		Password: payload.Password,
		Email:    payload.Email,
		Bio:      payload.Bio,
	}
}

func GetID(r *http.Request) (int, error) {
	ID := mux.Vars(r)["id"]
	id, err := strconv.Atoi(ID)
	if err != nil {
		return -1, err
	}
	return id, nil
}

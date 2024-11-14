package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	auth "github.com/m21power/GrinGram/Auth"
	"github.com/m21power/GrinGram/domain"
	"github.com/m21power/GrinGram/types"
	"github.com/m21power/GrinGram/usecases"
	"github.com/m21power/GrinGram/utils"
)

type UserHandler struct {
	usecase *usecases.UserUsecase
}

func NewUserHandler(usecase *usecases.UserUsecase) *UserHandler {
	return &UserHandler{
		usecase: usecase,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userPayload types.UserPayload
	err := json.NewDecoder(r.Body).Decode(&userPayload)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	user := utils.PayloadToDomainUser(userPayload)
	cu, err := h.usecase.CreateUser(user)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, cu)
}
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetID(r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	u, err := h.usecase.GetUserByID(id)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, u)

}

func (h *UserHandler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		utils.WriteError(w, nil)
		return

	}
	u, err := h.usecase.GetUserByEmail(email)
	if err != nil {
		utils.WriteError(w, err)
		return

	}
	utils.WriteJSON(w, http.StatusOK, u)

}
func (h *UserHandler) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		utils.WriteError(w, nil)
		return

	}
	u, err := h.usecase.GetUserByUsername(username)
	if err != nil {
		utils.WriteError(w, err)
		return

	}
	utils.WriteJSON(w, http.StatusOK, u)

}
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var userPayload types.UserPayload
	err := json.NewDecoder(r.Body).Decode(&userPayload)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	id, err := utils.GetID(r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	hashedPassword, err := auth.HashedPassword(userPayload.Password)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	userPayload.Password = hashedPassword
	oldUser, err := h.usecase.GetUserByID(id)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	user := utils.PayloadToDomainUser(userPayload)
	new := updateFunc(oldUser, user)
	err = h.usecase.UpdateUser(new)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, new)

}
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetID(r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	err = h.usecase.DeleteUser(id)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, "Deleted Succesffully!")

}

func updateFunc(oldUser *domain.User, newUser *domain.User) *domain.User {
	if newUser.Email != "" {
		oldUser.Email = newUser.Email
	}
	if newUser.Name != "" {
		oldUser.Name = newUser.Name
	}
	if newUser.Username != "" {
		oldUser.Username = newUser.Username
	}
	if newUser.Username != "" {
		oldUser.Password = newUser.Password
	}
	if newUser.Bio != "" {
		log.Println(newUser.Bio)
		oldUser.Bio = newUser.Bio
	}
	return oldUser

}

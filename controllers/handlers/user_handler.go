package handlers

import (
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
	ctx := r.Context()
	url, err := utils.GetProfileUrl(r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	userPayload, err := utils.GetUserPayload(w, r)
	if userPayload == nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"message": "Payload empty"})
		return
	}
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	user := &domain.User{Name: userPayload.Name, Username: userPayload.Username, Email: userPayload.Email, Password: userPayload.Password, Bio: userPayload.Bio, ProfileImageUrl: url}
	cu, err := h.usecase.CreateUser(ctx, user)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, cu)
}
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := utils.GetID(r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	u, err := h.usecase.GetUserByID(ctx, id)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, u)

}

func (h *UserHandler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	email := r.URL.Query().Get("email")
	if email == "" {
		utils.WriteError(w, nil)
		return

	}
	u, err := h.usecase.GetUserByEmail(ctx, email)
	if err != nil {
		utils.WriteError(w, err)
		return

	}
	utils.WriteJSON(w, http.StatusOK, u)

}
func (h *UserHandler) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	username := r.URL.Query().Get("username")
	if username == "" {
		utils.WriteError(w, nil)
		return

	}
	u, err := h.usecase.GetUserByUsername(ctx, username)
	if err != nil {
		utils.WriteError(w, err)
		return

	}
	utils.WriteJSON(w, http.StatusOK, u)

}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tx, err := h.usecase.BeginTransaction(ctx)
	defer func() { // this function runs at the end of the function before it returns, then before return it checks the value of err
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	id, err := utils.GetID(r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	url, err := h.usecase.GetProfileURL(ctx, tx, id)
	if err != nil {
		utils.WriteError(w, err)
		return
	}

	// now we have to delete first from the cloud
	err = utils.DeleteImageFromCloud(r, url)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	url, err = utils.GetProfileUrl(r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	userpy, err := utils.GetUserPayload(w, r)
	if userpy == nil {
		userpy = &types.UserPayload{}
	}
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	if userpy.Password != "" {
		userpy.Password, err = auth.HashedPassword(userpy.Password)
		if err != nil {
			utils.WriteError(w, err)
			return
		}
	}
	oldUser, err := h.usecase.GetUserByID(ctx, id)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	up := updateFunc(oldUser, *userpy, url)
	err = h.usecase.UpdateUser(ctx, up)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, up)

}
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := utils.GetID(r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	err = h.usecase.DeleteUser(ctx, id)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, "Deleted Succesffully!")

}

func (h *UserHandler) DeleteUserImage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tx, err := h.usecase.BeginTransaction(ctx)
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
	id, err := utils.GetID(r)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	url, err := h.usecase.GetProfileURL(ctx, tx, id)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	// now we have to delete first from the cloud
	err = utils.DeleteImageFromCloud(r, url)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	err = h.usecase.DeleteUserImage(ctx, tx, id)
	if err != nil {
		utils.WriteError(w, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Deleted successfully"})
}

func updateFunc(oldUser *domain.User, newUser types.UserPayload, image_url string) *domain.User {
	if newUser.Email != "" {
		oldUser.Email = newUser.Email
	}
	if newUser.Name != "" {
		oldUser.Name = newUser.Name
	}
	if newUser.Username != "" {
		oldUser.Username = newUser.Username
	}
	if newUser.Password != "" {
		oldUser.Password = newUser.Password
	}
	if newUser.Bio != "" {
		oldUser.Bio = newUser.Bio
	}
	if image_url != "" {
		oldUser.ProfileImageUrl = image_url
	}

	return oldUser

}

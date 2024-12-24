package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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

// @Summary User login
// @Description Authenticates a user and returns a JWT token in a cookie
// @Tags Login/Signup
// @Accept json
// @Produce json
// @Param login body domain.LoginPayload true "Login credentials"
// @Success 200 {object} utils.SuccessReponse
// @Failure 400 {object} utils.ErrorResponse "Bad request"
// @Router /login [post]
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var login domain.LoginPayload
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	token, err := h.usecase.Login(ctx, login)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token", // Cookie name
		Value:    token,   // Token as cookie value
		Path:     "/",     // Path for which the cookie is valid
		HttpOnly: true,    // Restrict access to HTTP(S) requests only
		Secure:   true,    // Ensure cookie is sent over HTTPS
		MaxAge:   3600 * 24 * 7,
		SameSite: http.SameSiteStrictMode, // Protect against CSRF
	})

	utils.WriteJSON(w, http.StatusOK, utils.ApiResponse{Message: "Login successful", Success: true})
}

// @Summary User logout
// @Description Delete token from a cookie
// @Tags Login/Signup
// @Accept json
// @Produce json
// @Success 200 {object} utils.SuccessReponse
// @Failure 400 {object} utils.ErrorResponse "Bad request"
// @Router /users/logout [post]
func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// Remove the cookie by setting it with an expiration in the past
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Unix(0, 0), // Set expiration to a time in the past
		MaxAge:   -1,              // MaxAge < 0 means delete the cookie
	})

	// Send a success response
	utils.WriteJSON(w, http.StatusOK, utils.ApiResponse{Message: "Logout successfull", Success: true})

}

// GetMe godoc
// @Summary Get current user information
// @Description Retrieves the information of the currently authenticated user based on the provided token.
// @Tags Users
// @Produce json
// @Success 200 {object} domain.User "User information"
// @Failure 400 {object} utils.ErrorResponse "Invalid request"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /users/me [get]
func (h *UserHandler) GetMe(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	token, err := auth.GetTokens(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	Token, err := auth.GetTokenValues(token)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	user, err := h.usecase.GetUserByID(ctx, Token.UserID)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	utils.WriteJSON(w, http.StatusOK, utils.ApiResponse{Message: "User found", Success: true, Data: user})
}

// @Summary Create a new user
// @Description Register a new user with email, username, and password
// @Tags Login/Signup
// @Accept  json
// @Produce  json
// @Param   user  body  types.UserPayload  true  "User registration info"
// @Success 201 {object} domain.User
// @Failure 400 {object} utils.ErrorResponse
// @Router /signup [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	url, err := utils.GetProfileUrl(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	userPayload, err := utils.GetUserPayload(w, r)
	if userPayload == nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: "User payload empty", Success: false})
		return
	}
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	user := &domain.User{Name: userPayload.Name, Username: userPayload.Username, Email: userPayload.Email, Password: userPayload.Password, Bio: userPayload.Bio, ProfileImageUrl: url}
	cu, err := h.usecase.CreateUser(ctx, user)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		utils.DeleteImageFromCloud(r, url)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, utils.ApiResponse{Message: "User created successfully", Success: true, Data: cu})
}

// @Summary Get a user by user ID
// @Description Get a user by their user ID, only admins can access it
// @Tags Users
// @Accept  json
// @Produce  json
// @Param   id  path  int  true  "User ID"
// @Success 200 {object} domain.User
// @Failure 400 {object} utils.ErrorResponse
// @Router /users/{id} [get]
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := utils.GetID(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	u, err := h.usecase.GetUserByID(ctx, id)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	utils.WriteJSON(w, http.StatusOK, utils.ApiResponse{Message: "User found", Success: true, Data: u})

}

// @Summary Get user by email
// @Description Retrieve a user by their email address from the database.
// @Tags Users
// @Produce json
// @Param email query string true "User Email"
// @Success 200 {object} domain.User
// @Failure 400 {object} utils.ErrorResponse
// @Router /users/email [get]
func (h *UserHandler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	email := r.URL.Query().Get("email")
	if email == "" {
		utils.WriteJSON(w, http.StatusNotFound, utils.ApiResponse{Message: "Email not found", Success: false})
		return

	}
	u, err := h.usecase.GetUserByEmail(ctx, email)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return

	}
	utils.WriteJSON(w, http.StatusOK, utils.ApiResponse{Message: "User found", Success: true, Data: u})

}

// @Summary Get user by username
// @Description Retrieve a user by their username from the database.
// @Tags Users
// @Produce json
// @Param username query string true "Username"
// @Success 200 {object} domain.User
// @Failure 400 {object} utils.ErrorResponse
// @Router /users/username [get]
func (h *UserHandler) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	username := r.URL.Query().Get("username")
	if username == "" {
		utils.WriteJSON(w, http.StatusNotFound, utils.ApiResponse{Message: "Username not found", Success: false})
		return

	}
	u, err := h.usecase.GetUserByUsername(ctx, username)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return

	}
	utils.WriteJSON(w, http.StatusOK, utils.ApiResponse{Message: "user found", Success: true, Data: u})

}

// @Summary Update a user's profile
// @Description Updates the user's profile information including password and profile picture.
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body types.UserPayload true "User Payload"
// @Param image_url query string true "Profile picture URL"
// @Success 200 {object} domain.User "Updated user information"
// @Failure 400 {object} utils.ErrorResponse "Bad request"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 404 {object} utils.ErrorResponse "User not found"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /users/{id} [put]
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
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	_, err = IsAllowed(r, id)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	url, err := h.usecase.GetProfileURL(ctx, tx, id)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}

	// now we have to delete first from the cloud
	err = utils.DeleteImageFromCloud(r, url)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	url, err = utils.GetProfileUrl(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	userpy, err := utils.GetUserPayload(w, r)
	if userpy == nil {
		userpy = &types.UserPayload{}
	}
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	if userpy.Password != "" {
		userpy.Password, err = auth.HashedPassword(userpy.Password)
		if err != nil {
			utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
			return
		}
	}
	oldUser, err := h.usecase.GetUserByID(ctx, id)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	up := updateFunc(oldUser, *userpy, url)
	err = h.usecase.UpdateUser(ctx, up)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	utils.WriteJSON(w, http.StatusOK, utils.ApiResponse{Message: "User updated successfully", Success: true, Data: up})

}

// @Summary Delete a user
// @Description Delete a user by their user ID, only admins or the user themselves can delete their account
// @Tags Users
// @Accept  json
// @Produce  json
// @Param   id  path  int  true  "User ID"
// @Success 200 {object} utils.DeleteResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /users/delete/{id} [delete]
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := utils.GetID(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	_, err = IsAllowed(r, id)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	err = h.usecase.DeleteUser(ctx, id)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	utils.WriteJSON(w, http.StatusOK, utils.ApiResponse{Message: "User deleted successfully", Success: true})
}

// DeleteUserImage godoc
// @Summary Delete user profile image
// @Description Deletes the user's profile image from the cloud and the database
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object}  utils.DeleteResponse
// @Failure 400 {object} utils.ErrorResponse "Invalid request"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /users/image/{id} [delete]
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
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	_, err = IsAllowed(r, id)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	url, err := h.usecase.GetProfileURL(ctx, tx, id)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	// now we have to delete first from the cloud
	err = utils.DeleteImageFromCloud(r, url)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	err = h.usecase.DeleteUserImage(ctx, tx, id)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
		return
	}
	utils.WriteJSON(w, http.StatusOK, utils.ApiResponse{Message: "Image deleted successfully", Success: true})
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

func IsAllowed(r *http.Request, userId int) (bool, error) {
	token, err := auth.GetTokens(r)
	if err != nil {
		return false, err
	}
	Token, err := auth.GetTokenValues(token)
	if err != nil {
		return false, err
	}
	if Token.Role == "user" && Token.UserID != userId {
		return false, fmt.Errorf("you are not allowed")
	}
	return true, nil
}

package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
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
	url, err := getImageUrl(r)
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
	url, err := getImageUrl(r)
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
	err = deleteProfileFromCloud(r, pr.ImageURL)
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

func getImageUrl(r *http.Request) (string, error) {
	cld, err := cloudinary.NewFromParams("dl6vahv6t", "639632577282947", "_qyu3umAppkfaRNR84QUuWiIa7U")
	if err != nil {
		return "", err
	}
	// Parse form data, including file
	err = r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		return "", err
	}
	// Get the file from the form data
	file, _, err := r.FormFile("image") // "image" is the key in the form
	if err != nil {
		return "", err
	}
	defer file.Close()
	// Upload image to Cloudinary
	// We will upload the file directly from the reader (file content)
	resp, err := cld.Upload.Upload(r.Context(), file, uploader.UploadParams{
		Folder: "images",
	})

	if err != nil {
		return "", err
	}
	return resp.URL, nil

}

func deleteProfileFromCloud(r *http.Request, link string) error {
	cld, err := cloudinary.NewFromParams("dl6vahv6t", "639632577282947", "_qyu3umAppkfaRNR84QUuWiIa7U")
	if err != nil {
		return err
	}
	publicID, err := extractPublicID(link)
	if err != nil {
		return err
	}
	// Delete the image
	resp, err := cld.Upload.Destroy(r.Context(), uploader.DestroyParams{
		PublicID: publicID,
	})
	if err != nil {
		return err
	}

	// Check the response and confirm if the image was deleted successfully
	if resp.Result == "ok" {
		return nil
	} else {
		log.Println("it is here right ")
		return fmt.Errorf("failed")
	}
}
func extractPublicID(cloudinaryURL string) (string, error) {
	// Split the URL by '/'
	parts := strings.Split(cloudinaryURL, "/")
	// Check if the URL is valid and has the necessary parts
	if len(parts) < 9 {
		return "", fmt.Errorf("invalid link")
	}
	// The public ID is located at position 8 (index 8) after /image/upload/v<version>/<public-id>.<format>
	// Example URL structure: https://res.cloudinary.com/<cloud-name>/image/upload/v<version>/<public-id>.<format>

	publicIDWithExtension := parts[8]                        // This will give the public ID with file extension (e.g., 'image.jpg')
	publicID := strings.Split(publicIDWithExtension, ".")[0] // Remove the file extension

	return publicID, nil
}

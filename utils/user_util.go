package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/m21power/GrinGram/types"
)

func GetProfileUrl(r *http.Request) (string, error) {
	cld, err := cloudinary.NewFromParams("dl6vahv6t", "639632577282947", "_qyu3umAppkfaRNR84QUuWiIa7U")
	if err != nil {
		return "", err
	}
	if r.ContentLength == 0 {
		return "", nil
	}
	// Parse form data, including file

	err = r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		return "", err
	}
	// Get the file from the form data
	file, _, err := r.FormFile("profile") // "profile" is the key in the form
	if file == nil {
		return "", nil
	}
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

func DeleteImageFromCloud(r *http.Request, link string) error {
	if link == "" {
		return nil
	}
	cld, err := cloudinary.NewFromParams("dl6vahv6t", "639632577282947", "_qyu3umAppkfaRNR84QUuWiIa7U")
	if err != nil {
		return err
	}

	// The public ID of the image to delete
	publicID, err := ExtractPublicID(link)
	if err != nil {
		return err
	}

	// Use the Admin API to delete the image
	resp, err := cld.Admin.DeleteAssets(r.Context(), admin.DeleteAssetsParams{
		PublicIDs: []string{publicID},
	})
	if err != nil {
		return err
	}
	if resp.Deleted[publicID] == "not_found" || resp.Deleted[publicID] == "deleted" {
		return nil
	}
	return fmt.Errorf("image deleting unsuccessfull")

}
func ExtractPublicID(cloudinaryURL string) (string, error) {
	// Split the URL by '/'
	parts := strings.Split(cloudinaryURL, "/")
	// Check if the URL is valid and has the necessary parts
	if len(parts) < 9 {
		return "", fmt.Errorf("invalid link")
	}
	publicIDWithExtension := parts[8]                                    // This will give the public ID with file extension (e.g., 'image.jpg')
	publicID := "images/" + strings.Split(publicIDWithExtension, ".")[0] // Remove the file extension

	return publicID, nil
}

func GetUserPayload(w http.ResponseWriter, r *http.Request) (*types.UserPayload, error) {
	if r.ContentLength == 0 {
		return nil, nil
	}
	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB max memory
	if err != nil {
		return nil, err
	}

	// Access the `data` part
	jsonData := r.FormValue("UserPayload") // Retrieve the `data` field from the form
	if jsonData == "" {
		return nil, err
	}

	// Unmarshal JSON into the struct
	var userPayload types.UserPayload
	err = json.Unmarshal([]byte(jsonData), &userPayload)
	if err != nil {
		return nil, err
	}
	return &userPayload, nil
}
func GetPayload(w http.ResponseWriter, r *http.Request, payload any) (any, error) {
	if r.ContentLength == 0 {
		return nil, nil
	}
	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB max memory
	if err != nil {
		return nil, err
	}

	// Access the `data` part
	jsonData := r.FormValue("DataPayload") // Retrieve the `data` field from the form
	if jsonData == "" {
		return nil, err
	}

	// Unmarshal JSON into the struct
	err = json.Unmarshal([]byte(jsonData), &payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}
func GetPostImagesURL(r *http.Request) ([]string, error) {
	cld, err := cloudinary.NewFromParams("dl6vahv6t", "639632577282947", "_qyu3umAppkfaRNR84QUuWiIa7U")
	if err != nil {
		return nil, err
	}
	if r.ContentLength == 0 {
		return nil, nil
	}
	// Parse form data, including file

	err = r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		return nil, err
	}
	// Get the file from the form data
	files := r.MultipartForm.File["PostImage"] // "profile" is the key in the form
	if files == nil {
		return nil, nil
	}
	var Urls []string
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			return nil, err
		}
		// Upload image to Cloudinary
		// We will upload the file directly from the reader (file content)
		resp, err := cld.Upload.Upload(r.Context(), file, uploader.UploadParams{
			Folder: "images",
		})
		defer file.Close()
		if err != nil {
			return nil, err
		}
		Urls = append(Urls, resp.URL)
	}
	return Urls, nil

}
func GetPostImageURL(r *http.Request) (string, error) {
	cld, err := cloudinary.NewFromParams("dl6vahv6t", "639632577282947", "_qyu3umAppkfaRNR84QUuWiIa7U")
	if err != nil {
		return "", err
	}
	if r.ContentLength == 0 {
		return "", nil
	}
	// Parse form data, including file

	err = r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		return "", err
	}
	// Get the file from the form data
	file, _, err := r.FormFile("PostImage") // "profile" is the key in the form
	if file == nil {
		return "", nil
	}
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

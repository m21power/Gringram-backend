package utils

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func GetImageUrl(r *http.Request) (string, error) {
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

func DeleteProfileFromCloud(r *http.Request, link string) error {
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

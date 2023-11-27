package media

import (
	"image"

	"github.com/Azpect3120/MediaStorageServer/internal/models"
)

// Compares two images to see if they are the same
func Compare(imgOne, imgTwo image.Image) bool {
	return imgOne.Bounds().Eq(imgTwo.Bounds())
}

// Compares an image to an array of images and returns the FIRST match
func CompareArray(target image.Image, images []*models.Image) (*models.Image, bool, error) {
	for _, img := range images {
		file, err := OpenImage("." + img.Path)
		if err != nil {
			return nil, false, err
		}

		if exists := Compare(target, file); exists {
			return img, true, nil
		}
	}
	return nil, false, nil
}

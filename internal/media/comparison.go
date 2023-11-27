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

// Compare a video to an array of videos and returns FIRST match
func CompareArrayVideos (target *models.Image, videos []*models.Image) (*models.Image, bool, error) {
	targetHash, err := CalculateHash("." + target.Path)
	if err != nil {
		return nil, false, err
	}
	
	for _, vid := range videos {
		hash, err := CalculateHash("." + vid.Path)
		if err != nil {
			return nil, false, err
		}

		if hash == targetHash {
			return vid, true, nil
		}
	}
	return nil, false, nil
}

package media

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

// Open an image as an 'image.Image' object
func OpenImage(path string) (image.Image, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		file.Seek(0, 0)
		img, err = png.Decode(file)
		if err != nil {
			return nil, err
		}
	}

	return img, nil
}

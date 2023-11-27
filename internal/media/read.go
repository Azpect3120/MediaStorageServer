package media

import (
	"crypto/md5"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"golang.org/x/image/webp"
	"io"
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
			// file2, err := os.Open(path)
			// if err != nil {
			// 	return nil, err
			// }
			// defer file2.Close()
			file.Seek(0, 0)
			img, err = webp.Decode(file)
			if err != nil {
				fmt.Printf("%+v\n", err)
				return nil, err
			}
		}
	}

	return img, nil
}

// Calculates a videos hash
func CalculateHash (path string) (string, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

package media

import (
	"crypto/md5"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"image/gif"
	"golang.org/x/image/webp"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	"github.com/mat/besticon/ico"
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
			file.Seek(0, 0)
			img, err = webp.Decode(file)
			if err != nil {
				file.Seek(0, 0)
				img, err = gif.Decode(file)
				if err != nil {
					file.Seek(0, 0)
					img, err = ico.Decode(file)
					if err != nil {
						file.Seek(0, 0)
						img, err = tiff.Decode(file)
						if err != nil {
							file.Seek(0, 0)
							img, err = bmp.Decode(file)
							if err != nil {
								return nil, err
							}
						}
					}
				}
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

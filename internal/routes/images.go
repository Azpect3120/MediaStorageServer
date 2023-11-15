package routes

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Azpect3120/MediaStorageServer/internal/database"
	"github.com/Azpect3120/MediaStorageServer/internal/models"
	"github.com/gin-gonic/gin"
)

// Creates a image
//
// File upload must have the name 'media_upload'
func CreateImage (db *database.Database, root string, ctx *gin.Context) {
	folderId := ctx.Param("id")

	// Get file from request
	file, err := ctx.FormFile("media_upload")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "status": http.StatusInternalServerError, "error": err.Error() })
		return
	}

	// Create image object with values 
	image := &models.Image{
		Name: file.Filename,
		FolderId: folderId,
		Size: file.Size,
		Format: file.Header.Get("Content-Type"),
	}

	// Create image in database: updates image object
	if err = db.CreateImage(image); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "status": http.StatusInternalServerError, "error": err.Error() })
		return
	}

	// Join path to abs path and get extension
	fullPath := filepath.Join(root, image.FolderId, image.ID) + filepath.Ext(image.Name)
	image.Path = filepath.Join("uploads", image.FolderId, image.ID) + filepath.Ext(image.Name)

	// Create image on file system using newly updated image object
	outputFile, err := os.Create(fullPath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "status": http.StatusInternalServerError, "error": err.Error() })
		return
	}

	defer outputFile.Close()

	// Save image to the file
	if err := ctx.SaveUploadedFile(file, outputFile.Name()); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "status": http.StatusInternalServerError, "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "image": image })
}

// Gets a image
func GetImage (db *database.Database, root string, ctx *gin.Context) {
	id := ctx.Param("id")

	image, err := db.GetImage(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "status": http.StatusInternalServerError, "error": err.Error() })
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{ "status": http.StatusOK, "image": image })
}

// Deletes a image
func DeleteImage (db *database.Database, root string, ctx *gin.Context) {
	id := ctx.Param("id")

	// Delete image from database
	if err := db.DeleteImage(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "status": http.StatusInternalServerError, "error": err.Error() })
		return
	}

	// Delete image from file system
	err := deleteFile(root, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "status": http.StatusInternalServerError, "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// Walk the file path and delete a file using its ID
func deleteFile (root, id string) error {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Checks if the file is not a directory and has the specified name
		if !info.IsDir() && strings.HasPrefix(info.Name(), id) {
			ext := filepath.Ext(info.Name())

			// Check again if the name matches the current file without its extension
			if strings.TrimSuffix(info.Name(), ext) == id {
				if err := os.Remove(path); err != nil {
					return err
				}
			}
		}
		return nil
	})
	return err	
}

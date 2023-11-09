package routes

import (
	"net/http"
	"os"
	"path/filepath"

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

// Updates a image
func UpdateImage (db *database.Database, root string, ctx *gin.Context) {

}

// Deletes a image
func DeleteImage (db *database.Database, root string, ctx *gin.Context) {

}

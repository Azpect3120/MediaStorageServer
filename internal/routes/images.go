package routes

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Azpect3120/MediaStorageServer/internal/cache"
	"github.com/Azpect3120/MediaStorageServer/internal/database"
	"github.com/Azpect3120/MediaStorageServer/internal/models"
	"github.com/gin-gonic/gin"
)

// Creates a image
//
// File upload must have the name 'media_upload'
func CreateImage (cache *cache.Cache, db *database.Database, root string, ctx *gin.Context) {
	folderId := ctx.Param("id")

	cache.ResetRequest("/folders/" + folderId)

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
	ch := make(chan error)
	go db.CreateImage(ch, image)
	err = <- ch

	if err != nil {
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
func GetImage (folderCache, imageCache *cache.Cache, db *database.Database, root string, ctx *gin.Context) {
	// Check cache for request
	request := ctx.Request.URL.String()

	if response, exists := imageCache.GetResponse(request); exists {
		var data GetImageResponse
		if err := json.Unmarshal(response, &data); err == nil {
			ctx.JSON(http.StatusOK, data)
			return
		}
	}

	id := ctx.Param("id")

	ch := make(chan models.ImageChannel)
	go db.GetImage(ch, id)
	res := <- ch

	if res.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "status": http.StatusInternalServerError, "error": res.Error.Error() })
		return
	}

	// Add response and request to cache
	response := GetImageResponse{ Status: http.StatusOK, Image: res.Image }
	responseData, err := json.Marshal(response)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "status": http.StatusBadRequest, "error": err.Error() })
		return
	}
	imageCache.AddResponse(request, responseData)

	// Remove the parent folder from the cache 
	folderCache.ResetRequest("/folders/" + res.Image.FolderId)
	
	ctx.JSON(http.StatusOK, gin.H{ "status": http.StatusOK, "image": res.Image })
}

// Deletes a image
func DeleteImage (folderCache, imageCache *cache.Cache, db *database.Database, root string, ctx *gin.Context) {
	// Remove image from cache
	imageCache.ResetRequest(ctx.Request.URL.String())

	id := ctx.Param("id")

	chFID := make(chan models.IDChannel)
	go db.GetFolderID(chFID, id)
	res := <- chFID

	if res.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "status": http.StatusInternalServerError, "error": res.Error.Error() })
		return
	}

	// Remove parent folder from cache
	folderCache.ResetRequest("/folders/" + res.ID)

	// Delete image from database
	ch := make(chan error)
	go db.DeleteImage(ch, id)
	err := <- ch

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "status": http.StatusInternalServerError, "error": err.Error() })
		return
	}

	// Delete image from file system
	chFS := make(chan error)
	go deleteFile(chFS, root, id)
	err = <- chFS

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "status": http.StatusInternalServerError, "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// Walk the file path and delete a file using its ID
func deleteFile (ch chan error, root, id string) {
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
	ch <- err
}

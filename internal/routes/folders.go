package routes

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/Azpect3120/MediaStorageServer/internal/cache"
	"github.com/Azpect3120/MediaStorageServer/internal/database"
	"github.com/Azpect3120/MediaStorageServer/internal/models"
	"github.com/gin-gonic/gin"
)

// Creates a folder
func CreateFolder(db *database.Database, root string, ctx *gin.Context) {
	var req CreateFolderRequest

	// Bind request to struct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
		return
	}

	// Valid folder name regex
	req.Name = strings.TrimSpace(req.Name)
	result := regexp.MustCompile(`\s+`).ReplaceAllString(req.Name, "_")
	validDirName := regexp.MustCompile("^[a-zA-Z0-9_\\-]+$")

	// Ensure folder name is valid
	if validDirName.MatchString(result) {
		req.Name = result
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": "Invalid folder name. Folder name must only contain letters, digits, underscores, and hyphens."})
		return
	}

	// Create folder in database
	folder, err := db.CreateFolder(req.Name)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
		return
	}

	// Create folder in file system
	newFolderPath := filepath.Join(root, folder.ID)

	err = os.Mkdir(newFolderPath, 0755)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
	}

	// Return newly created folder
	ctx.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "folder": folder})
}

// Gets a folder
func GetFolder(cache *cache.Cache, db *database.Database, root string, ctx *gin.Context) {
	// Check cache for request
	request := ctx.Request.URL.String()

	if response, exists := cache.GetResponse(request); exists {
		var data GetFolderResponse
		if err := json.Unmarshal(response, &data); err == nil {
			ctx.JSON(http.StatusOK, data)
			return
		}
	}

	id := ctx.Param("id")

	// Validate id
	if valid := ValidateID(id); !valid {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": "Please enter a valid id."})
		return
	}

	// Get folder meta data from database
	folder, err := db.GetFolder(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
		return
	}
	// Add response and request to cache
	response := GetFolderResponse{Status: http.StatusOK, Folder: folder}
	responseData, err := json.Marshal(response)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
		return
	}
	cache.AddResponse(request, responseData)

	ctx.JSON(http.StatusOK, response)
}

// Gets a list of the images in a folder
func GetFolderImages(cache *cache.Cache, db *database.Database, root string, ctx *gin.Context) {
	// Check cache for request
	request := ctx.Request.URL.String()

	if response, exists := cache.GetResponse(request); exists {
		var data GetFolderImagesResponse
		if err := json.Unmarshal(response, &data); err == nil {
			ctx.JSON(http.StatusOK, data)
			return
		}
	}

	// Get folder id
	id := ctx.Param("id")

	// Get url queries
	limit := ctx.Query("limit")
	page := ctx.Query("page")

	// Validate URL inputs
	if !ValidateID(id) {
		ctx.JSON(http.StatusBadRequest, gin.H{ "status": http.StatusBadRequest, "error": "Please provide a valid id." })
		return
	}
	if page == "" { page = "1" }
	if limit == "" { limit = "100" }

	// Convert strings to integers
	limitNum, err := strconv.Atoi(limit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "status": http.StatusBadRequest, "error": "Please provide a valid limit." })
		return
	}

	pageNum, err := strconv.Atoi(page)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "status": http.StatusBadRequest, "error": "Please provide a valid page." })
		return
	}

	// Calculate offset for SQL
	pageNum = (pageNum - 1) * limitNum

	// Get images from the database
	images, err := db.GetImages(id, limitNum, pageNum)

	if images == nil {
		images = []*models.Image{}
	}

	// Add request to cache
	response := GetFolderImagesResponse{ Status: http.StatusOK, Images: images }
	responseData, err := json.Marshal(response)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err.Error()})
		return
	}
	cache.AddResponse(request, responseData)

	ctx.JSON(200, response)
	return
}

// Updates a folder
func UpdateFolder(cache *cache.Cache, db *database.Database, root string, ctx *gin.Context) {
	cache.ResetRequest(ctx.Request.URL.String())

	id := ctx.Param("id")

	// Validate id
	if valid := ValidateID(id); !valid {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": "Please enter a valid id."})
		return
	}

	folder := &models.Folder{}

	if err := ctx.ShouldBindJSON(&folder); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
		return
	}

	// Update folder in database
	err := db.UpdateFolder(id, folder)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
		return
	}

	// Get updated folder from database
	folder, err = db.GetFolder(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "folder": folder})
}

// Deletes a folder
func DeleteFolder(folderCache, imageCache *cache.Cache, db *database.Database, root string, ctx *gin.Context) {
	folderCache.ResetRequest(ctx.Request.URL.String())

	// Clear entire image cache
	imageCache.Clear()

	id := ctx.Param("id")

	// Validate id
	if valid := ValidateID(id); !valid {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": "Please enter a valid id."})
		return
	}

	// Delete folder from database
	err := db.DeleteFolder(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
		return
	}

	// Delete folder from file system
	targetPath := filepath.Join(root, id)

	err = os.RemoveAll(targetPath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

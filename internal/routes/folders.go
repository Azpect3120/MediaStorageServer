package routes

import (
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/Azpect3120/MediaStorageServer/internal/database"
	"github.com/Azpect3120/MediaStorageServer/internal/models"
	"github.com/gin-gonic/gin"
)

// Creates a folder
func CreateFolder (db *database.Database, root string, ctx *gin.Context) {
	var req CreateFolderRequest

	// Bind request to struct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "status": http.StatusBadRequest, "error": err.Error() })
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
		ctx.JSON(http.StatusBadRequest, gin.H{ "status": http.StatusBadRequest, "error": "Invalid folder name. Folder name must only contain letters, digits, underscores, and hyphens." })
		return
	}

	// Create folder in database
	ch := make(chan models.FolderChannel)
	go db.CreateFolder(ch, req.Name)
	res := <- ch

	if res.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "status": http.StatusInternalServerError, "error": res.Error.Error() })
		return
	}

	// Create folder in file system
	newFolderPath := filepath.Join(root, res.Folder.ID)

	err := os.Mkdir(newFolderPath, 0755)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "status": http.StatusBadRequest, "error": err.Error() })
	}

	// Return newly created folder 
	ctx.JSON(http.StatusCreated, gin.H{ "status": http.StatusCreated, "folder": res.Folder })
}

// Gets a folder
func GetFolder (db *database.Database, root string, ctx *gin.Context) {
	id := ctx.Param("id")
	
	// Get folder meta data from database
	ch := make(chan models.FolderChannel)
	go db.GetFolder(ch, id)
	res := <- ch

	if res.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "status": http.StatusBadRequest, "error": res.Error.Error() })
		return
	}

	// Get images from the database
	chImg := make(chan models.ImagesChannel)
	go db.GetImages(chImg, res.Folder.ID)
	imgRes := <- chImg

	if imgRes.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "status": http.StatusBadRequest, "error": imgRes.Error.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "status": http.StatusOK, "folder": res.Folder, "images": imgRes.Images, "count": len(imgRes.Images) })
}

// Updates a folder
func UpdateFolder (db *database.Database, root string, ctx *gin.Context) {
	id := ctx.Param("id")

	folder := &models.Folder{}

	if err := ctx.ShouldBindJSON(&folder); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "status": http.StatusInternalServerError, "error": err.Error() })
		return
	}

	// Update folder in database
	ch := make(chan error)
	go db.UpdateFolder(ch, id, folder)
	err := <- ch
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "status": http.StatusInternalServerError, "error": err.Error() })
		return
	}

	// Get updated folder from database
	chF := make(chan models.FolderChannel)
	go db.GetFolder(chF, id)
	res := <- chF

	if res.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "status": http.StatusInternalServerError, "error": res.Error.Error() })
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{ "status": http.StatusOK, "folder": res.Folder })
}

// Deletes a folder
func DeleteFolder (db *database.Database, root string, ctx *gin.Context) {
	id := ctx.Param("id")

	// Delete folder from database
	ch := make(chan error)
	go db.DeleteFolder(ch, id)
	err := <- ch

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "status": http.StatusInternalServerError, "error": err.Error() })
		return
	}

	// Delete folder from file system
	targetPath := filepath.Join(root, id)

	err = os.RemoveAll(targetPath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "status": http.StatusInternalServerError, "error": err.Error() })
		return
	}
	
	ctx.JSON(http.StatusNoContent, nil)
}

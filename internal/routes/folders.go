package routes

import (
	"os"
	"net/http"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/Azpect3120/MediaStorageServer/internal/database"
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
	folder, err := db.CreateFolder(req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "status": http.StatusInternalServerError, "error": err.Error() })
		return
	}

	// Create folder in file system
	newFolderPath := filepath.Join(root, folder.Name)

	err = os.Mkdir(newFolderPath, 0755)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "status": http.StatusBadRequest, "error": err.Error() })
	}

	// Return newly created folder 
	ctx.JSON(http.StatusCreated, gin.H{ "status": http.StatusCreated, "folder": folder })
}

// Gets a folder
func GetFolder (db *database.Database, root string, ctx *gin.Context) {

}

// Updates a folder
func UpdateFolder (db *database.Database, root string, ctx *gin.Context) {

}

// Deletes a folder
func DeleteFolder (db *database.Database, root string, ctx *gin.Context) {

}

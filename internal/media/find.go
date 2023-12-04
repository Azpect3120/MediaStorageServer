package media

import (
	"github.com/Azpect3120/MediaStorageServer/internal/database"
	"github.com/Azpect3120/MediaStorageServer/internal/models"
)

// Return an array of images that are of the same size of the 'target'
func FindMatches (db *database.Database, target *models.Image) ([]*models.Image,error) {
	imgs, err := db.GetImageMatches(target.Size, target.FolderId, target.ID)

	if err != nil {
		return nil, err
	}

	return imgs, nil
}

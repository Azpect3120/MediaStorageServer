package media

import (
	"github.com/Azpect3120/MediaStorageServer/internal/database"
	"github.com/Azpect3120/MediaStorageServer/internal/models"
)

// Return an array of images that are of the same size of the 'target'
func FindMatches (db *database.Database, target *models.Image) ([]*models.Image,error) {
	ch := make(chan models.ImagesChannel)
	go db.GetImageMatches(ch, target.Size, target.FolderId, target.ID)
	res := <- ch

	if res.Error != nil {
		return nil, res.Error
	}

	return res.Images, nil
}

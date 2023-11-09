package database

import (
	"github.com/Azpect3120/MediaStorageServer/internal/models"
	"path/filepath"
)

func (db *Database) GetFolder (id string) (*models.Folder, error) {
	var statement string = "SELECT * FROM folders WHERE id = $1;"

	var folder models.Folder

	if err := db.database.QueryRow(statement, id).Scan(&folder.ID, &folder.Name, &folder.CreatedAt); err != nil {
		return nil, err
	}

	return &folder, nil
}

func (db *Database) GetImage (id string) (*models.Image, error) {
	var statement string = "SELECT * FROM images WHERE id = $1;"

	var image models.Image 

	if err := db.database.QueryRow(statement, id).Scan(&image.ID, &image.FolderId, &image.Name, &image.Size, &image.Format, &image.UploadedAt); err != nil {
		return nil, err
	}

	image.Path = filepath.Join("uploads", image.FolderId, image.ID) + filepath.Ext(image.Name)

	return &image, nil
}

package database

import (
	"github.com/Azpect3120/MediaStorageServer/internal/models"
	"path/filepath"
)

func (db *Database) GetFolder (ch chan models.FolderChannel, id string) {
	var statement string = "SELECT * FROM folders WHERE id = $1;"

	var folder models.Folder

	if err := db.database.QueryRow(statement, id).Scan(&folder.ID, &folder.Name, &folder.CreatedAt); err != nil {
		ch <- models.FolderChannel{ Folder: nil, Error: err }
	}

	ch <- models.FolderChannel{ Folder: &folder, Error: nil }
}

func (db *Database) GetImage (ch chan models.ImageChannel, id string) {
	var statement string = "SELECT * FROM images WHERE id = $1;"

	var image models.Image 

	if err := db.database.QueryRow(statement, id).Scan(&image.ID, &image.FolderId, &image.Name, &image.Size, &image.Format, &image.UploadedAt); err != nil {
		ch <- models.ImageChannel{ Image: nil, Error: err }
	}

	image.Path = filepath.Join("uploads", image.FolderId, image.ID) + filepath.Ext(image.Name)

	ch <- models.ImageChannel{ Image: &image, Error: nil }
}

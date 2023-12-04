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

	if err := db.database.QueryRow(statement, id).Scan(&image.ID, &image.FolderId, &image.Name, &image.Size, &image.Format, &image.UploadedAt, &image.Path); err != nil {
		return nil, err
	}

	image.Path = filepath.Join("uploads", image.FolderId, image.ID) + filepath.Ext(image.Name)

	return &image, nil
}

func (db *Database) GetImages (id string, limit, page int) ([]*models.Image, error) {
	var statement string = "SELECT * FROM images WHERE folderid = $1 ORDER BY uploadedat DESC LIMIT $2 OFFSET $3;"
	
	var images []*models.Image

	rows, err := db.database.Query(statement, id, limit, page);
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var image models.Image

		if err := rows.Scan(&image.ID, &image.FolderId, &image.Name, &image.Size, &image.Format, &image.UploadedAt, &image.Path); err != nil {
			return nil, err
		}

		// image.Path = filepath.Join("uploads", image.FolderId, image.ID) + filepath.Ext(image.Name)
		images = append(images, &image)
	}

	return images, nil
}

func (db *Database) GetFolderID (id string) (string, error) {
	var statement string = "SELECT folderid FROM images WHERE id = $1"

	var folderID string
	if err := db.database.QueryRow(statement, id).Scan(&folderID); err != nil {
		return "", err
	}

	return folderID, nil
}

func (db *Database) GetImageMatches (size int64, folderID, id string) ([]*models.Image, error) {
	var statement string = "SELECT * FROM images WHERE size = $1 AND folderid = $2 AND id != $3 ORDER BY uploadedat;"

	rows, err := db.database.Query(statement, size, folderID, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var images []*models.Image

	for rows.Next() {
		var image models.Image

		if err := rows.Scan(&image.ID, &image.FolderId, &image.Name, &image.Size, &image.Format, &image.UploadedAt, &image.Path); err != nil {
			return nil, err
		}

		// image.Path = filepath.Join("uploads", image.FolderId, image.ID) + filepath.Ext(image.Name)
		images = append(images, &image)
	}

	return images, nil
}

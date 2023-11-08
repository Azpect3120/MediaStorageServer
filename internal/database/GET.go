package database

import "github.com/Azpect3120/MediaStorageServer/internal/models"

func (db *Database) GetFolder (id string) (*models.Folder, error) {
	var statement string = "SELECT * FROM folders WHERE id = $1;"

	var folder models.Folder

	if err := db.database.QueryRow(statement, id).Scan(&folder.ID, &folder.Name, &folder.CreatedAt); err != nil {
		return nil, err
	}

	return &folder, nil
}

func (db *Database) GetImage (id string) (*models.Image, error) {
	return nil, nil
}

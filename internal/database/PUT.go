package database

import (
	"errors"

	"github.com/Azpect3120/MediaStorageServer/internal/models"
)

func (db *Database) UpdateFolder (id string, updated *models.Folder) error {
	var statement string = "UPDATE folders SET name = $1 WHERE id = $2;"

	result, err := db.database.Exec(statement, updated.Name, id)
	if err != nil {
		return err
	}

	numRows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if numRows == 0 {
		return errors.New("Folder with ID " +  id + " not found")
	}

	return nil
}

func (db *Database) UpdateImage (id string, update *models.Image) (*models.Image, error) {

	return nil, nil
}

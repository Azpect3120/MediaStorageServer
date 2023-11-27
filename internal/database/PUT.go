package database

import (
	"errors"

	"github.com/Azpect3120/MediaStorageServer/internal/models"
)

func (db *Database) UpdateFolder (ch chan error, id string, updated *models.Folder) {
	var statement string = "UPDATE folders SET name = $1 WHERE id = $2;"

	result, err := db.database.Exec(statement, updated.Name, id)
	if err != nil {
		ch <- err
		return
	}

	numRows, err := result.RowsAffected()
	if err != nil {
		ch <- err
		return
	}

	if numRows == 0 {
		ch <- errors.New("Folder with ID " +  id + " not found")
		return
	}

	ch <- nil
}

func (db *Database) UpdateImage (ch chan error, image *models.Image) {
	var statement string = "UPDATE images SET path = $1 WHERE id = $2;"

	result, err := db.database.Exec(statement, image.Path, image.ID)
	if err != nil {
		ch <- err
		return
	}

	numRows, err := result.RowsAffected()
	if err != nil {
		ch <- err
		return
	}

	if numRows == 0 {
		ch <- errors.New("Image not found")
		return
	}

	ch <- nil
}

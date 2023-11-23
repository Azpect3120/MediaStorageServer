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
	}

	numRows, err := result.RowsAffected()
	if err != nil {
		ch <- err
	}

	if numRows == 0 {
		ch <- errors.New("Folder with ID " +  id + " not found")
	}

	ch <- nil
}

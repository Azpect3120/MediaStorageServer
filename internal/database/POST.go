package database

import (
	"strings"
	"errors"

	"github.com/Azpect3120/MediaStorageServer/internal/models"
)

// Create a folder in the database
//
// Converts folder names to be lowercase
//
// Schema does not allow duplicate file names: an error will be thrown
func (db *Database) CreateFolder (name string) (*models.Folder, error) {
	name = strings.ToLower(name)

	// Insert into database
	var statement string = "INSERT INTO folders (name) VALUES ($1);"

	if _, err := db.database.Exec(statement, name); err != nil {
		return nil, err
	}

	// Query new data
	statement = "SELECT * FROM folders WHERE name = $1;"

	var folder models.Folder

	if err := db.database.QueryRow(statement, name).Scan(&folder.ID, &folder.Name, &folder.CreatedAt); err != nil {
		return nil, err
	}

	return &folder, nil
}

// Create an image in the database
//
// Converts image names to be lowercase
//
// Does not return a new image object. Updates the image object passed into the function.
func (db *Database) CreateImage (image *models.Image) error {
	image.Name = strings.ToLower(image.Name)

	// Insert into database
	var statement string = "INSERT INTO images (id, folderid, name, type, size, path) VALUES ($1, $2, $3, $4, $5,  $6);"

	result, err := db.database.Exec(statement, image.ID, image.FolderId, image.Name, image.Format, image.Size, image.Path)
	if err != nil {
		return err
	}

	numRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if numRows == 0 {
		return errors.New("New image was not created.")
	}

	return err
}

package database

import (
	"strings"
	"errors"

	"github.com/Azpect3120/MediaStorageServer/internal/models"
	"github.com/google/uuid"
)

// Create a folder in the database
//
// Converts folder names to be lowercase
//
// Schema does not allow duplicate file names: an error will be thrown
func (db *Database) CreateFolder (ch chan models.FolderChannel, name string) {
	name = strings.ToLower(name)

	// Insert into database
	var statement string = "INSERT INTO folders (name) VALUES ($1);"

	if _, err := db.database.Exec(statement, name); err != nil {
		ch <- models.FolderChannel{ Folder: nil, Error: err }
	}

	// Query new data
	statement = "SELECT * FROM folders WHERE name = $1;"

	var folder models.Folder

	if err := db.database.QueryRow(statement, name).Scan(&folder.ID, &folder.Name, &folder.CreatedAt); err != nil {
		ch <- models.FolderChannel{ Folder: nil, Error: err }
	}

	ch <- models.FolderChannel{ Folder: &folder, Error: nil }
}

// Create an image in the database
//
// Converts image names to be lowercase
//
// Does not return a new image object. Updates the image object passed into the function.
func (db *Database) CreateImage (ch chan error, image *models.Image) {
	image.Name = strings.ToLower(image.Name)

	image.ID = uuid.New().String()

	// Insert into database
	var statement string = "INSERT INTO images (id, folderid, name, type, size) VALUES ($1, $2, $3, $4, $5);"

	result, err := db.database.Exec(statement, image.ID, image.FolderId, image.Name, image.Format, image.Size)
	if err != nil {

		ch <- err 
	}

	numRows, err := result.RowsAffected()
	if err != nil {
		ch <- err
	}
	
	if numRows == 0 {
		ch <- errors.New("New image was not created.")
	}

	ch <- nil
}

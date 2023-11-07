package database

import (
	"github.com/Azpect3120/MediaStorageServer/internal/models"
	"strings"
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

func (db *Database) CreateImage (name string) (*models.Image, error) {
	return nil, nil
}

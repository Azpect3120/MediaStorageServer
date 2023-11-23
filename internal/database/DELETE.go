package database

import "errors"

// Deletes a folder from the database
//
// Returns an error if the ID cannot be found
func (db *Database) DeleteFolder (ch chan error, id string) {
	var statement string = "DELETE FROM folders WHERE id = $1"

	result, err := db.database.Exec(statement, id)
	if err != nil {
		ch <- err
	}

	numRows, err := result.RowsAffected()
	if err != nil {
		ch <- err
	}

	if numRows == 0 {
		ch <- errors.New("Folder with ID " + id +  " not found.")
	}

	ch <- nil
}

// Delete an image from the database
//
// Returns an error if the ID cannot be found
func (db *Database) DeleteImage (ch chan error, id string) {
	var statement string = "DELETE FROM images WHERE id = $1;"

	reuslt, err := db.database.Exec(statement, id)
	if err != nil {
		ch <- err
	}

	numRows, err := reuslt.RowsAffected()
	if err != nil {
		ch <- err
	}

	if numRows == 0 {
		ch <- errors.New("Image with the ID " + id + " not found.") 
	}

	ch <- nil
}

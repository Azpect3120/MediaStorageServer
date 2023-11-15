package database

import "errors"

// Deletes a folder from the database
//
// Returns an error if the ID cannot be found
func (db *Database) DeleteFolder (id string) error {
	var statement string = "DELETE FROM folders WHERE id = $1"

	result, err := db.database.Exec(statement, id)
	if err != nil {
		return err
	}

	numRows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if numRows == 0 {
		return errors.New("Folder with ID " + id +  " not found.")
	}

	return nil
}

// Delete an image from the database
//
// Returns an error if the ID cannot be found
func (db *Database) DeleteImage (id string) error {
	var statement string = "DELETE FROM images WHERE id = $1;"

	reuslt, err := db.database.Exec(statement, id)
	if err != nil {
		return err
	}

	numRows, err := reuslt.RowsAffected()
	if err != nil {
		return err
	}

	if numRows == 0 {
		return errors.New("Image with the ID " + id + " not found.") 
	}

	return nil
}

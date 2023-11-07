package database

import "errors"

// Deletes a folder form the database
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
		return errors.New("Folder with ID " + " not found")
	}

	return nil
}

func (db *Database) DeleteImage (name string) error {
	return nil
}

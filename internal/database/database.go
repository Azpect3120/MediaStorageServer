package database

import (
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type Database struct {
	connectionString string
	database         *sql.DB
}

// CreateDatabase creates a database object and attempts to connect to the database.
func CreateDatabase() *Database {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error loading environment variables: ", err)
	}

	database := &Database{
		connectionString: os.Getenv("db_url"),
	}

	db, err := sql.Open("postgres", database.connectionString)
	if err != nil {
		log.Fatalln("Error opening connection to database: ", err)
	}

	database.database = db

	return database
}

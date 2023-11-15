package main

import (
	"log"

	"github.com/Azpect3120/MediaStorageServer/internal"
	"github.com/Azpect3120/MediaStorageServer/internal/database"
)

func main() {
	server := internal.CreateServer()

	server.DefineUploadRoot("./uploads")

	database := database.CreateDatabase()

	server.LoadRoutes(database)

	err := server.Run("3000")
	if err != nil {
		log.Fatalln("Error running server: ", err)
	}
}

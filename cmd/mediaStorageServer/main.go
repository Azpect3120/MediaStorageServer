package main

import (
	"log"
	"os"

	"github.com/Azpect3120/MediaStorageServer/internal/server"
	"github.com/Azpect3120/MediaStorageServer/internal/database"
)

func main() {
	server := internal.CreateServer()

	server.DefineUploadRoot("./uploads")

	database := database.CreateDatabase()

	server.SetupCache()

	server.LoadRoutes(database)

	port := os.Getenv("MSS_PORT")
	if port == "" {
		port = "3000"
	}

	err := server.Run(port)
	if err != nil {
		log.Fatalln("Error running server: ", err)
	}
}

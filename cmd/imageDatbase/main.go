package main

import (
	"log"

	"github.com/Azpect3120/MediaStorageServer/internal"
)

func main() {
	server := internal.CreateServer()

	server.LoadRoutes()

	err := server.Run("3000")
	if err != nil {
		log.Fatalln("Error running server: ", err)
	}
}

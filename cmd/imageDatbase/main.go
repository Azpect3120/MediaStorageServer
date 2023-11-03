package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// CORS middle-ware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Add your frontend's port

	r.Use(cors.New(config))

	// Define the upload route
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println(file.Header)

		// Generate a unique filename
		// Store file on local machine
		filename := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)
		uploadPath := filepath.Join("uploads", filename)

		if err := c.SaveUploadedFile(file, uploadPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Here you can add code to save the file path and metadata to your SQL database
		// You might want to store the filename, user information, and other metadata.

		imageUrl := "/images/" + filename

		c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "file": file, "path": uploadPath, "filename": filename, "url": imageUrl})
	})

	r.GET("/images/:imageID", func(c *gin.Context) {
		// imageID := c.Param("imageID")

		// Retrieve the image data from the database based on imageID.
		// You can retrieve the byte array and other metadata.

		// Return the image data as a response.
		// c.Data(http.StatusOK, "image/jpeg", imageBytes) // Adjust the MIME type as needed.
	})

	// Serve static files from the "uploads" directory
	r.Static("/uploads", "./uploads")

	r.Run(":3000")
}

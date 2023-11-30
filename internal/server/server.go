package internal

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/Azpect3120/MediaStorageServer/internal/cache"
	"github.com/Azpect3120/MediaStorageServer/internal/database"
	"github.com/Azpect3120/MediaStorageServer/internal/routes"
)

type Server struct {
	Router      *gin.Engine
	Config      cors.Config
	UploadRoot  string
	FolderCache *cache.Cache
	ImageCache  *cache.Cache
}

// CreateServer creates a new server object with default values.
//
// Sets the mode to "debug mode" while testing is taking place and will be set to "release mode" when the server is active.
//
// Defines the allowed origins to "*".
func CreateServer() *Server {
	gin.SetMode(gin.DebugMode)

	var server *Server = &Server{
		Router:     gin.Default(),
		Config:     cors.DefaultConfig(),
		UploadRoot: "",
	}

	server.Config.AllowOrigins = []string{"*"}
	server.Router.Use(cors.New(server.Config))

	return server
}

// DefineUploadRoot defines the path to '/uploads' directory where the folders will be stored.
//
// Path should point to the '/uploads' directory. ex: '~/Documents/media-server/uploads'.
//
// Path can relative or absolute.
func (s *Server) DefineUploadRoot(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	s.UploadRoot = absPath
	return s.UploadRoot, nil
}

// SetupCache sets up the server's various caches
//
// Folder cache stores 10 most recent GET requests to the '/folders/:id' endpoint
//
// Image cache stores 25 most recent GET request to the '/images/:id' endpoint
func (s *Server) SetupCache() {
	s.FolderCache = cache.NewCache(10)
	s.ImageCache = cache.NewCache(25)
}

// Run runs the server on the provided port.
//
// Returns any errors the server may encounter when attempting to run.
//
// Initializes the uploads folder if it does not exist.
//
// Serves the uploads folder.
func (s *Server) Run(port string) error {
	if s.UploadRoot == "" {
		return errors.New("Servers upload root is not defined. Please define it with server.DefineUploadRoot.")
	}

	_, err := os.Stat(s.UploadRoot)

	if os.IsNotExist(err) {
		err = os.Mkdir(s.UploadRoot, 0755)
	}

	if err != nil {
		return err
	}

	s.Router.Static("/uploads", "./uploads")

	return s.Router.Run(":" + port)
}

// Loads the routes from the "/routes/*" directory
func (s *Server) LoadRoutes(db *database.Database) {
	s.Router.GET("/", routes.Index)
	s.Router.GET("/status", routes.Status)

	s.Router.GET("/v1/folders/:id", func(ctx *gin.Context) { routes.GetFolder(s.FolderCache, db, s.UploadRoot, ctx) })
	s.Router.POST("/v1/folders", func(ctx *gin.Context) { routes.CreateFolder(db, s.UploadRoot, ctx) })
	s.Router.PUT("/v1/folders/:id", func(ctx *gin.Context) { routes.UpdateFolder(s.FolderCache, db, s.UploadRoot, ctx) })
	s.Router.DELETE("/v1/folders/:id", func(ctx *gin.Context) { routes.DeleteFolder(s.FolderCache, s.ImageCache, db, s.UploadRoot, ctx) })
	
	s.Router.GET("/v1/folders/:id/images", func(ctx *gin.Context) { routes.GetFolderImages(s.FolderCache, db, s.UploadRoot, ctx) })

	s.Router.GET("/v1/images/:id", func(ctx *gin.Context) { routes.GetImage(s.FolderCache, s.ImageCache, db, s.UploadRoot, ctx) })
	s.Router.POST("/v1/images/:id", func(ctx *gin.Context) { routes.CreateImage(s.FolderCache, db, s.UploadRoot, ctx) })
	s.Router.DELETE("/v1/images/:id", func(ctx *gin.Context) { routes.DeleteImage(s.FolderCache, s.ImageCache, db, s.UploadRoot, ctx) })

	s.Router.GET("/v1/reports/:id/:email", func(ctx *gin.Context) { routes.SendReport(db, ctx) })
}

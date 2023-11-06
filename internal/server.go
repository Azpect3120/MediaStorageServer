package internal

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/Azpect3120/MediaStorageServer/internal/routes"
)

type Server struct {
	Router *gin.Engine
	Config cors.Config
}

// CreateServer creates a new server object with default values.
//
// Sets the mode to "debug mode" while testing is taking place and will be set to "release mode" when the server is active.
//
// Defines the allowed origins to "*".
func CreateServer() *Server {
	gin.SetMode(gin.DebugMode)

	var server *Server = &Server{
		Router: gin.Default(),
		Config: cors.DefaultConfig(),
	}

	server.Config.AllowOrigins = []string{"*"}

	return server
}

// Run runs the server on the provided port.
//
// Returns any errors the server may encounter when attempting to run.
func (s *Server) Run(port string) error {
	return s.Router.Run(":" + port)
}

// Loads the routes from the "/routes/*" directory 
func (s *Server) LoadRoutes () {
	s.Router.GET("/", routes.Index)
	s.Router.GET("/status", routes.Status)
}

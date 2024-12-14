package api

import (
	"github.com/Sahas001/some-project/db/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *controller.Store
	router *gin.Engine
}

func NewServer(store *controller.Store) *Server {
	server := &Server{
		store:  store,
		router: gin.Default(),
	}

	// Configure and add the CORS middleware
	server.router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Replace with your React app's URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	server.setupRoutes()
	return server

}

func (s *Server) setupRoutes() {
	// user crud operations
	s.router.POST("/user", s.CreateUser)
	s.router.GET("/user/:id", s.GetUser)
	s.router.PUT("/user/:id", s.UpdateUser)
	s.router.DELETE("/user/:id", s.DeleteUser)

	// tasks crud operations
	s.router.POST("/tasks", s.CreateTask)
	s.router.GET("/tasks/user/:id", s.GetTasksByUserID)
	s.router.DELETE("/tasks/:id", s.DeleteTask)
	s.router.PUT("/tasks/:id", s.UpdateTask)
	s.router.DELETE("/tasks/user/:id", s.DeleteTasksByUserID)

	//TODO: add routes to delete tasks with tasks id

	// s.router.GET("/tasks/:id", s.GetTasksByID)
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func okResponse(data interface{}) gin.H {
	return gin.H{"data": data}
}

// func

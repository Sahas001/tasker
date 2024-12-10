package api

import (
	"github.com/Sahas001/some-project/db/controller"
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

	server.setupRoutes()
	return server

}

func (s *Server) setupRoutes() {
	// user crud operations
	s.router.POST("/users", s.CreateUser)
	s.router.GET("/users/:id", s.GetUser)
	s.router.PUT("/users/:id", s.UpdateUser)
	s.router.DELETE("/users/:id", s.DeleteUser)

	// tasks crud operations
	s.router.POST("/tasks", s.CreateTask)
	s.router.GET("/tasks/:id", s.GetTasksByUserID)
	s.router.DELETE("/tasks/:id", s.DeleteTask)
	s.router.PUT("/tasks/:id", s.UpdateTask)
	s.router.DELETE("/tasks/user/:id", s.DeleteTasksByUserID)
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
package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Server is the server struct
type Server struct {
	counter int64
	server  *http.Server
	router  *gin.Engine
}

// New creates a new server
func New() *Server {
	router := gin.Default()
	server := &Server{
		counter: int64(0),
		router:  router,
	}

	router.GET("/", server.CounterHandler)
	router.GET("/health_checks", server.HealthHandler)

	return server
}

// HealthHandler returns a 200 OK response
func (s *Server) HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

// CounterHandler increments the counter and returns its value
func (s *Server) CounterHandler(c *gin.Context) {
	s.counter++
	c.JSON(http.StatusOK, gin.H{
		"counter": s.counter,
	})
}

// Run starts the server
func (s *Server) Start(address string) error {
	s.server = &http.Server{
		Addr:        address,
		Handler:     s.router,
		ReadTimeout: 10 * time.Second,
	}

	log.Printf("Server is running on %s\n", address)

	return s.server.ListenAndServe()
}

// Stop stops the server
func (s *Server) Stop() error {
	log.Println("Stopping server...")

	if s.server == nil {
		return nil
	}

	return s.server.Close()
}

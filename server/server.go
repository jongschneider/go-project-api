package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Server represents the server that supports our api
type Server struct {
	Port       int
	Addr       string
	HTTPServer *http.Server
}

// Start starts the server
func (s *Server) Start() {
	log.Printf("Server starting on port %d", s.Port)
	log.Fatal(s.HTTPServer.ListenAndServe())
}

// New creates a new server
func New(port int) *Server {
	addr := fmt.Sprintf(":%d", port)

	return &Server{
		Port: port,
		Addr: addr,
		HTTPServer: &http.Server{
			Addr:           addr,
			Handler:        nil,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

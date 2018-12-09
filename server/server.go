package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/jongschneider/go-project/router"
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
	r := router.New()

	handler := handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Origin", "Cache-Control", "X-App-Token"}),
		handlers.ExposedHeaders([]string{}),
		handlers.MaxAge(1000),
		handlers.AllowCredentials(),
	)(r.Router))

	handler = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)

	return &Server{
		Port: port,
		Addr: addr,
		HTTPServer: &http.Server{
			Addr:           addr,
			Handler:        handler,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

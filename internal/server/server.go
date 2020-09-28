package server

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

// Server is a base server config
type Server struct {
	server *http.Server
}

// New inicialize a new server with config
func New(port string) (*Server, error) {
	r := chi.NewRouter()

	s := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server := Server{server: s}

	return &server, nil
}

// Close server resources
func (s *Server) Close() error {
	// TODO
	return nil
}

// Start the server
func (s *Server) Start() {
	log.Printf("Server running on http://localhost%s", s.server.Addr)
	log.Fatal(s.server.ListenAndServe())
}

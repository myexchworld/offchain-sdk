package server

import (
	"context"
	"net/http"
)

// Handler is a handler.
type Handler struct {
	Path    string
	Handler http.Handler
}

// Server is a server.
type Server struct {
	mux *http.ServeMux
}

// New creates a new server.
func New() *Server {
	return &Server{
		mux: http.NewServeMux(),
	}
}

// RegisterHandler registers a handler.
func (s *Server) RegisterHandler(h Handler) {
	s.mux.Handle(h.Path, h.Handler)
}

// Start starts the server.
func (s *Server) Start(_ context.Context) {
	if err := http.ListenAndServe(":8080", s.mux); err != nil { //nolint:gosec // todo fix.
		panic(err)
	}
}

// Stop stops the server.
func (s *Server) Stop() {}

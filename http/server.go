package http

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"time"
)

// ShutdownTimeout is the time given for any outstanding requests to finish before server shutdown.
const ShutdownTimeout = 5 * time.Second

// Server represents an HTTP server. It is meant to wrap all HTTP functionality
// used by the application so that dependent packages don't have to import net/http.
type Server struct {
	ln     net.Listener
	server *http.Server
	router *mux.Router
}

// NewServer returns a new instance of Server.
func NewServer() *Server {
	// Create a new server that wraps the net/http server & mux router.
	s := &Server{
		server: &http.Server{},
		router: mux.NewRouter(),
	}

	// TODO: Add middlewares here.

	// Use gorrila/mux router as the server handler.
	s.server.Handler = s.router

	// Set up home page route.
	s.router.HandleFunc("/", s.handleIndex).Methods("GET")

	// Set up error handling route.
	s.router.NotFoundHandler = http.HandlerFunc(s.handleNotFound)

	return s
}

// handleIndex handles requests for the home page.
func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

// handleNotFound handles requests for unknown routes.
func (s *Server) handleNotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Your page cannot be found.", http.StatusNotFound)
}

// Open opens the server for requests.
func (s *Server) Open() (err error) {
	fmt.Println("Server is starting on port 8080")
	// Start listening on port 8080.
	if s.ln, err = net.Listen("tcp", ":8080"); err != nil {
		return err
	}
	// Begin serving requests.
	return s.server.Serve(s.ln)
}

// Close gracefully shuts down the server.
func (s *Server) Close() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()
	return s.server.Shutdown(ctx)

}

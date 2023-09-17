package main

import (
	"gochat/http"
	"os"
)

func main() {
	m := NewMain()

	// Execute the program.
	if err := m.Run(); err != nil {
		m.Close()
		// Report error here
		os.Exit(1)
	}
}

// Main represents the main application.
type Main struct {
	HTTPServer *http.Server
}

// NewMain returns a new instance of Main.
func NewMain() *Main {
	return &Main{
		HTTPServer: http.NewServer(),
	}
}

// Close gracefully shuts down the application.
func (m *Main) Close() (err error) {
	// Close the HTTP server.
	if err = m.HTTPServer.Close(); err != nil {
		return err
	}
	return nil
}

// Run runs the application.
func (m *Main) Run() (err error) {
	// Open the HTTP server.
	if err = m.HTTPServer.Open(); err != nil {
		return err
	}
	return nil
}

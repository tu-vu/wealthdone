package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"wealthfront/http"
)

func main() {
	m := NewMain()

	// Setup signal handlers.
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() { <-c; cancel() }() // Release resources on interrupt signal (Ctrl+C).

	// Execute the program.
	if err := m.Run(ctx); err != nil {
		m.Close()
		os.Exit(1)
	}

	// Wait for Ctrl+C.
	<-ctx.Done()

	// Clean up program.
	if err := m.Close(); err != nil {
		fmt.Fprintln(os.Stderr, err)
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
func (m *Main) Run(ctx context.Context) (err error) {
	// Open the HTTP server.
	if err = m.HTTPServer.Open(); err != nil {
		return err
	}
	return nil
}

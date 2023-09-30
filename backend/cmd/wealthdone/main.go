package main

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
	"os"
	"os/signal"
	fb "wealthfront/firebase"
	"wealthfront/http"
)

const (
	// DefaultCredentialsPath is the default path to the service account key file.
	DefaultCredentialsPath = "./serviceAccountKey.json"
)

func main() {
	// Setup signal handlers.
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() { <-c; cancel() }() // Release resources on interrupt signal (Ctrl+C).

	// Instantiate a new type to represent our application.
	m, err := NewMain(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

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
	FireApp    *firebase.App
	HTTPServer *http.Server
}

// NewMain returns a new instance of Main.
func NewMain(ctx context.Context) (*Main, error) {
	// Create a new firebase app.
	opt := option.WithCredentialsFile(DefaultCredentialsPath)
	fireApp, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, errors.Errorf("failed to create Firebase app: %v", err)
	}

	return &Main{
		HTTPServer: http.NewServer(),
		FireApp:    fireApp,
	}, nil
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
func (m *Main) Run(ctx context.Context) error {
	// Instantiate Firebase-backed services.
	authService, err := fb.NewAuthService(ctx, m.FireApp)
	if err != nil {
		return err
	}

	// Attach underlying services to HTTP server.
	m.HTTPServer.AuthService = authService

	// Open the HTTP server.
	if err := m.HTTPServer.Open(); err != nil {
		return err
	}
	return nil
}

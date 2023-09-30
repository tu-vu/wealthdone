package firebase

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/pkg/errors"
)

// AuthService represents a service for managing auths.
type AuthService struct {
	FireAuth *auth.Client
}

// NewAuthService returns a new instance of AuthService.
func NewAuthService(ctx context.Context, app *firebase.App) (*AuthService, error) {
	authClient, err := app.Auth(ctx)
	if err != nil {
		return nil, errors.Errorf("error getting auth client: %v", err)
	}
	return &AuthService{
		FireAuth: authClient,
	}, nil
}

package wealthdone

import (
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

// AuthService represents a service for managing auths.
type AuthService struct {
	FireAuth *auth.Client
}

// NewAuthService returns a new instance of AuthService.
func NewAuthService(app *firebase.App) *AuthService {
	return &AuthService{
		FireAuth: app.Auth,
	}
}

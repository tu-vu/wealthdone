package wealthdone

import "time"

// User represents a user account.
type User struct {
	ID int `json:"id"`

	// User's basic information.
	Name  string `json:"name"`
	Email string `json:"email"`

	// Password is the user's hashed password.
	Password string `json:"password"`

	// Timestamps for user creation & last update.
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

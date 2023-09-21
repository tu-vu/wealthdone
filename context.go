package wealthdone

import "context"

// contextKey represents an internal key for adding context fields.
type contextKey int

// List of context keys, which are used to store request-scoped data.
const (
	// Stores the current logged-in user in the context.
	userContextKey = contextKey(iota + 1)
)

// NewContextWithUser returns a new context with the given user.
func NewContextWithUser(ctx context.Context, user *User) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}

// UserFromContext returns user stored in context, if any.
func UserFromContext(ctx context.Context) *User {
	u, ok := ctx.Value(userContextKey).(*User)
	if !ok {
		return nil
	}
	return u
}

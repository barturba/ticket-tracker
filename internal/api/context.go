package api

import (
	"context"
	"net/http"

	"github.com/barturba/ticket-tracker/internal/models"
)

// Define a custom context type, based on the underlying string type.
type contextKey string

// Convert the string "user" to a contextKey type and assign it to the
// userContextKey constant.
const userContextKey = contextKey("user")

func ContextSetUser(r *http.Request, user *models.User) *http.Request {
	ctx := context.WithValue(r.Context(), userContextKey, user)
	return r.WithContext(ctx)
}

func ContextGetUser(r *http.Request) *models.User {
	user, ok := r.Context().Value(userContextKey).(*models.User)
	if !ok {
		panic("missing user value in request context")
	}
	return user
}

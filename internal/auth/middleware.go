package auth

import (
	"context"
	"election-api/internal/entity"
	"github.com/dgrijalva/jwt-go"
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/go-ozzo/ozzo-routing/v2/auth"
)

type contextKey int

const (
	userKey contextKey = iota
)

// Handler returns a JWT-based authentication middleware.
func Handler(verificationKey string) routing.Handler {
	return auth.JWT(verificationKey, auth.JWTOptions{TokenHandler: handleToken})
}

// handleToken stores the user identity in the request context so that it can be accessed elsewhere.
func handleToken(c *routing.Context, token *jwt.Token) error {
	ctx := WithUser(
		c.Request.Context(),
		token.Claims.(jwt.MapClaims)["id"].(string),
		token.Claims.(jwt.MapClaims)["name"].(string),
	)
	c.Request = c.Request.WithContext(ctx)
	return nil
}

// WithUser returns a context that contains the user identity from the given JWT.
func WithUser(ctx context.Context, id, name string) context.Context {
	return context.WithValue(ctx, userKey, entity.User{})
}

// CurrentUser returns the user identity from the given context.
// Nil is returned if no user identity is found in the context.
func CurrentUser(ctx context.Context) Identity {
	if user, ok := ctx.Value(userKey).(entity.User); ok {
		return user
	}
	return nil
}

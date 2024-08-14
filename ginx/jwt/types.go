package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// RegisteredClaims is the registered claims.
type RegisteredClaims[T any] struct {
	Data T `json:"data,omitempty"`
	jwt.RegisteredClaims
}

type Manager[T any] interface {
	// MiddlewareBuilder returns a new middleware builder.
	MiddlewareBuilder() *MiddlewareBuilder[T]

	// Refresh returns a new refresh handler.
	Refresh(ctx *gin.Context)

	// GenerateRefreshToken generates a new refresh token.
	// must set refresh token to the context, else return errEmptyRefreshOpts error.
	GenerateRefreshToken(data T) (string, error)

	// GenerateAccessToken generates a new access token.
	// must set access token to the context, else return errEmptyAccessOpts error.
	GenerateAccessToken(data T) (string, error)

	// VerifyRefreshToken verifies a refresh token.
	VerifyRefreshToken(token string, opts ...jwt.ParserOption) (RegisteredClaims[T], error)

	// VerifyAccessToken verifies an access token.
	VerifyAccessToken(token string, opts ...jwt.ParserOption) (RegisteredClaims[T], error)

	// SetClaims sets the claims to the context.
	SetClaims(ctx *gin.Context, claims RegisteredClaims[T])
}

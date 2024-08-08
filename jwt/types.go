package jwt

import "github.com/golang-jwt/jwt/v5"

type RegisteredClaims[T any] struct {
	Data T `json:"data,omitempty"`
	jwt.RegisteredClaims
}

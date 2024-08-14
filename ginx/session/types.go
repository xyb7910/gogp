package session

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	Uid  int64
	SSId string
	Data map[string]string
}

func (c *Claims) Get(key string) (any, error) {
	val, ok := c.Data[key]
	if !ok {
		return nil, errors.New("key not exists")
	}
	return val, nil
}

type Session interface {
	// Set the data to the session
	Set(ctx *context.Context, key string, val any) error

	// Get the data from the session
	Get(ctx *context.Context, key string) (any, error)

	// Del the data from the session
	Del(ctx *context.Context, key string) error

	// Destroy the session
	Destroy(ctx *context.Context) error

	// Claims encode data to jwt token
	Claims() Claims
}

// Provider is the interface of the session provider
type Provider interface {
	// NewSession create a new session
	// ctx is the gin.Context
	// uid is the user id
	// jwtData will be stored in the jwt token
	// sessData will be stored in the session
	NewSession(ctx *gin.Context, uid int64, jwtData map[string]string, sessData map[string]any) (Session, error)

	// GetSession get the session by the context
	// if session exists, return the session
	// else return nil, error
	GetSession(ctx *gin.Context) (Session, error)

	// UpdateClaims update the claims of the session
	// the claims is stored in the jwt token,and the jwt token is unchangeable,
	// so, we need to update the claims in the jwt token by generating a new jwt token
	// and must have the same session id
	UpdateClaims(ctx *gin.Context, claims Claims) error

	// RenewAccessToken renew the access token of the session
	RenewAccessToken(ctx *gin.Context) error
}

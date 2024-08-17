package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/xyb7910/gogp/generics/set"
	"log/slog"
	"net/http"
	"time"
)

// MiddlewareBuilder is a builder for JWT middleware, use it to check login.
type MiddlewareBuilder[T any] struct {
	// ignorePath ignores the path, if the path is ignored, the middleware will not check the login.
	ignoredPath func(path string) bool
	manager     *Management[T]
	nowFunc     func() time.Time
}

// NewMiddlewareBuilder returns a new MiddlewareBuilder.
func NewMiddlewareBuilder[T any](m *Management[T]) *MiddlewareBuilder[T] {
	return &MiddlewareBuilder[T]{
		manager: m,
		ignoredPath: func(path string) bool {
			return false
		},
		nowFunc: m.nowFunc,
	}
}

// IgnorePath sets the paths that will be ignored.
func (m *MiddlewareBuilder[T]) IgnorePath(path ...string) *MiddlewareBuilder[T] {
	return m.IgnoredPathFunc(staticIgnorePaths(path))
}

// staticIgnorePaths returns a function that checks if the path is ignored.
func staticIgnorePaths(paths []string) func(path string) bool {
	s := set.NewMapSet[string](len(paths))
	for _, path := range paths {
		s.Add(path)
	}
	return func(path string) bool {
		return s.Exist(path)
	}
}

// IgnoredPathFunc sets the function that checks if the path is ignored.
func (m *MiddlewareBuilder[T]) IgnoredPathFunc(fn func(path string) bool) *MiddlewareBuilder[T] {
	m.ignoredPath = fn
	return m
}

// Build returns the middleware.
func (m *MiddlewareBuilder[T]) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if m.ignoredPath(ctx.Request.URL.Path) {
			return
		}

		// extract token from header
		tokenStr := m.manager.extractTokenString(ctx)
		if tokenStr == "" {
			slog.Debug("token is empty")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// parse token and check token
		claim, err := m.manager.VerifyAccessToken(tokenStr, jwt.WithTimeFunc(m.nowFunc))
		if err != nil {
			slog.Debug("token is invalid")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// set claim
		m.manager.SetClaims(ctx, claim)
	}
}

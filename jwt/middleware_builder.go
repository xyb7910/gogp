package jwt

// MiddlewareBuilder is a builder for JWT middleware, use it to check login.
type MiddlewareBuilder struct {
	// ignorePath ignores the path, if the path is ignored, the middleware will not check the login.
	ignoredPath func(path string) bool
	manager
}

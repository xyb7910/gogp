package jwt

type Management[T any] struct {
	allowTokenHeader    string
	exposeAccessHeader  string
	exposeRefreshHeader string

	accessJWTOptions Options
}

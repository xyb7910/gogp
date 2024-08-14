package ratelimit

import "context"

//go:generate mockgen -source=types.go -destination=./mocks/mock_types.go -package=limitmocks Limiter
type Limiter interface {
	// Limit returns true if the key is limited, otherwise false.
	Limit(ctx context.Context, key string) (bool, error)
}

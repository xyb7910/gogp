package ratelimit

import (
	"context"
	_ "embed"
	"github.com/redis/go-redis/v9"
	"time"
)

//go:embed slide_window.lua
var luaSlideWindow string

type RedisSlidingWindowRateLimiter struct {
	Cmd redis.Cmdable

	// windowSize is the size of the sliding window.
	windowSize time.Duration

	// Ratio is the ratio of the number of requests allowed in the window.
	Ratios int
}

func (r *RedisSlidingWindowRateLimiter) Limit(ctx context.Context, key string) (bool, error) {
	return r.Cmd.Eval(ctx, luaSlideWindow, []string{key},
		r.windowSize.Milliseconds(), r.Ratios, time.Now().UnixMilli()).Bool()
}

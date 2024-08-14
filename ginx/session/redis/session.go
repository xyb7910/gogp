package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/xyb7910/gogp/session"
	"time"
)

var _ session.Session = (*Session)(nil)

type Session struct {
	client     redis.Cmdable
	key        string
	claims     session.Claims
	expiration time.Duration
}

func NewRedisSession(
	ssid string,
	expiration time.Duration,
	client redis.Cmdable,
	cl session.Claims) *Session {
	return &Session{
		client:     client,
		key:        "redis_session:" + ssid,
		expiration: expiration,
		claims:     cl,
	}
}

func (s Session) Set(ctx context.Context, key string, value any) error {
	return s.client.HMSet(ctx, s.key, key, value).Err()
}

func (s Session) Get(ctx context.Context, key string) (any, error) {
	res, err := s.client.HGet(ctx, s.key, key).Result()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s Session) Del(ctx context.Context, key string) error {
	return s.client.HDel(ctx, s.key, key).Err()
}

func (s Session) Destroy(ctx context.Context) error {
	return s.client.Del(ctx, s.key).Err()
}

func (s Session) Claims() session.Claims {
	return s.claims
}

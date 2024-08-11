package session

import (
	"context"
	"errors"
	"github.com/xyb7910/gogp/gctx"
)

type Claims struct {
	Uid  int64
	SSId string
	Data map[string]string
}

type Session interface {
	//Set 将数据写进 session
	Set(ctx context.Context, key string, value any) error
	//Get 从 session 中获取数据
	Get(ctx context.Context, key string) (any, error)
	// Del 从 session 中删除数据
	Del(ctx context.Context, key string) error
	// Destroy 销毁 session
	Destroy(ctx context.Context) error
	// Claims 获取 session 的 claims
	Claims() Claims
}

func (c Claims) Get(key string) (any, error) {
	val, ok := c.Data[key]
	if !ok {
		return nil, errors.New("key not found")
	}
	return val, nil
}

// Provider 定义了 Session 的提供者，所有的 session 都支持 jwt
type Provider interface {
	// NewSession 创建一个新的 session， jwtData 是 jwt 的数据， sessData 是 session 的数据
	NewSession(ctx *gctx.Context, uid int64,
		jwtData map[string]string, sessData map[string]string) (Session, error)

	// Get 获取 session
	Get(ctx *gctx.Context) (Session, error)

	// UpdateClaims 更新 session 的 claims,必须传进去 SSid 否则无法更新
	UpdateClaims(ctx *gctx.Context, claims Claims) error

	// RenewAccessToken 刷新 access token
	RenewAccessToken(ctx *gctx.Context) error
}

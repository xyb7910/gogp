package session

import (
	"github.com/xyb7910/gogp/gctx"
)

var defaultProvider Provider

func NewSession(ctx *gctx.Context, uid int64,
	jwtData map[string]string, sessData map[string]string) (Session, error) {
	return defaultProvider.NewSession(ctx, uid, jwtData, sessData)
}

func Get(ctx *gctx.Context) (Session, error) {
	return defaultProvider.Get(ctx)
}

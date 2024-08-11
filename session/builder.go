package session

import "github.com/xyb7910/gogp/gctx"

type Builder struct {
	ctx      *gctx.Context
	uid      int64
	jwtData  map[string]string
	sessData map[string]string
	sp       Provider
}

func NewSessionBuilder(ctx *gctx.Context, uid int64) *Builder {
	return &Builder{
		ctx: ctx,
		uid: uid,
		sp:  defaultProvider,
	}
}

func (b *Builder) SetProvider(p Provider) *Builder {
	b.sp = p
	return b
}

func (b *Builder) SetJWTData(data map[string]string) *Builder {
	b.jwtData = data
	return b
}

func (b *Builder) SetSessionData(data map[string]string) *Builder {
	b.sessData = data
	return b
}

func (b *Builder) Build() (Session, error) {
	return b.sp.NewSession(b.ctx, b.uid, b.jwtData, b.sessData)
}

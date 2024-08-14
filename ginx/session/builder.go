package session

import "github.com/gin-gonic/gin"

// Builder assists in building a session
type Builder struct {
	Ctx      *gin.Context
	Uid      int64
	jwtData  map[string]string
	sessData map[string]any
	sp       Provider
}

func NewSessionBuilder(ctx *gin.Context, uid int64, sp Provider) *Builder {
	return &Builder{
		Ctx: ctx,
		Uid: uid,
		sp:  sp,
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

func (b *Builder) SetSessionData(data map[string]any) *Builder {
	b.sessData = data
	return b
}

func (b *Builder) Builder() (Session, error) {
	return b.sp.NewSession(b.Ctx, b.Uid, b.jwtData, b.sessData)
}

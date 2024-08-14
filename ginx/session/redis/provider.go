package redis

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	ijwt "github.com/xyb7910/gogp/ginx/jwt"
	"github.com/xyb7910/gogp/ginx/session"
	"time"
)

var _ session.Provider = (*SessionProvider)(nil)

type SessionProvider struct {
	client      redis.Cmdable
	m           ijwt.Manager[session.Claims]
	tokenHeader string
	atHeader    string
	rtHeader    string
	expiration  time.Duration
}

func (s *SessionProvider) NewSession(ctx *gin.Context, uid int64, jwtData map[string]string, sessData map[string]any) (session.Session, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SessionProvider) GetSession(ctx *gin.Context) (session.Session, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SessionProvider) UpdateClaims(ctx *gin.Context, claims session.Claims) error {
	//TODO implement me
	panic("implement me")
}

func (s *SessionProvider) RenewAccessToken(ctx *gin.Context) error {
	//TODO implement me
	panic("implement me")
}

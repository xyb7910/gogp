package session

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

var CTX_SESSION = "_session"

type MiddlewareBuilder struct {
	sp Provider
}

func (b *MiddlewareBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sess, err := b.sp.GetSession(ctx)
		if err != nil {
			slog.Error("get session error", err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set(CTX_SESSION, sess)
	}
}

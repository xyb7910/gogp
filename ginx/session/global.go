package session

import "github.com/gin-gonic/gin"

var defaultProvider Provider

func RegisterProvider(provider Provider) {
	defaultProvider = provider
}

func NewSession(ctx *gin.Context, uid int64, jwtData map[string]string,
	sessData map[string]any) (Session, error) {
	return defaultProvider.NewSession(ctx, uid, jwtData, sessData)
}

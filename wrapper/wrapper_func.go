package wrapper

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/xyb7910/gogp"
	"github.com/xyb7910/gogp/gctx"
	"github.com/xyb7910/gogp/ginx/session"
	"log/slog"
	"net/http"
)

type Result struct {
	Code int
	Msg  string
	Data any
}

type Context = gctx.Context

func W(fn func(ctx *gin.Context) (Result, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res, err := fn(ctx)
		if errors.Is(err, gogp.ErrNoResponse) {
			slog.Debug("不需要响应", slog.Any("err", err))
			return
		}
		if errors.Is(err, gogp.ErrUnauthorized) {
			slog.Debug("未授权", slog.Any("err", err))
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if err != nil {
			slog.Error("执行业务逻辑失败", slog.Any("err", err))
			ctx.PureJSON(http.StatusInternalServerError, res)
			return
		}
		ctx.PureJSON(http.StatusOK, res)
	}
}

func B[Req any](fn func(ctx *gin.Context, req Req) (Result, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req Req
		if err := ctx.Bind(&req); err != nil {
			slog.Debug("绑定参数失败", slog.Any("err", err))
		}
		res, err := fn(ctx, req)
		if errors.Is(err, gogp.ErrNoResponse) {
			slog.Debug("不需要返回值", slog.Any("err", err))
		}
		if errors.Is(err, gogp.ErrUnauthorized) {
			slog.Debug("未授权", slog.Any("err", err))
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if err != nil {
			slog.Error("执行业务逻辑失败", slog.Any("err", err))
			ctx.PureJSON(http.StatusInternalServerError, res)
			return
		}
		ctx.PureJSON(http.StatusOK, res)
	}
}

// BS 带session的业务逻辑包装器
func BS[Req any](fn func(ctx *gin.Context, req Req, sess session.Session) (Result, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sess, err := session.Get(ctx)
		if err != nil {
			slog.Debug("获取session失败", slog.Any("err", err))
			return
		}
		var req Req
		if err := ctx.Bind(&req); err != nil {
			slog.Debug("绑定参数失败", slog.Any("err", err))
			return
		}
		res, err := fn(ctx, req, sess)
		if errors.Is(err, gogp.ErrNoResponse) {
			slog.Debug("不需要响应", slog.Any("err", err))
			return
		}
		// 如果里面有权限校验，那么会返回 401 错误（目前来看，主要是登录态校验）
		if errors.Is(err, gogp.ErrUnauthorized) {
			slog.Debug("未授权", slog.Any("err", err))
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if err != nil {
			slog.Error("执行业务逻辑失败", slog.Any("err", err))
			ctx.PureJSON(http.StatusInternalServerError, res)
			return
		}
		ctx.PureJSON(http.StatusOK, res)
	}
}

func S(fn func(ctx *gin.Context, sess session.Session) (Result, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sess, err := session.Get(ctx)
		if err != nil {
			slog.Debug("获取 Session 失败", slog.Any("err", err))
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		res, err := fn(ctx, sess)
		if errors.Is(err, gogp.ErrNoResponse) {
			slog.Debug("不需要响应", slog.Any("err", err))
			return
		}
		// 如果里面有权限校验，那么会返回 401 错误（目前来看，主要是登录态校验）
		if errors.Is(err, gogp.ErrUnauthorized) {
			slog.Debug("未授权", slog.Any("err", err))
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if err != nil {
			slog.Error("执行业务逻辑失败", slog.Any("err", err))
			ctx.PureJSON(http.StatusInternalServerError, res)
			return
		}
		ctx.PureJSON(http.StatusOK, res)
	}
}

package jwt

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/xyb7910/gogp/ginx/option"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

const bearerPrefix = "Bearer"

var (
	errEmptyRefreshOpts = errors.New("refreshJWTOptions are nil")
)

type Management[T any] struct {
	allowTokenHeader    string // require header with token
	exposeAccessHeader  string // expose access token
	exposeRefreshHeader string // expose refresh token

	accessJWTOptions   Options          // access token options
	refreshJWTOptions  *Options         // refresh token options
	rotateRefreshToken bool             // converting refresh token to access token
	nowFunc            func() time.Time // control jwt token time
}

func (m *Management[T]) MiddlewareBuilder() *MiddlewareBuilder[T] {
	return NewMiddlewareBuilder[T](m)
}

// NewManagement create a new management
func NewManagement[T any](accessJWTOptions Options,
	opts ...option.Option[Management[T]]) *Management[T] {
	DOpts := defaultManagementOptions[T]()
	DOpts.accessJWTOptions = accessJWTOptions
	option.WithOption[Management[T]](&DOpts, opts...)

	return &DOpts
}

func defaultManagementOptions[T any]() Management[T] {
	return Management[T]{
		allowTokenHeader:    "authorization",
		exposeAccessHeader:  "xxx-access-token",
		exposeRefreshHeader: "xxx-refresh-token",
		rotateRefreshToken:  false,
		nowFunc:             time.Now,
	}
}

func WithAllowTokenHeader[T any](header string) option.Option[Management[T]] {
	return func(m *Management[T]) {
		m.allowTokenHeader = header
	}
}

func WithExposeAccessHeader[T any](header string) option.Option[Management[T]] {
	return func(m *Management[T]) {
		m.exposeAccessHeader = header
	}
}

func WithExposeRefreshHeader[T any](header string) option.Option[Management[T]] {
	return func(m *Management[T]) {
		m.exposeRefreshHeader = header
	}
}

func WithRefreshJWTOptions[T any](refreshOpts Options) option.Option[Management[T]] {
	return func(m *Management[T]) {
		m.refreshJWTOptions = &refreshOpts
	}
}

func WithRotateRefreshToken[T any](isRotate bool) option.Option[Management[T]] {
	return func(m *Management[T]) {
		m.rotateRefreshToken = isRotate
	}
}

func WithNowFunc[T any](nowFunc func() time.Time) option.Option[Management[T]] {
	return func(m *Management[T]) {
		m.nowFunc = nowFunc
	}
}

// extractTokenString extract token string from header
func (m *Management[T]) extractTokenString(ctx *gin.Context) string {
	authCode := ctx.GetHeader(m.allowTokenHeader)
	if authCode == "" {
		return ""
	}
	var b strings.Builder
	b.WriteString(bearerPrefix)
	b.WriteString("")
	prefix := b.String()
	if strings.HasPrefix(authCode, prefix) {
		return authCode[len(prefix):]
	}
	return ""
}

// VerifyAccessToken verify access token with options
func (m *Management[T]) VerifyAccessToken(token string, opts ...jwt.ParserOption) (RegisteredClaims[T], error) {
	t, err := jwt.ParseWithClaims(token, &RegisteredClaims[T]{},
		func(*jwt.Token) (interface{}, error) {
			return []byte(m.accessJWTOptions.DecryptionKey), nil
		},
		opts...,
	)
	if err != nil || !t.Valid {
		return RegisteredClaims[T]{}, err
	}
	clm, _ := t.Claims.(*RegisteredClaims[T])
	return *clm, nil
}

// VerifyRefreshToken verify refresh token with options
func (m *Management[T]) VerifyRefreshToken(token string, opts ...jwt.ParserOption) (RegisteredClaims[T], error) {
	if m.refreshJWTOptions == nil {
		return RegisteredClaims[T]{}, errEmptyRefreshOpts
	}
	t, err := jwt.ParseWithClaims(token, &RegisteredClaims[T]{},
		func(*jwt.Token) (interface{}, error) {
			return []byte(m.refreshJWTOptions.DecryptionKey), nil
		},
		opts...,
	)
	if err != nil || !t.Valid {
		return RegisteredClaims[T]{}, err
	}
	clm, _ := t.Claims.(*RegisteredClaims[T])
	return *clm, nil
}

// GenerateAccessToken generate access token with data
func (m *Management[T]) GenerateAccessToken(data T) (string, error) {
	nowTime := m.nowFunc()
	claims := RegisteredClaims[T]{
		Data: data,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    m.accessJWTOptions.Issuer,
			ExpiresAt: jwt.NewNumericDate(nowTime.Add(m.accessJWTOptions.Expire)),
			IssuedAt:  jwt.NewNumericDate(nowTime),
			ID:        m.accessJWTOptions.genIDFn(),
		},
	}
	token := jwt.NewWithClaims(m.accessJWTOptions.Method, claims)
	return token.SignedString([]byte(m.accessJWTOptions.EncryptionKey))
}

// GenerateRefreshToken generate refresh token with data
func (m *Management[T]) GenerateRefreshToken(data T) (string, error) {
	if m.refreshJWTOptions == nil {
		return "", errEmptyRefreshOpts
	}

	nowTime := m.nowFunc()
	claims := RegisteredClaims[T]{
		Data: data,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    m.refreshJWTOptions.Issuer,
			ExpiresAt: jwt.NewNumericDate(nowTime.Add(m.refreshJWTOptions.Expire)),
			IssuedAt:  jwt.NewNumericDate(nowTime),
			ID:        m.refreshJWTOptions.genIDFn(),
		},
	}

	token := jwt.NewWithClaims(m.refreshJWTOptions.Method, claims)
	return token.SignedString([]byte(m.refreshJWTOptions.EncryptionKey))
}

// Refresh refresh token handler
func (m *Management[T]) Refresh(ctx *gin.Context) {
	if m.refreshJWTOptions == nil {
		slog.Error("refresh token options is nil")
		ctx.Status(http.StatusInternalServerError)
		return
	}
	tokenStr := m.extractTokenString(ctx)
	claims, err := m.VerifyAccessToken(tokenStr,
		jwt.WithTimeFunc(m.nowFunc))
	if err != nil {
		slog.Error("verify access token error: ", err)
		ctx.Status(http.StatusUnauthorized)
		return
	}
	accessToken, err := m.GenerateAccessToken(claims.Data)
	if err != nil {
		slog.Error("fail to generate access token: ", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.Header(m.exposeAccessHeader, accessToken)

	if m.rotateRefreshToken {
		refreshedToken, err := m.GenerateRefreshToken(claims.Data)
		if err != nil {
			slog.Error("fail to generate refresh token: ", err)
			ctx.Status(http.StatusInternalServerError)
			return
		}
		ctx.Header(m.exposeRefreshHeader, refreshedToken)
	}
	ctx.Status(http.StatusNoContent)
}

// SetClaims set claims to context
func (m *Management[T]) SetClaims(ctx *gin.Context, claims RegisteredClaims[T]) {
	ctx.Set("Claims", claims)
}

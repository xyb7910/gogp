package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/xyb7910/gogp/ginx/option"
	"time"
)

type Options struct {
	Expire        time.Duration     // expire time
	EncryptionKey string            // encryption key
	DecryptionKey string            // decryption key
	Method        jwt.SigningMethod // encryption method
	Issuer        string            // issuer
	genIDFn       func() string     // generate id function
}

func NewOptions(expire time.Duration, encryptionKey string,
	opts ...option.Option[Options]) Options {
	Opts := Options{
		Expire:        expire,
		EncryptionKey: encryptionKey,
		DecryptionKey: encryptionKey,
		Method:        jwt.SigningMethodES256,
		genIDFn:       func() string { return "" },
	}
	option.WithOption[Options](&Opts, opts...)
	return Opts
}

func WithDecryptionKey(key string) option.Option[Options] {
	return func(o *Options) {
		o.DecryptionKey = key
	}
}

func WithMethod(method jwt.SigningMethod) option.Option[Options] {
	return func(o *Options) {
		o.Method = method
	}
}

func WithIssuer(issuer string) option.Option[Options] {
	return func(o *Options) {
		o.Issuer = issuer
	}
}

func WithGenIDFn(fn func() string) option.Option[Options] {
	return func(o *Options) {
		o.genIDFn = fn
	}
}

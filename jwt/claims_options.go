package jwt

import (
	"github.com/golang-jwt/jwt/v5"
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

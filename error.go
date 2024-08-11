package gogp

import (
	"errors"
)

var ErrNoResponse = errors.New("不需要返回 response")
var ErrUnauthorized = errors.New("未授权")
var ErrSessionKeyNotFound = errors.New("session 中没有找到这个 key")

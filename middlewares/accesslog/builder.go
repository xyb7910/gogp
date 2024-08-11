package accesslog

import "time"

type AccessLog struct {
	// http 请求类型
	Method string
	// url 整个请求的url
	Url string
	// 请求体
	ReqBody string
	// 响应体
	RespBody string
	// 处理时间
	Duration time.Duration
	// 状态码
	StatusCode int
}

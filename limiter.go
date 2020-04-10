package ratelimit

// 限流适配接口
type Limiter interface {
	Allow() bool
}

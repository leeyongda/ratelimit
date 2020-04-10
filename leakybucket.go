package ratelimit

import "go.uber.org/ratelimit"

// 漏桶
type LeakyBucket struct {
	ratelimit.Limiter
}

// 漏桶算法限流
func NewLeakyBucket(rate int) *LeakyBucket {
	rl := ratelimit.New(rate)
	return &LeakyBucket{Limiter: rl}
}

func (l *LeakyBucket) Allow() bool {
	l.Limiter.Take()
	return true
}

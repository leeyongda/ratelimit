package ratelimit

import "github.com/juju/ratelimit"

type TokenBucket struct {
	*ratelimit.Bucket
}

// 令牌桶算法限流
func NewTokenBucket(rate float64, capacity int64) *TokenBucket {
	b := ratelimit.NewBucketWithRate(rate, capacity)
	return &TokenBucket{b}
}

func (t *TokenBucket) Allow() bool {
	return t.TakeAvailable(1) > 0
}

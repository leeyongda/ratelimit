package ratelimit

type Adapter struct {
	Limiter
}

func NewAdapter(l Limiter) Limiter {
	if l == nil {
		return nil
	}
	return &Adapter{Limiter: l}
}

// 单例实现接口
func (a *Adapter) Allow() bool {
	if a != nil {
		return a.Limiter.Allow()
	}
	return false
}

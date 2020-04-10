### 限流适配器
#### 支持令牌桶，漏桶算法限流
#### 使用第三方开源限流库
漏桶算法
> go.uber.org/ratelimit  

令牌桶算法
> github.com/juju/ratelimit

### demo
```go
func main() {
	// 漏桶限流
	// b := ratelimit.NewLeakyBucket(5)
	// 令牌桶限流
	b := ratelimit.NewTokenBucket(5, 10)
	ad := ratelimit.NewAdapter(b)
	if ad.Allow() {
		fmt.Println("allow", time.Now())
	} else {
		fmt.Println("too many requests")
	}
}
```
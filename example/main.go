package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leeyongda/ratelimit"
)

func Limiter(l ratelimit.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		if l == nil {
			c.Next()
			return
		}
		if !l.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"msg": "too many requests",
			})
			return
		}
		c.Next()
	}
}

func NewRouter(l ratelimit.Limiter) *gin.Engine {
	g := gin.New()
	g.Use(Limiter(l))
	return g
}

func main() {
	// b := ratelimit.NewLeakyBucket(5)
	b := ratelimit.NewTokenBucket(5, 10)
	ad := ratelimit.NewAdapter(b)
	g := NewRouter(b)
	g.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "pong"})
	})
	log.Fatal(g.Run(":8080"))
}

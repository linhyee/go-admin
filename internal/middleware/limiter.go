package middleware

import (
	"blog/pkg/app"
	"blog/pkg/errcode"
	"blog/pkg/limiter"

	"github.com/gin-gonic/gin"
)

func RateLimter(l limiter.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				app.NewResponse(c).ToErrResponse(errcode.TooManyRequest)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

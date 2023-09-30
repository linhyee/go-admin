package limiter

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

type MethodLimit struct {
	*Limit
}

func NewMethodLimit() Limiter {
	l := &Limit{limitBuckets: make(map[string]*ratelimit.Bucket)}
	return MethodLimit{Limit: l}
}

func (l MethodLimit) Key(c *gin.Context) string {
	uri := c.Request.RequestURI
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	}
	return uri[:index]
}

func (l MethodLimit) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := l.limitBuckets[key]
	return bucket, ok
}

func (l MethodLimit) AddBuckets(rules ...LimitBucketRule) Limiter {
	for _, rule := range rules {
		if _, ok := l.limitBuckets[rule.Key]; !ok {
			bucket := ratelimit.NewBucketWithQuantum(
				rule.FillInterval,
				rule.Capacity,
				rule.Quantum,
			)
			l.limitBuckets[rule.Key] = bucket
		}
	}
	return l
}

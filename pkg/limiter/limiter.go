package limiter

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

type LimiterBucketRule struct {
	Key          string
	FillInterval time.Duration
	Capacity     int64
	Quantm       int64
}

type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

type LimiterIface interface {
	Key(c *gin.Context) string
	GetBucket(key string) (*ratelimit.Bucket, bool)
	AddBuckets(rules ...LimiterBucketRule) LimiterIface
}

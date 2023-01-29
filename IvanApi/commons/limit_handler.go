package commons

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"time"
)

var buckets = make(map[string]*ratelimit.Bucket)

func LimitHandler(duration time.Duration, capacity int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		route := c.Request.RequestURI
		b, err := buckets[route]
		if !err {
			b = ratelimit.NewBucket(duration, capacity)
			buckets[route] = b
		}
		before := b.Available()
		tokenGet := b.TakeAvailable(1)
		fmt.Println("route:", route)
		if tokenGet != 0 {
			fmt.Println("获取令牌成功, 之前数量：", before, "之后数量：", b.Available())
		} else {
			fmt.Println("获取令牌失败")
			c.JSON(http.StatusBadRequest, gin.H{
				"status": -1,
				"msg":    "点击过于频繁，请稍后重试",
				"data":   nil,
			})
			c.Abort()
			return
		}
	}
}

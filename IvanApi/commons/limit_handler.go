package commons

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"time"
)

func LimitHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		before := b.Available()
		tokenGet := b.TakeAvailable(1)
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

var b *ratelimit.Bucket

func LimitInit(duration time.Duration, capacity int64) {
	b = ratelimit.NewBucket(duration, capacity)
}

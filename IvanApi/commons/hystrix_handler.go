package commons

import (
	"errors"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HystrixWrapper() gin.HandlerFunc {
	return func(c *gin.Context) {
		hystrix.Do("hystrixwrapper", func() error {
			c.Next()
			code := c.Writer.Status()
			if code != http.StatusOK {
				return errors.New(fmt.Sprintf("status code %d", code))
			}
			return nil
		}, func(err error) error {
			if err == nil {
				c.JSON(http.StatusServiceUnavailable, gin.H{
					"msg": err.Error(),
				})
			}
			return nil
		})
	}
}

func HystrixBreakWrapper(c *gin.Context) {
	hystrix.Do("hystrixwrapper", func() error {
		c.Next()
		code := c.Writer.Status()
		if code != http.StatusOK {
			return errors.New(fmt.Sprintf("status code %d", code))
		}
		return nil
	}, func(err error) error {
		if err != nil {

		}
		return nil
	})
}

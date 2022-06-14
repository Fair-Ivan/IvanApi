package router

import (
	"IvanApi/apis"
	"IvanApi/commons"
	docs "IvanApi/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(Recover)
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/ops")
		{
			eg.PUT("/broadcast", apis.UpdateBroadCast)
			eg.GET("/getbroadcast", apis.GetBroadCastList)
			eg.POST("/broadcast", apis.AddBroadCast)
			eg.DELETE("/broadcast", apis.RemoveBroadCast)
		}
		ag := v1.Group("/test")
		ag.Use(commons.LimitHandler())
		{
			ag.GET("/", apis.TestApi)
			ag.GET("/second", apis.TestApi2)
			ag.GET("/third", apis.TestApi3)
		}
		sg := v1.Group("/login")
		{
			sg.GET("", apis.Login)
		}
		fg := v1.Group("/oss")
		{
			fg.POST("callback", apis.OssCheckCallback)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Print("panic: %v\n", r)
			c.JSON(http.StatusOK, gin.H{"msg": ErrorToString(r)})
		}
	}()
	c.Next()
}

func ErrorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}

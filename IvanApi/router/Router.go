package router

import (
	"IvanApi/apis"
	docs "IvanApi/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.New()
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
		{
			//ag.Use(commons.JWTAuth()).GET("", apis.TestApi)
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

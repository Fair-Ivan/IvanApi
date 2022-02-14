package Router

import (
	"IvanApi/Apis"
	"IvanApi/Commons"
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
			eg.PUT("/broadcast", Apis.UpdateBroadCast)
			eg.GET("/getbroadcast", Apis.GetBroadCastList)
			eg.POST("/broadcast", Apis.AddBroadCast)
			eg.DELETE("/broadcast", Apis.RemoveBroadCast)
		}
		ag := v1.Group("/test")
		{
			ag.Use(Commons.JWTAuth()).GET("", Apis.TestApi)
		}
		sg := v1.Group("/login")
		{
			sg.GET("", Apis.Login)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}

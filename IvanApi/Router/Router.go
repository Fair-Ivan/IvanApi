package Router

import (
	"IvanApi/Apis"
	docs "IvanApi/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router()  *gin.Engine{
	r := gin.New()
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/ops")
		{
			eg.GET("/getbroadcast",Apis.GetBroadCastList)
		}
	}

	return r
}
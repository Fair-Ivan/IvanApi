package Apis

import (
	"IvanApi/Commons"
	"github.com/gin-gonic/gin"
)

// @Summary 测试
// @Schemes
// @Description test
// @Tags test
// @Accept json
// @Produce json
// @Success 200 {string} TestApi
// @Router /test [get]
func TestApi(g *gin.Context) {
	Commons.PublishMessage("发送消息")
}

package apis

import (
	"IvanApi/commons"
	"IvanApi/model"
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary oss callback
// @Schemes
// @Description oss callback
// @Tags oss
// @Param input body model.OssForm true "input"
// @Accept json
// @Produce json
// @Success 200 {string} OssCheckCallback
// @Router /oss/callback [post]
func OssCheckCallback(g *gin.Context) {
	var ossInput model.OssForm
	if err := g.ShouldBind(&ossInput); err != nil {
		g.JSON(
			http.StatusBadRequest,
			gin.H{"msg": err.Error()},
		)
		return
	}
	config := commons.GetConfigJson()
	if config == nil {
		g.JSON(
			http.StatusBadRequest,
			gin.H{"msg": "配置错误"},
		)
		return
	}
	flag := CheckCheckSum(config.OssConfig.UserId, config.OssConfig.Seed, ossInput.Content, ossInput.CheckSum)
	if !flag {
		g.JSON(
			http.StatusBadRequest,
			gin.H{"msg": "签名错误"},
		)
		return

	}
	//  这里是收到回调之后的处理逻辑，随意编写

	g.JSON(200, "success")
}

func CheckCheckSum(userid string, seed string, content string, checksum string) bool {
	flag := false
	data := userid + seed + content
	sign := Sha256(data)
	if sign == checksum {
		flag = true
	}
	return flag
}

func Sha256(src string) string {
	m := sha256.New()
	m.Write([]byte(src))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

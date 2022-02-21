package apis

import (
	"IvanApi/commons"
	"IvanApi/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// @Summary 登录
// @Schemes
// @Description login
// @Tags login
// @Param id query string true "id"
// @Param name query string true "name"
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseResult Login
// @Router /login [get]
func Login(c *gin.Context) {
	var loginReq model.LoginInfo
	if c.ShouldBindQuery(&loginReq) == nil {
		// 登陆逻辑校验(查库，验证用户是否存在以及登陆信息是否正确)
		isPass, user := model.LoginCheck(&loginReq)
		// 验证通过后为该次请求生成token
		if isPass {
			generateToken(c, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":   -1,
				"msg":    "验证失败",
				"result": nil,
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":   -1,
			"msg":    "用户数据解析失败",
			"result": nil,
		})
	}
}

// token生成器
// md 为上面定义好的middleware中间件
func generateToken(c *gin.Context, user *model.LoginInfo) {
	// 构造SignKey: 签名和解签名需要使用一个值
	j := commons.NewJWT()

	// 构造用户claims信息(负荷)
	claims := commons.CustomClaims{
		user.Name,
		user.Id,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 1000), // 签名过期时间
			Issuer:    "gg",                            // 签名颁发者
		},
	}

	// 根据claims生成token对象
	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
			"data":   nil,
		})
	}

	log.Println(token)
	// 封装一个响应数据,返回用户名和token
	data := model.LoginResult{
		Name:  user.Name,
		Token: token,
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登陆成功",
		"data":   data,
	})
	return

}

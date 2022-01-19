package Apis

import (
	"IvanApi/Model"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
	"strconv"
)

// @BasePath /api/v1
// @Summary 获取广播
// @Schemes
// @Description get users
// @Tags ops
// @Param Authorization header string false "Bearer 用户令牌"
// @Param gameId query int true "游戏Id"
// @Param gameVersion query string true "游戏版本"
// @Param pageIndex query int true "pageIndex"
// @Param pageSize query int true "pageSize"
// @Accept json
// @Produce json
// @Success 200 {object} Model.PageResult GetBroadCast
// @Router /ops/getbroadcast [get]
func GetBroadCastList(g *gin.Context) {
	connArgs := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", "root", "密码", "Ip", 3306, "数据库")
	db, err := gorm.Open("mysql", connArgs)
	if err != nil {
		fmt.Println("%s", err.Error())
		panic("连接数据库失败")
	}
	defer db.Close()
	var gameId = g.Query("gameId")
	var gameVersion = g.Query("gameVersion")
	var pageIndex = g.Query("pageIndex")
	var pageSize = g.Query("pageSize")
	pagesize, _ := strconv.Atoi(pageSize)
	pageindex, _ := strconv.Atoi(pageIndex)
	skipNum := (pageindex - 1) * pagesize
	var total int
	var broadcasts []Model.BroadCast
	result := db.Where("game_id = ? And game_version = ?", gameId, gameVersion).Offset(skipNum).Limit(pagesize).Find(&broadcasts).Count(&total)
	if result.Error == nil {
		p := Model.PageResult{
			Total:     total,
			PageSize:  pagesize,
			PageIndex: pageindex,
			Result:    broadcasts,
		}
		g.JSON(http.StatusOK, p)
	} else {
		fmt.Println(result.Error)
		g.JSON(http.StatusBadRequest, nil)
	}
}

func UpdateBroadCast(g *gin.Context) {

}

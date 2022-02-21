package apis

import (
	"IvanApi/commons"
	"IvanApi/model"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
	"strconv"
)

// @Summary 获取广播
// @Schemes
// @Description get users
// @Tags ops
// @Param gameId query int true "游戏Id"
// @Param gameVersion query string true "游戏版本"
// @Param pageIndex query int true "pageIndex"
// @Param pageSize query int true "pageSize"
// @Accept json
// @Produce json
// @Success 200 {object} model.PageResult GetBroadCast
// @Router /ops/getbroadcast [get]
func GetBroadCastList(g *gin.Context) {
	db := commons.Getdb()
	val, err := commons.RedisClient.Get("123456").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
	var gameId = g.Query("gameId")
	var gameVersion = g.Query("gameVersion")
	var pageIndex = g.Query("pageIndex")
	var pageSize = g.Query("pageSize")
	pagesize, _ := strconv.Atoi(pageSize)
	pageindex, _ := strconv.Atoi(pageIndex)
	skipNum := (pageindex - 1) * pagesize
	var total int
	var broadcasts []model.BroadCast
	result := db.Where("game_id = ? And game_version = ?", gameId, gameVersion).Offset(skipNum).Limit(pagesize).Find(&broadcasts).Count(&total)
	if result.Error == nil {
		p := model.PageResult{
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

// @Summary 修改
// @Schemes
// @Description update
// @Tags ops
// @Param broadInput body model.BroadCastUpdateInput true "入参"
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseResult UpdateBroadCast
// @Router /ops/broadcast [put]
func UpdateBroadCast(g *gin.Context) {
	var broadInput model.BroadCastUpdateInput
	if err := g.ShouldBind(&broadInput); err != nil {
		g.JSON(
			http.StatusBadRequest,
			gin.H{"msg": err.Error()},
		)
		return
	}
	db := commons.Getdb()
	count := db.Table("broadcastinfo").Where("id = ?", broadInput.Id).Update(model.BroadCast{
		BroadcastText: broadInput.BroadcastText,
	}).RowsAffected
	var response model.ResponseResult
	if count > 0 {
		response.Code = 0
		response.Msg = "修改成功"
		response.Result = true
	}

	g.JSON(200, response)
}

// @Summary 新增
// @Schemes
// @Description update
// @Tags ops
// @Param broadInput body model.BroadCastUpdateInput true "信息"
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseResult AddBroadCast
// @Router /ops/broadcast [post]
func AddBroadCast(g *gin.Context) {
	var broadInput model.BroadCastUpdateInput
	if err := g.ShouldBind(&broadInput); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	db := commons.Getdb()
	var broadCast model.BroadCast
	broadCast.IsDeleted = 0
	broadCast.BroadcastText = broadInput.BroadcastText
	broadCast.BroadcastPosition = broadInput.BroadcastPosition
	broadCast.ChannelId = broadInput.ChannelId
	broadCast.EndTime = broadInput.EndTime
	broadCast.StartTime = broadInput.StartTime
	broadCast.GameId = broadInput.GameId
	broadCast.PartnerId = broadInput.PartnerId
	broadCast.GameVersion = broadInput.GameVersion
	broadCast.WorldId = broadInput.WorldId
	db.Table("broadcastinfo").Create(&broadCast)
	var response model.ResponseResult
	response.Code = 0
	response.Msg = "修改成功"
	response.Result = true
	g.JSON(200, response)
}

// @Summary 删除
// @Schemes
// @Description delete
// @Tags ops
// @Param id query int true "Id"
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseResult RemoveBroadCast
// @Router /ops/broadcast [delete]
func RemoveBroadCast(g *gin.Context) {
	var deleteInput model.BroadCastRemoveInput
	if err := g.ShouldBind(&deleteInput); err != nil {
		g.JSON(http.StatusOK, gin.H{"msg": err})
		return
	}
	var response model.ResponseResult
	db := commons.Getdb()
	count := db.Table("broadcastinfo").Where("id = ?", deleteInput.Id).Update(model.BroadCast{
		IsDeleted: 1,
	}).RowsAffected
	if count > 0 {
		response.Code = 0
		response.Msg = "删除成功"
		response.Result = true
	} else {
		response.Code = -1
		response.Msg = "删除失败"
		response.Result = false
	}
	g.JSON(200, response)
}

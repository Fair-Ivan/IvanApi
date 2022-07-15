package apis

import (
	"IvanApi/commons"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

// @Summary 测试
// @Schemes
// @Description test
// @Tags test
// @Accept json
// @param Authorization header string true "验证头"
// @Produce json
// @Success 200 {string} TestApi
// @Router /test [get]
func TestApi(g *gin.Context) {
	//commons.PublishMessage("发送消息")
	g.JSON(http.StatusOK, gin.H{
		"msg": "发送消息",
	})
}

// @Summary 测试2
// @Schemes
// @Description test2
// @Tags test
// @Accept json
// @Produce json
// @Success 200 {string} TestApi2
// @Router /test/second [get]
func TestApi2(g *gin.Context) {
	client := commons.GetMongoClient()
	// get collection
	collection := client.Database("GMCommand").Collection("Email")
	findOptions := options.Find()
	findOptions.SetLimit(2)
	findOptions.SetSkip(2)
	cur, err := collection.Find(context.Background(), bson.M{"WorldId": 4000}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	var all []*Email
	err = cur.All(context.Background(), &all)
	g.JSON(http.StatusOK, all)
}

// @Summary 测试3
// @Schemes
// @Description test3
// @Param id query int true "id"
// @Tags test
// @Accept json
// @Produce json
// @Success 200 {string} TestApi3
// @Router /test/third [get]
func TestApi3(g *gin.Context) {
	type input struct {
		Id int `form:"id" json:"id"`
	}
	var param input
	if err := g.ShouldBindQuery(&param); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"msg": "你好"})
		return
	}

	//var list = [5]int{1, 2, 3, 4, 5}
	//for i := 1; i <= len(list); i++ {
	//	item := list[i]
	//	if item == 10 {
	//		return
	//	}
	//}
	g.JSON(http.StatusOK, gin.H{"msg": "你好"})
	return
}

// @Summary 测试4
// @Schemes
// @Description test4
// @Tags test
// @Accept json
// @Produce json
// @Success 200 {string} TestApi4
// @Router /test/fifth [get]
func TestApi4(g *gin.Context) {
	panic("错误")
}

type Email struct {
	Id      primitive.ObjectID `bson:"_id"`
	EmailId string
	WorldId int
}

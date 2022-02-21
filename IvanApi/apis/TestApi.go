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
	commons.PublishMessage("发送消息")
	g.JSON(http.StatusOK, "发送消息")
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

type Email struct {
	Id      primitive.ObjectID `bson:"_id"`
	EmailId string
	WorldId int
}
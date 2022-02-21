package commons

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var client *mongo.Client
var clientOptions *options.ClientOptions

// 初始化MongoDB
func MongoInit() {
	config := GetConfigJson()
	clientOptions = options.Client().ApplyURI(config.MongoDBConnString)
	// Connect to MongoDB
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
}

// 获取mongo client
func GetMongoClient() *mongo.Client {
	if client.Ping(context.TODO(), nil) != nil {
		client, _ = mongo.Connect(context.TODO(), clientOptions)
	}
	return client
}

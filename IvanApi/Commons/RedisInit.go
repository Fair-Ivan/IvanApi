package Commons

import (
	"fmt"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func RedisInit() {
	redisConfig := GetConfigJson().RedisConfig

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host,
		Password: redisConfig.Password,
		DB:       redisConfig.Db,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		fmt.Println(err.Error())
		panic("redis ping error")
	}
}

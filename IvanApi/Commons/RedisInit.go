package Commons

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var RedisClient *redis.Client

func RedisInit() {
	redisConfig := GetConfigJson().RedisConfig
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         redisConfig.Host,
		Password:     redisConfig.Password,
		DB:           redisConfig.Db,
		MinIdleConns: 10,
		PoolSize:     15,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		fmt.Println(err.Error())
		panic("redis ping error")
	}
}

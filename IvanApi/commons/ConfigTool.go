package commons

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

type ConfigJson struct {
	ConnectionString   string      `json:"connectionString"`
	RedisConfig        RedisConfig `json:"redisConfig"`
	RabbitMqConnString string      `json:"rabbitMqConnString"`
	MongoDBConnString  string      `json:"MongoDBConnString"`
	OssConfig          OssConfig   `json:"ossConfig"`
}

type RedisConfig struct {
	Host     string `json:"host"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

type OssConfig struct {
	Seed   string `json:"seed"`
	UserId string `json:"userId"`
}

var (
	configJson *ConfigJson
	configMux  sync.RWMutex
)

func GetConfigJson() *ConfigJson {
	return configJson
}

func InitConfigJson(fliepath string) error {
	var config ConfigJson
	file, err := ioutil.ReadFile(fliepath)
	if err != nil {
		fmt.Println("配置文件读取错误,找不到配置文件", err)
		return err
	}
	if err = json.Unmarshal(file, &config); err != nil {
		fmt.Println("配置文件读取失败", err)
		return err
	}
	configMux.Lock()
	configJson = &config
	configMux.Unlock()
	return nil
}

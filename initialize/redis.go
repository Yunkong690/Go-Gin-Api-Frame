package initialize

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"personFrame/pkg/common"
)

var Redis *redis.Client

func InitRedis() *redis.Client {
	host := common.Conf.Redis.Host
	port := common.Conf.Redis.Port
	Password := common.Conf.Redis.Password
	db := common.Conf.Redis.DB
	addr := fmt.Sprintf("%s:%s", host, port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: Password,
		DB:       db,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		if err != nil {
			panic("连接Redis失败" + err.Error())
		}
	}
	Redis = client
	return client
}

func GetRedis() *redis.Client {
	return Redis
}

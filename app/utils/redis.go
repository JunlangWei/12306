package utils

import (
	"github.com/go-redis/redis/v8"
	"github.com/mamachengcheng/12306/app/common"
	"gopkg.in/ini.v1"
)

func GetRedisClient() *redis.Client {
	cfg, _ := ini.Load(common.ConfFilePath)

	redisCfg := cfg.Section("redis")
	address := redisCfg.Key("host").String() + ":" + redisCfg.Key("port").String()
	db, _ := redisCfg.Key("db").Int()

	client := redis.NewClient(&redis.Options{
		Addr:     address,
		DB:       db,
	})

	return client
}

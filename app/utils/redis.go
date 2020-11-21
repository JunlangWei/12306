package utils

import (
	"github.com/go-redis/redis/v8"
	"gopkg.in/ini.v1"
)

func GetRedisClient() *redis.Client {
	cfg, _ := ini.Load("app/conf/config.ini")

	redisCfg := cfg.Section("redis")
	address := redisCfg.Key("host").String() + ":" + redisCfg.Key("port").String()
	db, _ := redisCfg.Key("db").Int()

	client := redis.NewClient(&redis.Options{
		Addr:     address,
		DB:       db,
	})

	return client
}

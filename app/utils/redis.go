package utils

import (
	"github.com/go-redis/redis/v8"
	"gopkg.in/ini.v1"
)

func GetRedisClient() *redis.Client {
	cfg, _ := ini.Load("app/conf/config.ini")

	mysqlCfg := cfg.Section("redis")
	address := mysqlCfg.Key("http").String() + ":" + mysqlCfg.Key("port").String()
	password := mysqlCfg.Key("password").String()
	db, _ := mysqlCfg.Key("DB").Int()

	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})

	return client
}

package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/mamachengcheng/12306/app/utils"
)

var ctx = context.Background()

func TestRedis(c *gin.Context) {
	redisClient := utils.GetRedisClient()

	err := redisClient.Get(ctx, "key").Err()

	if err != redis.Nil{
		redisClient.Incr(ctx, "key")
	} else {
		redisClient.Set(ctx,"key", 0, 0)
	}

	val := redisClient.Get(ctx, "key").Val()

	c.JSON(200, gin.H{
		"message": val,
	})
}

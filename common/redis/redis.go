package redis

import (
	"TaoCoupon/common/logger"
	"TaoCoupon/config"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.Conf.Redis.Addr,
		Password: config.Conf.Redis.Password,
		DB:       config.Conf.Redis.DB,
	})
	logger.Logger.Info(fmt.Sprintf("redis init success: %v \n", RedisClient))
}

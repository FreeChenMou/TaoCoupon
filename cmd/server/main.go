package main

import (
	"TaoCoupon/app/server/router"
	"TaoCoupon/common/database"
	"TaoCoupon/common/logger"
	"TaoCoupon/common/redis"
	"TaoCoupon/common/snowflake"
	"TaoCoupon/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()
	logger.InitLogger()
	database.InitDB()
	redis.InitRedis()
	snowflake.InitSnowflake(1, 1)
	r := gin.Default()
	api := r.Group("/api/admin")

	router.InitCouponTemplateRouter(api)
	router.InitCouponTaskRouter(api)

	r.Run()
}

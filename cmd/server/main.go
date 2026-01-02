package main

import (
	"TaoCoupon/app/server/router"
	"TaoCoupon/common/database"
	"TaoCoupon/common/logger"
	"TaoCoupon/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()
	logger.InitLogger()
	database.InitDB()
	r := gin.Default()
	api := r.Group("/api/admin")

	router.InitCouponTemplateRouter(api)
	router.InitCouponTaskRouter(api)

	r.Run()
}

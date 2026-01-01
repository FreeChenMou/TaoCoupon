package main

import (
	"TaoCoupon/app/server/router"
	"TaoCoupon/common/database"
	"TaoCoupon/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()

	database.InitDB()
	r := gin.Default()
	api := r.Group("/api/admin")

	router.InitCouponTemplateRouter(api)
	router.InitCouponTaskRouter(api)

	r.Run()
}

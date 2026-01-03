package main

import (
	"TaoCoupon/app/engine/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api := r.Group("/api/engine")
	{
		router.InitUserCouponRouter(api)
		router.InitCouponTemplateRouter(api)
		router.InitCouponTemplateRemindRouter(api)
	}
	r.Run()
}

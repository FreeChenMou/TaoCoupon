package main

import (
	"TaoCoupon/app/consumer/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api := r.Group("/consumer")
	{
		router.InitCouponRouter(api)
	}
	r.Run()
}

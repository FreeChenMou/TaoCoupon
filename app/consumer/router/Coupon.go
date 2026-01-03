package router

import (
	"TaoCoupon/app/consumer/apis"
	"github.com/gin-gonic/gin"
)

func InitCouponRouter(r *gin.RouterGroup) gin.IRoutes {
	api := apis.Coupon{}
	group := r.Group("/couponTemplate")
	{
		group.GET("/query", api.ListQueryCoupons)
		group.GET("/query-sync", api.ListQueryCouponsBySync)
	}

	return r
}

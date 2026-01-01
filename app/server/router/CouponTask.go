package router

import (
	"TaoCoupon/app/server/apis"
	"github.com/gin-gonic/gin"
)

func InitCouponTaskRouter(r *gin.RouterGroup) gin.IRoutes {
	api := apis.CouponTask{}
	group := r.Group("/couponTask")
	{
		group.GET("/page", api.PageQueryCouponTask)
		group.GET("/find", api.FindCouponTaskById)
		group.POST("/create", api.CreateCouponTask)
	}
	return r
}

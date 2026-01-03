package router

import (
	"TaoCoupon/app/engine/apis"
	"github.com/gin-gonic/gin"
)

func InitCouponTemplateRemindRouter(r *gin.RouterGroup) gin.IRoutes {
	api := apis.CouponTemplateRemind{}
	group := r.Group("/remind")
	{
		group.GET("/list", api.ListCouponRemind)
		group.POST("/create", api.CreateCouponRemind)
		group.POST("/cancel", api.CancelCouponRemind)
	}

	return r
}

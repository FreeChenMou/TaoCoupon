package router

import (
	"TaoCoupon/app/server/apis"
	"github.com/gin-gonic/gin"
)

func InitCouponTemplateRouter(r *gin.RouterGroup) gin.IRoutes {
	api := apis.CouponTemplate{}
	group := r.Group("/couponTemplate")
	{
		group.GET("/page", api.PageQueryCouponTemplate)
		group.GET("/find", api.FindCouponTemplate)
		group.POST("/create", api.CreateCouponTemplate)
		group.POST("/increase", api.IncreaseNumberCouponTemplate)
		group.POST("/terminate", api.TerminateCouponTemplate)
	}

	return r
}

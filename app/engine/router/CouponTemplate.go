package router

import (
	"TaoCoupon/app/server/apis"
	"github.com/gin-gonic/gin"
)

func InitCouponTemplateRouter(r *gin.RouterGroup) gin.IRoutes {
	api := apis.CouponTemplate{}
	group := r.Group("/couponTemplate")
	{
		group.GET("/query", api.FindCouponTemplate)
	}

	return r
}

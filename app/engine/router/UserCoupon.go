package router

import (
	"TaoCoupon/app/engine/apis"
	"github.com/gin-gonic/gin"
)

func InitUserCouponRouter(r *gin.RouterGroup) gin.IRouter {
	api := apis.UserCoupon{}
	group := r.Group("/user-coupon")
	{
		group.POST("/redeem", api.RedeemUserCoupon)
		group.POST("/redeem-mq", api.RedeemUserCouponByMQ)
		group.POST("/process-payment", api.ProcessPayment)
		group.POST("/create-payment-record", api.CreatePaymentRecord)
		group.POST("/process-refund", api.ProcessRefund)
	}
	return r
}

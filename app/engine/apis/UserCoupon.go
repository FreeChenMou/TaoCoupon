package apis

import (
	"TaoCoupon/common/apis"
	"github.com/gin-gonic/gin"
)

type UserCoupon struct {
	apis.Api
}

func (a UserCoupon) RedeemUserCoupon(c *gin.Context) {

}

func (a UserCoupon) RedeemUserCouponByMQ(c *gin.Context) {}

func (a UserCoupon) CreatePaymentRecord(c *gin.Context) {}

func (a UserCoupon) ProcessPayment(c *gin.Context) {}

func (a UserCoupon) ProcessRefund(c *gin.Context) {}

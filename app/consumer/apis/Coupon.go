package apis

import (
	"TaoCoupon/common/apis"
	"github.com/gin-gonic/gin"
)

type Coupon struct {
	apis.Api
}

func (a Coupon) ListQueryCoupons(c *gin.Context) {

}

func (a Coupon) ListQueryCouponsBySync(c *gin.Context) {}

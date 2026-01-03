package apis

import (
	"TaoCoupon/common/apis"
	"github.com/gin-gonic/gin"
)

type CouponTemplateRemind struct {
	apis.Api
}

func (a CouponTemplateRemind) CreateCouponRemind(c *gin.Context) {

}

func (a CouponTemplateRemind) ListCouponRemind(c *gin.Context) {}

func (a CouponTemplateRemind) CancelCouponRemind(c *gin.Context) {}

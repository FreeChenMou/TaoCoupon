package models

import "TaoCoupon/app/models"

type CouponSettlement struct {
	ID            int     `json:"id"`
	UserID        int     `json:"user_id"`
	CouponID      int     `json:"coupon_id"`
	ShopNumber    int     `json:"shop_number"`
	TotalAmount   float64 `json:"total_amount"`   //结算总金额
	PayableAmount float64 `json:"payable_amount"` //应付金额
	CouponAmount  float64 `json:"coupon_amount"`  //优惠卷金额
	Status        int     `json:"status"`         //结算状态
	models.ModelTime
}

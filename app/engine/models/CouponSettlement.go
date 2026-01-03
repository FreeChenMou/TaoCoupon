package models

import "TaoCoupon/app/models"

type CouponSettlement struct {
	ID       int `json:"id"`
	OrderID  int `json:"order_id"`
	UserID   int `json:"user_id"`
	CouponID int `json:"coupon_id"`
	Status   int `json:"status"`
	models.ModelTime
}

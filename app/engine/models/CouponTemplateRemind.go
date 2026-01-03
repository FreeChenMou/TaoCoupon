package models

import "time"

type CouponTemplateRemindModel struct {
	UserId           int       `json:"userId"`
	CouponTemplateId int       `json:"coupon_template_id"`
	Information      int       `json:"information"`
	ShopNumber       int       `json:"shop_number"`
	StartTime        time.Time `json:"start_time"` //优惠卷开抢时间
}

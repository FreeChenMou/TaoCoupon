package models

import (
	"TaoCoupon/app/models"
	"time"
)

type UserCoupon struct {
	ID               int       `json:"id"`
	UserID           int       `json:"user_id"`
	CouponTemplateID int       `json:"coupon_template_id"` //优惠卷模板ID
	ReceiveTime      time.Time `json:"receive_time"`       //领取时间
	ReceiveCount     int       `json:"receive_count"`      //领取次数
	ValidStartTime   time.Time `json:"valid_start_time"`   //有效期开始时间
	ValidEndTime     time.Time `json:"valid_end_time"`     //有效期结束时间
	UseTime          time.Time `json:"use_time"`           //使用时间
	Source           int       `json:"source"`             //优惠卷来源 0：店铺卷 1：平台卷
	Status           int       `json:"status"`             //优惠卷状态 0：生效中 1：已结束
	models.ModelTime
}

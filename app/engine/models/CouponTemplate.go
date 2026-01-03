package models

import (
	"TaoCoupon/app/models"
	"time"
)

type CouponTemplate struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`             //优惠卷名称
	Source         int       `json:"source"`           //优惠卷来源 0：店铺卷 1：平台卷
	ShopNumber     int       `json:"shop_number"`      //店铺编号
	Target         int       `json:"target"`           //优惠对象 0：商品专属 1：全店通用
	Goods          int       `json:"goods"`            //优惠商品编码
	Type           int       `json:"type"`             //优惠类型 0：立减卷 1：满减卷 2：折扣卷
	ValidStartTime time.Time `json:"valid_start_time"` //有效期开始时间
	ValidEndTime   time.Time `json:"valid_end_time"`   //有效期结束时间
	Stock          int       `json:"stock"`            //库存
	ReceiveRule    int       `json:"receive_rule"`     //领取规则
	ConsumeRule    int       `json:"consume_rule"`     //消耗规则
	Status         int       `json:"status"`           //优惠卷状态 0：生效中 1：已结束
	models.ModelTime
}

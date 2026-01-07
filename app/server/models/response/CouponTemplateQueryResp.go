package models

import (
	"gorm.io/datatypes"
	"time"
)

type CouponTemplateQueryResp struct {
	Name           string         `json:"name"`                          //优惠卷名称
	Source         int            `json:"source"`                        //优惠卷来源 0：店铺卷 1：平台卷
	Target         int            `json:"target"`                        //优惠对象 0：商品专属 1：全店通用
	Goods          int            `json:"goods"`                         //优惠商品编码
	Type           int            `json:"type"`                          //优惠类型 0：立减卷 1：满减卷 2：折扣卷
	ValidStartTime time.Time      `json:"valid_start_time"`              //有效期开始时间
	ValidEndTime   time.Time      `json:"valid_end_time"`                //有效期结束时间
	Stock          int            `json:"stock"`                         //库存
	ReceiveRule    datatypes.JSON `json:"receive_rule" gorm:"type:json"` //领取规则
	ConsumeRule    datatypes.JSON `json:"consume_rule" gorm:"type:json"` //消耗规则
	Status         int            `json:"status"`                        //优惠卷状态 0：生效中 1：已结束
}

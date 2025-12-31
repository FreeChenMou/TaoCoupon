package models

import "TaoCoupon/app/models"

type CouponTemplateLog struct {
	Id               int    `json:"id"`
	ShopNumber       int    `json:"shop_number"`        //店铺编号
	CouponTemplateId int    `json:"coupon_template_id"` //优惠卷模板ID
	OperatorId       int    `json:"operator_id"`        //操作人ID
	OperatorLog      string `json:"operator_log"`       //操作日志
	OriginalData     string `json:"original_data"`      //原始数据
	ModifiedData     string `json:"modified_data"`      //修改后数据
	models.ModelTime
}

package request

type CouponTemplateQueryReq struct {
	Name   string `json:"name" form:"name"`     //优惠卷名称
	Target int    `json:"target" form:"target"` //优惠对象 0：商品专属 1：全店通用
	Goods  int    `json:"goods" form:"goods"`   //优惠商品编码
	Type   int    `json:"type" form:"type"`     //优惠类型 0：立减卷 1：满减卷 2：折扣卷
}

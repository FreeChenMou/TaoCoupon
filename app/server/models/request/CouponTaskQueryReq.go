package models

type CouponTaskQueryReq struct {
	BatchId          int    `json:"batch_id"`           //任务批次ID
	TaskName         string `json:"task_name"`          //优惠券批次任务名称
	CouponTemplateId int    `json:"coupon_template_id"` //优惠卷模板ID
	Status           int    `json:"status"`             //状态 0：待执行 1：执行中 2：执行失败 3：执行成功 4：取消
}

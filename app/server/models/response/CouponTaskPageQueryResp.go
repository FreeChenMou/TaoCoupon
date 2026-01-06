package models

import (
	"time"
)

type CouponTaskPageQueryResp struct {
	BatchId          int       `json:"batch_id"`           //任务批次ID
	TaskName         string    `json:"task_name"`          //优惠券批次任务名称
	SendNum          int       `json:"send_num"`           //优惠卷发放数量
	NotifyType       int       `json:"notify_type"`        //通知方式 0：站内信 1：弹框推送 2：邮箱 3：短信
	CouponTemplateId int       `json:"coupon_template_id"` //优惠卷模板ID
	SendType         int       `json:"send_type"`          //发送类型 0：立即发送 1：定时发送
	SendTime         time.Time `json:"send_time"`          //发送时间
	Status           int       `json:"status"`             //状态 0：待执行 1：执行中 2：执行失败 3：执行成功 4：取消
	CompletionDate   time.Time `json:"completion_date"`    //完成时间
	OperatorId       int       `json:"operator_id"`        //操作人ID
}

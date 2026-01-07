package request

import (
	"time"
)

type CouponTask struct {
	TaskName         string    `json:"task_name"`          //优惠券批次任务名称
	FileAddress      string    `json:"file_address"`       //文件地址
	NotifyType       int       `json:"notify_type"`        //通知方式 0：站内信 1：弹框推送 2：邮箱 3：短信
	CouponTemplateId int       `json:"coupon_template_id"` //优惠卷模板ID
	SendType         int       `json:"send_type"`          //发送类型 0：立即发送 1：定时发送
	SendTime         time.Time `json:"send_time"`          //发送时间
}

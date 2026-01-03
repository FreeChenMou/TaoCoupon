package models

type CouponTaskFail struct {
	ID         int    `json:"id"`
	BatchId    string `json:"batch_id"`
	JsonObject string `json:"json_object"`
}

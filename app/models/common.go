package models

import "time"

type ModelTime struct {
	CreatedAt time.Time `json:"created_time" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_time" gorm:"autoUpdateTime"`
	DelFlag   int       `json:"del_flag" gorm:"type:int;default:0"`
}

package models

import "time"

type ModelTime struct {
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	DelFlag   int       `json:"del_flag" gorm:"type:int;default:0"`
}

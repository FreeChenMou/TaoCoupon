package service

import "gorm.io/gorm"

type Service struct {
	DB    *gorm.DB
	Msg   string
	MsgID string
	Error error
}

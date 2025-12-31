package models

import "TaoCoupon/app/models"

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	ShopNumber int    `json:"shop_number"` //店铺编号
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Password   string `json:"password"`
	models.ModelTime
}

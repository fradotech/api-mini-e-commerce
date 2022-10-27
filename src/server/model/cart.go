package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserId    string `json:"user_id"`
	ProductId string `json:"product_id"`
	Qty       uint16 `json:"qty"`
	Subtotal  uint32 `json:"subtotal"`

	User    User    `json:"user"`
	Product Product `json:"product"`
}

var Carts = []Cart{}

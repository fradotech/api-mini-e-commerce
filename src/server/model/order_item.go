package model

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderId   string `json:"order_id"`
	ProductId string `json:"product_id"`
	Qty       uint16 `json:"qty"`
	SubTotal  uint32 `json:"subtotal"`

	Order   Order   `json:"order"`
	Product Product `json:"product"`
}

var OrderItems = []OrderItem{}

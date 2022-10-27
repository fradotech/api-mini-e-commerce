package model

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	OrderId  string `json:"order_id"`
	Amount   uint32 `json:"amount"`
	Provider string `json:"provider"`
	Status   string `json:"status"`

	Order Order `json:"order"`
}

var Payments = []Payment{}

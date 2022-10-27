package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserId string `json:"user_id"`
	Total  uint32 `json:"total"`

	OrderItem []OrderItem `json:"order_item"`
	Payment   []Payment   `json:"payment"`
	User      User        `json:"user"`
}

var Orders = []Order{}

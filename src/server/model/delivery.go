package model

import "gorm.io/gorm"

type Delivery struct {
	gorm.Model
	OrderId       string `json:"order_id"`
	ReceiptNumber string `json:"receipt_number"`
	Expedition    string `json:"expedition"`
	Weight        uint32 `json:"weight"`
	ServiceType   string `json:"service_type"`
	Cost          uint32 `json:"cost"`
	Estimation    string `json:"estimation"`

	Order Order `json:"Order"`
}

var Deliveries = []Delivery{}

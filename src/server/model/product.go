package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name     	string `json:"name"`
	Price      	uint32 `json:"price"`
	Stock      	uint16 `json:"stock"`
	Description string `json:"description"`
	CategoryId 	string `json:"category_id"`
	UserId   	string `json:"user_id"`
	
	Category Category `json:"category"`
	User User `json:"user"`
}

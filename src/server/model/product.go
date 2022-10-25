package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name     string `json:"name"`
	Category string `json:"category"`
	Desc     string `json:"desc"`
	UserId   string `json:"user_id"`

	User User `json:"user"`
}
